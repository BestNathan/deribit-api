package models

import models2 "github.com/xingxing/deribit-api/clients/websocket/models"

type StopOrder struct {
	Trigger        string        `json:"trigger"`
	Timestamp      int64         `json:"timestamp"`
	StopPrice      float64       `json:"stop_price"`
	StopID         string        `json:"stop_id"`
	OrderState     string        `json:"order_state"`
	Request        string        `json:"request"`
	Price          models2.Price `json:"price"`
	OrderID        string        `json:"order_id"`
	Offset         float64       `json:"offset"`
	InstrumentName string        `json:"instrument_name"`
	Amount         float64       `json:"amount"`
	Direction      string        `json:"direction"`
}
