package websocket

import (
	"github.com/xingxing/deribit-api/pkg/models"
)

func (c *DeribitWSClient) GetBookSummaryByCurrency(params *models.GetBookSummaryByCurrencyParams) (result []models.BookSummary, err error) {
	err = c.Call("public/get_book_summary_by_currency", params, &result)
	return
}

func (c *DeribitWSClient) GetBookSummaryByInstrument(params *models.GetBookSummaryByInstrumentParams) (result []models.BookSummary, err error) {
	err = c.Call("public/get_book_summary_by_instrument", params, &result)
	return
}

func (c *DeribitWSClient) GetContractSize(params *models.GetContractSizeParams) (result models.GetContractSizeResponse, err error) {
	err = c.Call("public/get_contract_size", params, &result)
	return
}

func (c *DeribitWSClient) GetCurrencies() (result []models.Currency, err error) {
	err = c.Call("public/get_currencies", nil, &result)
	return
}

func (c *DeribitWSClient) GetFundingChartData(params *models.GetFundingChartDataParams) (result models.GetFundingChartDataResponse, err error) {
	err = c.Call("public/get_funding_chart_data", params, &result)
	return
}

func (c *DeribitWSClient) GetHistoricalVolatility(params *models.GetHistoricalVolatilityParams) (result models.GetHistoricalVolatilityResponse, err error) {
	err = c.Call("public/get_historical_volatility", params, &result)
	return
}

func (c *DeribitWSClient) GetIndex(params *models.GetIndexParams) (result models.GetIndexResponse, err error) {
	err = c.Call("public/get_index", params, &result)
	return
}

func (c *DeribitWSClient) GetInstruments(params *models.GetInstrumentsParams) (result []models.Instrument, err error) {
	err = c.Call("public/get_instruments", params, &result)
	return
}

func (c *DeribitWSClient) GetLastSettlementsByCurrency(params *models.GetLastSettlementsByCurrencyParams) (result models.GetLastSettlementsResponse, err error) {
	err = c.Call("public/get_last_settlements_by_currency", params, &result)
	return
}

func (c *DeribitWSClient) GetLastSettlementsByInstrument(params *models.GetLastSettlementsByInstrumentParams) (result models.GetLastSettlementsResponse, err error) {
	err = c.Call("public/get_last_settlements_by_instrument", params, &result)
	return
}

func (c *DeribitWSClient) GetLastTradesByCurrency(params *models.GetLastTradesByCurrencyParams) (result models.GetLastTradesResponse, err error) {
	err = c.Call("public/get_last_trades_by_currency", params, &result)
	return
}

func (c *DeribitWSClient) GetLastTradesByCurrencyAndTime(params *models.GetLastTradesByCurrencyAndTimeParams) (result models.GetLastTradesResponse, err error) {
	err = c.Call("public/get_last_trades_by_currency_and_time", params, &result)
	return
}

func (c *DeribitWSClient) GetLastTradesByInstrument(params *models.GetLastTradesByInstrumentParams) (result models.GetLastTradesResponse, err error) {
	err = c.Call("public/get_last_trades_by_instrument", params, &result)
	return
}

func (c *DeribitWSClient) GetLastTradesByInstrumentAndTime(params *models.GetLastTradesByInstrumentAndTimeParams) (result models.GetLastTradesResponse, err error) {
	err = c.Call("public/get_last_trades_by_instrument_and_time", params, &result)
	return
}

func (c *DeribitWSClient) GetOrderBook(params *models.GetOrderBookParams) (result models.GetOrderBookResponse, err error) {
	err = c.Call("public/get_order_book", params, &result)
	return
}

func (c *DeribitWSClient) GetTradeVolumes() (result models.GetTradeVolumesResponse, err error) {
	err = c.Call("public/get_trade_volumes", nil, &result)
	return
}

func (c *DeribitWSClient) GetTradingviewChartData(params *models.GetTradingviewChartDataParams) (result models.GetTradingviewChartDataResponse, err error) {
	err = c.Call("public/get_tradingview_chart_data", params, &result)
	return
}

func (c *DeribitWSClient) Ticker(params *models.TickerParams) (result models.TickerResponse, err error) {
	err = c.Call("public/ticker", params, &result)
	return
}

func (c *DeribitWSClient) GetMarkPriceHistory(params *models.GetMarkPriceHistoryParams) (resut models.MarkPriceHistory, err error) {
	err = c.Call("public/get_mark_price_history", params, &resut)
	return
}
