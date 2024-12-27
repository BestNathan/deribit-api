package main

import (
	"encoding/json"
	"fmt"
	"github.com/BestNathan/deribit-api/clients/websocket"
	"github.com/BestNathan/deribit-api/pkg/deribit"
	"github.com/BestNathan/deribit-api/pkg/models"
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

	//client.On("book.BTC-PERPETUAL.raw", func(e *models.OrderBookRawNotification) {
	//	jsonData, _ := json.Marshal(e)
	//	fmt.Printf("%s\n", string(jsonData))
	//})

	client.On("book.BTC-PERPETUAL.raw", func(e *models.OrderBookRawNotification) {
		jsonData, err := json.MarshalIndent(e, "", "    ")
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v\n", err)
			return
		}

		fmt.Printf("%s\n", string(jsonData))
	})

	client.Subscribe([]string{
		"book.BTC-PERPETUAL.raw",
	})

	forever := make(chan bool)
	<-forever
}
