package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/BestNathan/deribit-api/clients/websocket"
	"github.com/BestNathan/deribit-api/pkg/deribit"
)

func main() {
	cfg := deribit.GetConfig()

	cfg.Client = &http.Client{
		Transport: &http.Transport{
			Proxy: func(r *http.Request) (*url.URL, error) {
				return url.Parse("socks5://127.0.0.1:7890")
			},
		},
	}

	client := websocket.NewDeribitWsClient(cfg)

	_, gErr := client.GetTime()
	if gErr != nil {
		return
	}
	_, tErr := client.Test()
	if tErr != nil {
		return
	}

	btcdvolch := websocket.ChannelDeribitVolatilityIndex(websocket.DERIBIT_VOLATILITY_INDEX_NAME_BTC)

	client.On(btcdvolch, func(data string) {
		log.Println(data)
	})

	client.Subscribe([]string{
		btcdvolch,
	})

	forever := make(chan bool)
	<-forever
}
