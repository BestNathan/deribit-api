package models

import models2 "github.com/joaquinbejar/deribit-api/internal/websocket/models"

type UserChangesNotification struct {
	Trades    []UserTrade     `json:"trades"`
	Positions []Position      `json:"positions"`
	Orders    []models2.Order `json:"orders"`
}