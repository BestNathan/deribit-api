package deribit

import "context"

const (
	RealBaseURL     = "wss://www.deribit.com/ws/api/v2/"
	TestBaseURL     = "wss://test.deribit.com/ws/api/v2/"
	RealRestBaseURL = "https://www.deribit.com/ws/api/v2/"
	TestRestBaseURL = "https://test.deribit.com/ws/api/v2/"
)

const (
	MaxTryTimes = 10000
)

type Configuration struct {
	Ctx           context.Context
	Addr          string `json:"addr"`
	ApiKey        string `json:"api_key"`
	SecretKey     string `json:"secret_key"`
	AutoReconnect bool   `json:"auto_reconnect"`
	DebugMode     bool   `json:"debug_mode"`
	WSBaseURL     string `json:"ws_base_url"`
	RestBaseURL   string `json:"rest_base_url"`
}
