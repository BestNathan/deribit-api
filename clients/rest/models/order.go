package models

import (
	"github.com/BestNathan/deribit-api/pkg/models"
)

// Assuming Trade, Direction, decimalFormat are defined elsewhere

type OrderResult struct {
	Trades []models.Trade `json:"trades"`
	Order  Order          `json:"order"`
}

type OrderType string

const (
	Market OrderType = "market"
	Limit  OrderType = "limit"
)

type OrderStatus string

const (
	Open      OrderStatus = "open"
	Filled    OrderStatus = "filled"
	Cancelled OrderStatus = "cancelled"
	Rejected  OrderStatus = "rejected"
)

type Order struct {
	Web            bool        `json:"web"`
	TimeInForce    string      `json:"time_in_force"`
	Replaced       bool        `json:"replaced"`
	ReduceOnly     bool        `json:"reduce_only"`
	Price          float64     `json:"price"                    decimal_format:"deserialize"`
	PostOnly       bool        `json:"post_only"`
	OrderType      OrderType   `json:"order_type"`
	Status         OrderStatus `json:"order_state"`
	OrderID        string      `json:"order_id"`
	MaxShow        float64     `json:"max_show"                 decimal_format:"deserialize"`
	LastUpdate     int64       `json:"last_update_timestamp"`
	Label          string      `json:"label"`
	IsRebalance    *bool       `json:"is_rebalance"`
	IsLiquidation  bool        `json:"is_liquidation"`
	InstrumentName string      `json:"instrument_name"`
	FilledAmount   float64     `json:"filled_amount"            decimal_format:"deserialize"`
	Direction      Direction   `json:"direction"`
	Timestamp      int64       `json:"creation_timestamp"`
	AveragePrice   float64     `json:"average_price"            decimal_format:"deserialize"`
	API            bool        `json:"api"`
	Amount         float64     `json:"amount"                   decimal_format:"deserialize"`
}

type Orders []Order

func (orders Orders) Len() int {
	return len(orders)
}

func (orders Orders) IsEmpty() bool {
	return len(orders) == 0
}

func (orders Orders) Iter() []Order {
	return orders
}

type CancelOrderResponse struct {
	Triggered           *bool    `json:"triggered"`
	Trigger             *string  `json:"trigger"`
	TimeInForce         string   `json:"time_in_force"`
	TriggerPrice        *float64 `json:"trigger_price"`
	ReduceOnly          bool     `json:"reduce_only"`
	Price               float64  `json:"price"`
	PostOnly            bool     `json:"post_only"`
	OrderType           string   `json:"order_type"`
	OrderState          string   `json:"order_state"`
	OrderID             string   `json:"order_id"`
	MaxShow             float64  `json:"max_show"`
	LastUpdateTimestamp int64    `json:"last_update_timestamp"`
	Label               string   `json:"label"`
	IsRebalance         *bool    `json:"is_rebalance"`
	IsLiquidation       bool     `json:"is_liquidation"`
	InstrumentName      string   `json:"instrument_name"`
	Direction           string   `json:"direction"`
	CreationTimestamp   int64    `json:"creation_timestamp"`
	API                 bool     `json:"api"`
	Amount              float64  `json:"amount"`
}
