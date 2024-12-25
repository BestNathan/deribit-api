package models

import models2 "github.com/BestNathan/deribit-api/clients/websocket/models"

type ClosePositionResponse struct {
	Trades []Trade       `json:"trades"`
	Order  models2.Order `json:"order"`
}
