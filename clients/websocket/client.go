package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	websocketmodels "github.com/BestNathan/deribit-api/clients/websocket/models"
	"github.com/BestNathan/deribit-api/pkg/deribit"
	"github.com/BestNathan/deribit-api/pkg/models"
	"github.com/sirupsen/logrus"

	"github.com/chuckpreslar/emission"
	"github.com/coder/websocket"
	"github.com/sourcegraph/jsonrpc2"
)

var (
	ErrUnAuthorized          = errors.New("websocket unauthorized")
	ErrWebsocketNotConnected = errors.New("websocket not connected")
)

type DeribitWSClient struct {
	ctx context.Context
	cfg *deribit.WebsocketConfiguration

	client *http.Client

	credential     deribit.Credential
	authentication *websocketmodels.Authentication

	// conn
	conn            *websocket.Conn
	rpcConn         *jsonrpc2.Conn
	stopheartbeatch chan struct{}
	connected       atomic.Bool

	// subs
	subscriptions    []string
	subscriptionsMap map[string]struct{}

	// pub/sub
	emitter *emission.Emitter

	logger *logrus.Logger
}

func NewDeribitWsClient(cfg *deribit.Configuration) *DeribitWSClient {
	ctx := cfg.Ctx
	if ctx == nil {
		ctx = context.Background()
	}

	client := &DeribitWSClient{
		ctx:              ctx,
		cfg:              cfg.WebsocketConfiguration,
		client:           cfg.Client,
		credential:       cfg.Credential,
		logger:           cfg.Logger,
		subscriptionsMap: make(map[string]struct{}),
		emitter:          emission.NewEmitter(),
	}

	if cfg.AutoStart {
		if err := client.start(ctx); err != nil {
			client.logger.WithContext(ctx).Warnln("websocket start fail", err)
			panic(err)
		}
	}

	return client
}

func (c *DeribitWSClient) Start(ctx context.Context) error {
	if c.IsConnected() {
		return nil
	}

	return c.start(ctx)
}

// setIsConnected sets state for isConnoected
func (c *DeribitWSClient) setIsConnected(state bool) {
	c.connected.Store(state)
}

// IsConnected returns the WebSocket connection state
func (c *DeribitWSClient) IsConnected() bool {
	return c.connected.Load()
}

func (c *DeribitWSClient) Subscribe(channels []string) {
	c.subscriptions = append(c.subscriptions, channels...)
	c.subscribe()
}

func (c *DeribitWSClient) subscribe() {
	var publicChannels []string
	var privateChannels []string

	for _, v := range c.subscriptions {
		if _, ok := c.subscriptionsMap[v]; ok {
			continue
		}
		if strings.HasPrefix(v, "user.") {
			privateChannels = append(privateChannels, v)
		} else {
			publicChannels = append(publicChannels, v)
		}
	}

	if len(publicChannels) > 0 {
		_, err := c.PublicSubscribe(&models.SubscribeParams{
			Channels: publicChannels,
		})
		if err != nil {
			return
		}
	}
	if len(privateChannels) > 0 {
		_, err := c.PrivateSubscribe(&models.SubscribeParams{
			Channels: privateChannels,
		})
		if err != nil {
			return
		}
	}

	allChannels := append(publicChannels, privateChannels...)
	for _, v := range allChannels {
		c.subscriptionsMap[v] = struct{}{}
	}
}

func (c *DeribitWSClient) start(ctx context.Context) error {
	ctx = context.WithoutCancel(ctx)

	c.setIsConnected(false)
	c.subscriptionsMap = make(map[string]struct{})
	c.conn = nil
	c.rpcConn = nil

	for i := 0; i < deribit.MaxTryTimes; i++ {
		conn, _, err := c.dial(ctx)
		if err != nil {
			dur := time.Duration((i+1)*5) * time.Second

			c.logger.
				WithContext(ctx).
				Warnf("websocket dial fail(%d), and will retry in %s\n", i, dur)

			time.Sleep(dur)

			continue
		}

		c.conn = conn
		break
	}

	if c.conn == nil {
		return errors.New("websocket dial fail")
	}

	// Create a new object stream with the websocket connection
	stream := websocketmodels.NewObjectStream(c.conn)

	// Initialize the JSON-RPC connection with the stream
	c.rpcConn = jsonrpc2.NewConn(ctx, stream, c)

	c.setIsConnected(true)

	// Authenticate if credentials are provided
	if c.credential.ApiKey != "" && c.credential.SecretKey != "" {
		if err := c.Auth(c.credential.ApiKey, c.credential.SecretKey); err != nil {
			return fmt.Errorf("auth: %w", err)
		}
	}

	// Subscribe to channels
	c.subscribe()

	// Set heartbeat
	_, err := c.SetHeartbeat(&models.SetHeartbeatParams{Interval: c.cfg.HeartBeatInterval})
	if err != nil {
		return fmt.Errorf("set heartbeat: %w", err)
	}

	// Start reconnection handler if enabled
	if c.cfg.AutoReconnect {
		go c.reconnect(ctx)
	}

	// Start heartbeat routine
	go c.startheartbeat(ctx)

	return nil
}

