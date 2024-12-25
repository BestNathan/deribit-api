package rest

import (
	"encoding/json"
	"github.com/BestNathan/deribit-api/pkg/deribit"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewDeribitRestClient(t *testing.T) {
	logger := logrus.New()
	cfg := &deribit.Configuration{
		ApiKey:    "test-key",
		SecretKey: "test-secret",
		RestAddr:  "http://test.deribit.com",
		Logger:    logger,
	}

	client := NewDeribitRestClient(cfg)

	assert.NotNil(t, client)
	assert.Equal(t, cfg.ApiKey, client.ClientID)
	assert.Equal(t, cfg.SecretKey, client.ApiSecret)
	assert.Equal(t, cfg.RestAddr, client.BaseURL)
	assert.Equal(t, logger, client.Logger)
}

func TestGetAuthToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/public/auth", r.URL.Path)

		response := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result": map[string]interface{}{
				"access_token": "test-token",
			},
		}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}))
	defer server.Close()

	client := &DeribitRestClient{
		Client:    http.DefaultClient,
		ClientID:  "test-key",
		ApiSecret: "test-secret",
		BaseURL:   server.URL,
		Logger:    logrus.New(),
	}

	token, err := client.GetAuthToken()
	assert.NoError(t, err)
	assert.Equal(t, "test-token", token)
}

func TestGetOrderbook(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/public/get_order_book", r.URL.Path)

		response := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result": map[string]interface{}{
				"asks":      [][]float64{{9000.5, 1.0}},
				"bids":      [][]float64{{8999.5, 1.0}},
				"timestamp": 1234567890,
			},
		}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}))
	defer server.Close()

	client := &DeribitRestClient{
		Client:  http.DefaultClient,
		BaseURL: server.URL,
		Logger:  logrus.New(),
	}

	orderbook, err := client.GetOrderbook("BTC-PERPETUAL", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, orderbook.Asks)
	assert.NotEmpty(t, orderbook.Bids)
}

func TestPlaceLimitOrder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "/private/")

		response := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result": map[string]interface{}{
				"order": map[string]interface{}{
					"order_id":  "test_order_id",
					"price":     9000.5,
					"amount":    1.0,
					"direction": "buy",
				},
			},
		}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}))
	defer server.Close()

	client := &DeribitRestClient{
		Client:      http.DefaultClient,
		BaseURL:     server.URL,
		Logger:      logrus.New(),
		AccessToken: stringPtr("test-token"),
	}

	price := decimal.NewFromFloat(9000.5)
	amount := decimal.NewFromFloat(1.0)

	order, err := client.PlaceLimitOrder("BTC-PERPETUAL", price, amount, "buy")
	assert.NoError(t, err)
	assert.Equal(t, "test_order_id", order.OrderID)
}

func TestGetRecentTrades(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/public/get_last_trades_by_instrument", r.URL.Path)

		response := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result": map[string]interface{}{
				"trades": []map[string]interface{}{
					{
						"trade_id":  "test_trade_id",
						"price":     9000.5,
						"amount":    1.0,
						"direction": "buy",
						"timestamp": 1234567890000,
					},
				},
			},
		}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}))
	defer server.Close()

	client := &DeribitRestClient{
		Client:  http.DefaultClient,
		BaseURL: server.URL,
		Logger:  logrus.New(),
	}

	trades, err := client.GetRecentTrades("BTC-PERPETUAL", 1)
	assert.NoError(t, err)
	assert.Len(t, trades, 1)
	assert.Equal(t, "test_trade_id", trades[0].TradeID)
}

func TestGetFundingRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/public/get_funding_rate_value", r.URL.Path)

		response := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result":  0.0001,
		}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}))
	defer server.Close()

	client := &DeribitRestClient{
		Client:  http.DefaultClient,
		BaseURL: server.URL,
		Logger:  logrus.New(),
	}

	startTime := time.Now().Add(-1 * time.Hour)
	endTime := time.Now()

	rate, err := client.GetFundingRate("BTC-PERPETUAL", startTime, endTime)
	assert.NoError(t, err)
	assert.Equal(t, 0.0001, rate.Rate)
}

func TestGetCurrentFundingRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/public/get_funding_rate_value", r.URL.Path)

		response := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result":  0.0001,
		}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}))
	defer server.Close()

	client := &DeribitRestClient{
		Client:  http.DefaultClient,
		BaseURL: server.URL,
		Logger:  logrus.New(),
	}

	rate, err := client.GetCurrentFundingRate("BTC-PERPETUAL")
	assert.NoError(t, err)
	assert.Equal(t, 0.0001, rate.Rate)
}

func stringPtr(s string) *string {
	return &s
}

func TestGetBookSummaryByInstrument(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/public/get_book_summary_by_instrument", r.URL.Path)
		assert.Equal(t, "ETH-22FEB19-140-P", r.URL.Query().Get("instrument_name"))

		response := `{
            "jsonrpc": "2.0",
            "id": 3659,
            "result": [{
                "volume": 0.55,
                "underlying_price": 121.38,
                "underlying_index": "index_price",
                "quote_currency": "USD",
                "price_change": -26.7793594,
                "open_interest": 0.55,
                "mid_price": 0.2444,
                "mark_price": 0.179112,
                "low": 0.34,
                "last": 0.34,
                "interest_rate": 0.207,
                "instrument_name": "ETH-22FEB19-140-P",
                "high": 0.34,
                "creation_timestamp": 1550227952163,
                "bid_price": 0.1488,
                "base_currency": "ETH",
                "ask_price": 0.34
            }]
        }`
		_, err := w.Write([]byte(response))
		if err != nil {
			return
		}
	}))
	defer server.Close()

	client := &DeribitRestClient{
		Client:  http.DefaultClient,
		BaseURL: server.URL,
		Logger:  logrus.New(),
	}

	summaries, err := client.GetBookSummary("ETH-22FEB19-140-P")
	assert.NoError(t, err)
	assert.NotNil(t, summaries)
	assert.Len(t, summaries, 1)

	summary := summaries[0]

	assert.Equal(t, 0.55, summary.Volume)
	assert.Equal(t, 121.38, summary.UnderlyingPrice)
	assert.Equal(t, "index_price", summary.UnderlyingIndex)
	assert.Equal(t, "USD", summary.QuoteCurrency)
	assert.Equal(t, 0.55, summary.OpenInterest)
	assert.Equal(t, 0.2444, summary.MidPrice)
	assert.Equal(t, 0.179112, summary.MarkPrice)
	assert.Equal(t, 0.34, summary.Low)
	assert.Equal(t, 0.34, summary.Last)
	assert.Equal(t, 0.207, summary.InterestRate)
	assert.Equal(t, "ETH-22FEB19-140-P", summary.InstrumentName)
	assert.Equal(t, 0.34, summary.High)
	assert.Equal(t, int64(1550227952163), summary.CreationTimestamp)
	assert.Equal(t, 0.1488, summary.BidPrice)
	assert.Equal(t, "ETH", summary.BaseCurrency)
	assert.Equal(t, 0.34, summary.AskPrice)
}
