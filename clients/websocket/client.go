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
	ErrAuthenticationIsRequired = errors.New("authentication is required")
)

type DeribitWSClient struct {
	ctx           context.Context
	url           string
	credential    deribit.Credential
	autoReconnect bool
	client        *http.Client
	conn          *websocket.Conn
	rpcConn       *jsonrpc2.Conn
	heartCancel   chan struct{}
	isConnected   atomic.Bool

	auth struct {
		token   string
		refresh string
	}

	subscriptions    []string
	subscriptionsMap map[string]struct{}

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
		url:              cfg.WebsocketConfiguration.Url,
		client:           cfg.Client,
		credential:       cfg.Credential,
		autoReconnect:    cfg.AutoReconnect,
		logger:           cfg.Logger,
		subscriptionsMap: make(map[string]struct{}),
		emitter:          emission.NewEmitter(),
	}
	err := client.start()
	if err != nil {
		client.logger.WithContext(ctx).Warnln("start fail", err)

		panic(err)
	}
	return client
}

// setIsConnected sets state for isConnoected
func (c *DeribitWSClient) setIsConnected(state bool) {
	c.isConnected.Store(state)
}

// IsConnected returns the WebSocket connection state
func (c *DeribitWSClient) IsConnected() bool {
	return c.isConnected.Load()
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

func (c *DeribitWSClient) start() error {
	c.setIsConnected(false)
	c.subscriptionsMap = make(map[string]struct{})
	c.conn = nil
	c.rpcConn = nil
	c.heartCancel = make(chan struct{})

	for i := 0; i < deribit.MaxTryTimes; i++ {
		conn, _, err := c.connect()
		if err != nil {
			tm := (i + 1) * 5
			dur := time.Duration(tm) * time.Second

			c.logger.
				WithContext(c.ctx).
				Warnf("websocket connect fail(%d), and will retry in %s\n", i, dur)

			time.Sleep(dur)

			continue
		}

		c.conn = conn
		break
	}

	if c.conn == nil {
		return errors.New("websocket connect fail")
	}

	// Create a new object stream with the websocket connection
	stream := websocketmodels.NewObjectStream(c.conn)

	// Initialize the JSON-RPC connection with the stream
	c.rpcConn = jsonrpc2.NewConn(c.ctx, stream, c)

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
	_, err := c.SetHeartbeat(&models.SetHeartbeatParams{Interval: 30})
	if err != nil {
		return fmt.Errorf("set heartbeat: %w", err)
	}

	// Start reconnection handler if enabled
	if c.autoReconnect {
		go c.reconnect()
	}

	// Start heartbeat routine
	go c.heartbeat()

	return nil
}

// Call issues JSONRPC v2 calls
func (c *DeribitWSClient) Call(method string, params interface{}, result interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	if !c.IsConnected() {
		return errors.New("not connected")
	}
	if params == nil {
		params = websocketmodels.EmptyParams
	}

	if token, ok := params.(websocketmodels.PrivateParams); ok {
		if c.auth.token == "" {
			return ErrAuthenticationIsRequired
		}
		token.SetToken(c.auth.token)
	}

	return c.rpcConn.Call(c.ctx, method, params, result)
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

func (c *DeribitWSClient) heartbeat() {
	t := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-t.C:
			if _, err := c.Test(); err != nil {
				c.logger.WithContext(c.ctx).Warnln("test fail", err)
				return
			}
		case <-c.heartCancel:
			return
		}
	}
}

func (c *DeribitWSClient) reconnect() {
	notify := c.rpcConn.DisconnectNotify()
	<-notify
	c.setIsConnected(false)

	c.logger.WithContext(c.ctx).Debugln("reconnecting...")

	close(c.heartCancel)

	time.Sleep(1 * time.Second)

	err := c.start()
	if err != nil {
		c.logger.WithContext(c.ctx).Warnln("reconnect fail", err)

		return
	} else {
		c.logger.WithContext(c.ctx).Debugln("reconnect success")
	}
}

func (c *DeribitWSClient) connect() (*websocket.Conn, *http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, resp, err := websocket.Dial(ctx, c.url, &websocket.DialOptions{
		HTTPClient: c.client,
	})
	if err == nil {
		conn.SetReadLimit(32768 * 64)
	}
	return conn, resp, err
}