// Call issues JSONRPC v2 calls
func (c *DeribitWSClient) Call(method string, params interface{}, result interface{}) (err error) {
	timeout := c.cfg.CallTimeout
	if timeout == 0 {
		timeout = time.Minute
	}

	ctx, cancel := context.WithTimeout(context.WithoutCancel(c.ctx), timeout)
	defer cancel()

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recover: %v", r)
		}
	}()

	if !c.IsConnected() {
		return ErrWebsocketNotConnected
	}

	if params == nil {
		params = websocketmodels.EmptyParams
	}

	if token, ok := params.(websocketmodels.PrivateParams); ok {
		if c.authentication == nil || c.authentication.AccessToken == "" {
			return ErrUnAuthorized
		}

		token.SetToken(c.authentication.AccessToken)
	}

	if err := c.rpcConn.Call(ctx, method, params, result); err != nil {
		return fmt.Errorf("jsonrpc call: %w", err)
	} else {
		return nil
	}
}

// Handle implements jsonrpc2.Handler
func (c *DeribitWSClient) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	if req.Method == "subscription" {
		// update events
		if req.Params != nil && len(*req.Params) > 0 {
			var event websocketmodels.Event
			if err := json.Unmarshal(*req.Params, &event); err != nil {
				c.logger.WithContext(ctx).Warnln("websocket unmarshal event fail", err)
				return
			}

			if err := c.subscriptionsProcess(&event); err != nil {
				c.logger.WithContext(ctx).Warnln("websocket subscription process fail", err)
			}
		}
	}
}

func (c *DeribitWSClient) startheartbeat(ctx context.Context) {
	c.stopheartbeatch = make(chan struct{})

	t := time.NewTicker(c.cfg.TestDuration)
	for {
		select {
		case <-t.C:
			if _, err := c.Test(); err != nil {
				c.logger.WithContext(ctx).Warnln("heartbeat test fail", err)
				return
			}
		case <-c.stopheartbeatch:
			c.logger.WithContext(ctx).Debugln("heartbeat stop")
			return
		case <-ctx.Done():
			c.logger.WithContext(ctx).Debugln("heartbeat ctx done")
			return
		}
	}
}

func (c *DeribitWSClient) stopheartbeat() {
	if c.stopheartbeatch != nil {
		close(c.stopheartbeatch)
	}
}

func (c *DeribitWSClient) reconnect(ctx context.Context) {
	notify := c.rpcConn.DisconnectNotify()
	<-notify

	c.stopheartbeat()

	c.logger.WithContext(ctx).Debugln("jsonrpc conn disconnected, reconnecting...")

	if c.cfg.ReconnectDuration > 0 {
		time.Sleep(c.cfg.ReconnectDuration)
	}

	err := c.start(ctx)
	if err != nil {
		c.logger.WithContext(ctx).Warnln("reconnect fail", err)
	} else {
		c.logger.WithContext(ctx).Debugln("reconnect success")
	}
}

func (c *DeribitWSClient) dial(ctx context.Context) (*websocket.Conn, *http.Response, error) {
	if c.client != nil && c.client.Timeout == 0 {
		c.client.Timeout = c.cfg.DialWebsocketTimeout
	}

	conn, resp, err := websocket.Dial(ctx, c.cfg.Url, &websocket.DialOptions{
		HTTPClient: c.client,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("websocket dial: %w", err)
	}

	if c.cfg.ReadLimit != 0 {
		conn.SetReadLimit(c.cfg.ReadLimit)
	}

	return conn, resp, nil
}
