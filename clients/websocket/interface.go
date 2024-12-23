package websocket

import (
	models2 "github.com/xingxing/deribit-api/clients/websocket/models"
	"github.com/xingxing/deribit-api/pkg/models"
)

type Behavior interface {
	AccountBehavior
	MarketBehavior
	TradingBehavior
}

type MarketBehavior interface {
	GetOrderBook(*models.GetOrderBookParams) (models.GetOrderBookResponse, error)

	GetMarkPriceHistory(*models.GetMarkPriceHistoryParams) (models.MarkPriceHistory, error)
}

type AccountBehavior interface {
	GetPositions(*models.GetPositionsParams) ([]models.Position, error)
}

type TradingBehavior interface {
	Buy(*models.BuyParams) (models.BuyResponse, error)
	Sell(*models.SellParams) (models.SellResponse, error)

	GetOrderHistoryByInstrument(*models.GetOrderHistoryByInstrumentParams) ([]models2.Order, error)
	CancellByLabel(*models.CancelByLabelParams) (int, error)
}

var _ Behavior = (*DeribitWSClient)(nil)
