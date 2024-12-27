package deribit

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	BaseURL     = "https://www.deribit.com/api/v2"
	TestBaseURL = "https://test.deribit.com/ws/api/v2"
	WSURL       = "wss://www.deribit.com/ws/api/v2"
	TestWSURL   = "wss://test.deribit.com/ws/api/v2"
)

const (
	MaxTryTimes = 10000
)

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

type Credential struct {
	ApiKey    string
	SecretKey string
}

type WebsocketConfiguration struct {
	Url                  string
	AutoReconnect        bool
	ReconnectDuration    time.Duration
	AutoStart            bool
	ReadLimit            int64
	DialWebsocketTimeout time.Duration
	CallTimeout          time.Duration
	TestDuration         time.Duration
	HeartBeatInterval    float64
}

type HttpConfiguration struct {
	BaseUrl string
}

type Configuration struct {
	*WebsocketConfiguration
	*HttpConfiguration
	Ctx        context.Context
	Client     *http.Client
	Debug      bool
	Credential Credential
	Logger     *logrus.Logger
}

func GetConfig() *Configuration {
	autoReconnect, _ := strconv.ParseBool(getEnvWithDefault("DERIBIT_AUTO_RECONNECT", "true"))
	debugMode, _ := strconv.ParseBool(getEnvWithDefault("DERIBIT_DEBUG_MODE", "true"))
	realMode, _ := strconv.ParseBool(getEnvWithDefault("DERIBIT_REAL_MODE", "false"))

	wsUrl := TestWSURL
	baseUrl := TestBaseURL
	if realMode {
		wsUrl = WSURL
		baseUrl = BaseURL
	}

	// Configure logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	if debugMode {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	return &Configuration{
		Credential: Credential{
			ApiKey:    getEnvWithDefault("DERIBIT_API_KEY", ""),
			SecretKey: getEnvWithDefault("DERIBIT_API_SECRET", ""),
		},
		WebsocketConfiguration: &WebsocketConfiguration{
			Url:                  wsUrl,
			AutoReconnect:        autoReconnect,
			AutoStart:            true,
			TestDuration:         time.Second * 3,
			ReconnectDuration:    time.Second,
			CallTimeout:          time.Minute,
			HeartBeatInterval:    30,
			DialWebsocketTimeout: time.Second * 10,
		},
		HttpConfiguration: &HttpConfiguration{
			BaseUrl: baseUrl,
		},
		Client: http.DefaultClient,
		Debug:  debugMode,
		Logger: logger,
	}
}
