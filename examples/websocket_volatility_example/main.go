package main

import (
	"log"

	"github.com/BestNathan/deribit-api/clients/websocket"
	"github.com/BestNathan/deribit-api/pkg/deribit"
)

func main() {
	cfg := deribit.GetConfig()
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
