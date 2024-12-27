package main

import (
	"context"
	"log"
	"net/http"
	"net/url"

	"github.com/BestNathan/deribit-api/clients/websocket"
	"github.com/BestNathan/deribit-api/pkg/deribit"
)

func main() {
	cfg := deribit.GetConfig()

	cfg.AutoStart = false
	cfg.Client = &http.Client{
		Transport: &http.Transport{
			Proxy: func(r *http.Request) (*url.URL, error) {
				return url.Parse("socks5://127.0.0.1:7890")
			},
		},
	}

	client := websocket.NewDeribitWsClient(cfg)

	if err := client.Start(context.Background()); err != nil {
		log.Fatal("start fail ", err)
	}

	if _, err := client.GetTime(); err != nil {
		log.Fatal("GetTime fail ", err)
	}

	if _, err := client.Test(); err != nil {
		log.Fatal("Test fail ", err)
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
