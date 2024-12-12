package websocket

import (
	"encoding/json"
	websocketmodels "github.com/xingxing/deribit-api/clients/websocket/models"
	"github.com/xingxing/deribit-api/pkg/deribit"
	"github.com/xingxing/deribit-api/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newClient() *DeribitWSClient {
	cfg := &deribit.Configuration{
		WsAddr:        deribit.RealBaseURL,
		ApiKey:        "AsJTU16U",
		SecretKey:     "mM5_K8LVxztN6TjjYpv_cJVGQBvk4jglrEpqkw1b87U",
		AutoReconnect: true,
		DebugMode:     true,
	}
	client := NewDeribitWsClient(cfg)
	return client
}

func TestClient_GetTime(t *testing.T) {
	client := newClient()
	tm, err := client.GetTime()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", tm)
}

func TestClient_Test(t *testing.T) {
	client := newClient()
	result, err := client.Test()
	assert.Nil(t, err)
	t.Logf("%v", result)
}

func TestClient_GetBookSummaryByCurrency(t *testing.T) {
	client := newClient()
	params := &models.GetBookSummaryByCurrencyParams{
		Currency: "BTC",
		Kind:     "future",
	}
	result, err := client.GetBookSummaryByCurrency(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", result)
}

func TestClient_GetBookSummaryByInstrument(t *testing.T) {
	client := newClient()
	params := &models.GetBookSummaryByInstrumentParams{
		InstrumentName: "BTC-PERPETUAL",
	}
	result, err := client.GetBookSummaryByInstrument(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", result)
}

func TestClient_GetOrderBook(t *testing.T) {
	client := newClient()
	params := &models.GetOrderBookParams{
		InstrumentName: "BTC-PERPETUAL",
		Depth:          5,
	}
	result, err := client.GetOrderBook(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", result)
}

func TestClient_Ticker(t *testing.T) {
	client := newClient()
	params := &models.TickerParams{
		InstrumentName: "BTC-PERPETUAL",
	}
	result, err := client.Ticker(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", result)
}

func TestClient_GetPosition(t *testing.T) {
	client := newClient()
	params := &models.GetPositionParams{
		InstrumentName: "BTC-PERPETUAL",
	}
	result, err := client.GetPosition(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", result)
}

func TestClient_BuyMarket(t *testing.T) {
	client := newClient()
	params := &models.BuyParams{
		InstrumentName: "BTC-PERPETUAL",
		Amount:         10,
		Price:          0,
		Type:           "market",
	}
	result, err := client.Buy(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", result)
}

func TestClient_Buy(t *testing.T) {
	client := newClient()
	params := &models.BuyParams{
		InstrumentName: "BTC-PERPETUAL",
		Amount:         40,
		Price:          6000.0,
		Type:           "limit",
	}
	result, err := client.Buy(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", result)
}

func TestJsonOmitempty(t *testing.T) {
	params := &models.BuyParams{
		InstrumentName: "BTC-PERPETUAL",
		Amount:         40,
		//Price:          6000.0,
		Type:        "limit",
		TimeInForce: "good_til_cancelled",
		MaxShow:     websocketmodels.Float64Pointer(40.0),
	}
	data, _ := json.Marshal(params)
	t.Log(string(data))
}

func TestOnBook(t *testing.T) {
	client := newClient()
	client.On("book.BTC-PERPETUAL.none.1.100ms", func(e *models.OrderBookGroupNotification) {
		t.Logf("bids: %s\tasks: %s", e.Bids, e.Asks)
	})

	client.Subscribe([]string{"book.BTC-PERPETUAL.none.1.100ms"})

	select {}
}
