package websocket

import (
	models2 "github.com/xingxing/deribit-api/clients/websocket/models"
	"github.com/xingxing/deribit-api/pkg/models"
)

func (c *DeribitWSClient) Buy(params *models.BuyParams) (result models.BuyResponse, err error) {
	err = c.Call("private/buy", params, &result)
	return
}

func (c *DeribitWSClient) Sell(params *models.SellParams) (result models.SellResponse, err error) {
	err = c.Call("private/sell", params, &result)
	return
}

func (c *DeribitWSClient) Edit(params *models.EditParams) (result models.EditResponse, err error) {
	err = c.Call("private/edit", params, &result)
	return
}

func (c *DeribitWSClient) Cancel(params *models.CancelParams) (result models2.Order, err error) {
	err = c.Call("private/cancel", params, &result)
	return
}

func (c *DeribitWSClient) CancelAll() (result string, err error) {
	err = c.Call("private/cancel_all", nil, &result)
	return
}

func (c *DeribitWSClient) CancelAllByCurrency(params *models.CancelAllByCurrencyParams) (result string, err error) {
	err = c.Call("private/cancel_all_by_currency", params, &result)
	return
}

func (c *DeribitWSClient) CancelAllByInstrument(params *models.CancelAllByInstrumentParams) (result string, err error) {
	err = c.Call("private/cancel_all_by_instrument", params, &result)
	return
}

func (c *DeribitWSClient) CancellByLabel(params *models.CancelByLabelParams) (result int, err error) {
	err = c.Call("private/cancel_by_label", params, &result)
	return
}

func (c *DeribitWSClient) ClosePosition(params *models.ClosePositionParams) (result models.ClosePositionResponse, err error) {
	err = c.Call("private/close_position", params, &result)
	return
}

func (c *DeribitWSClient) GetMargins(params *models.GetMarginsParams) (result models.GetMarginsResponse, err error) {
	err = c.Call("private/get_margins", params, &result)
	return
}

func (c *DeribitWSClient) GetOpenOrdersByCurrency(params *models.GetOpenOrdersByCurrencyParams) (result []models2.Order, err error) {
	err = c.Call("private/get_open_orders_by_currency", params, &result)
	return
}

func (c *DeribitWSClient) GetOpenOrdersByInstrument(params *models.GetOpenOrdersByInstrumentParams) (result []models2.Order, err error) {
	err = c.Call("private/get_open_orders_by_instrument", params, &result)
	return
}

func (c *DeribitWSClient) GetOrderHistoryByCurrency(params *models.GetOrderHistoryByCurrencyParams) (result []models2.Order, err error) {
	err = c.Call("private/get_order_history_by_currency", params, &result)
	return
}

func (c *DeribitWSClient) GetOrderHistoryByInstrument(params *models.GetOrderHistoryByInstrumentParams) (result []models2.Order, err error) {
	err = c.Call("private/get_order_history_by_instrument", params, &result)
	return
}

func (c *DeribitWSClient) GetOrderMarginByIDs(params *models.GetOrderMarginByIDsParams) (result models.GetOrderMarginByIDsResponse, err error) {
	err = c.Call("private/get_order_margin_by_ids", params, &result)
	return
}

func (c *DeribitWSClient) GetOrderState(params *models.GetOrderStateParams) (result models2.Order, err error) {
	err = c.Call("private/get_order_state", params, &result)
	return
}

func (c *DeribitWSClient) GetStopOrderHistory(params *models.GetStopOrderHistoryParams) (result models.GetStopOrderHistoryResponse, err error) {
	err = c.Call("private/get_stop_order_history", params, &result)
	return
}

func (c *DeribitWSClient) GetUserTradesByCurrency(params *models.GetUserTradesByCurrencyParams) (result models.GetUserTradesResponse, err error) {
	err = c.Call("private/get_user_trades_by_currency", params, &result)
	return
}

func (c *DeribitWSClient) GetUserTradesByCurrencyAndTime(params *models.GetUserTradesByCurrencyAndTimeParams) (result models.GetUserTradesResponse, err error) {
	err = c.Call("private/get_user_trades_by_currency_and_time", params, &result)
	return
}

func (c *DeribitWSClient) GetUserTradesByInstrument(params *models.GetUserTradesByInstrumentParams) (result models.GetUserTradesResponse, err error) {
	err = c.Call("private/get_user_trades_by_instrument", params, &result)
	return
}

func (c *DeribitWSClient) GetUserTradesByInstrumentAndTime(params *models.GetUserTradesByInstrumentAndTimeParams) (result models.GetUserTradesResponse, err error) {
	err = c.Call("private/get_user_trades_by_instrument_and_time", params, &result)
	return
}

func (c *DeribitWSClient) GetUserTradesByOrder(params *models.GetUserTradesByOrderParams) (result models.GetUserTradesResponse, err error) {
	err = c.Call("private/get_user_trades_by_order", params, &result)
	return
}

func (c *DeribitWSClient) GetSettlementHistoryByInstrument(params *models.GetSettlementHistoryByInstrumentParams) (result models.GetSettlementHistoryResponse, err error) {
	err = c.Call("private/get_settlement_history_by_instrument", params, &result)
	return
}

func (c *DeribitWSClient) GetSettlementHistoryByCurrency(params *models.GetSettlementHistoryByCurrencyParams) (result models.GetSettlementHistoryResponse, err error) {
	err = c.Call("private/get_settlement_history_by_currency", params, &result)
	return
}
