package websocket

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/shopspring/decimal"
	websocketmodels "github.com/xingxing/deribit-api/clients/websocket/models"
	"github.com/xingxing/deribit-api/pkg/deribit"
	"github.com/xingxing/deribit-api/pkg/models"

	"github.com/stretchr/testify/assert"
)

func newClient() *DeribitWSClient {
	cfg := &deribit.Configuration{
		WsAddr:        deribit.RealBaseURL,
		ApiKey:        os.Getenv("DERIBIT_KEY"),
		SecretKey:     os.Getenv("DERIBIT_SECRET"),
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
		Amount:         decimal.NewFromInt(10),
		Price:          decimal.NewFromInt(0),
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
		Amount:         decimal.NewFromInt(40),
		Price:          decimal.NewFromInt(6000),
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
		Amount:         decimal.NewFromInt(40),
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

	r, err := client.Buy(&models.BuyParams{
		InstrumentName: "BTC-PERPETUAL",
		Amount:         decimal.NewFromInt(10),
		Type:           "limit",
		Price:          decimal.NewFromFloat(95683.0),
		Label:          "xingxing-test",
	})

	t.Logf("%#v %v", r, err)

	// r, err := client.Buy(&models.BuyParams{
	// 	InstrumentName: "BTC-PERPETUAL",
	// 	Amount:         decimal.NewFromInt(10),
	// 	Type:           "limit",
	// 	Price:          decimal.NewFromFloat(95683.0),
	// 	Label:          "xingxing",
	// })

	// t.Logf("%#v %v", r, err)

	// sr, err := client.CancellByLabel(&models.CancelByLabelParams{
	// 	Label:    "xingxing",
	// 	Currency: "BTC",
	// })

	// t.Logf("%#v %v", sr, err)

	// r, _ := client.Sell(&models.SellParams{
	// 	InstrumentName: "BTC-PERPETUAL",
	// 	Amount:         10,
	// 	Type:           "market",
	// })

	// t.Logf("%#v", r)

	// Only for futures, position size in base currency
	// r, err := client.GetPositions(&models.GetPositionsParams{
	// 	Currency: "BTC",
	// 	Kind:     "future",
	// })

	// t.Logf("currency size %s \n %v", r[0].SizeCurrency, err)

	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// t.Logf("%#v", p)

	// r, _ := client.GetOrderBook(&models.GetOrderBookParams{
	// 	InstrumentName: "BTC-PERPETUAL",
	// 	Depth:          1,
	// })

	// t.Logf("%s", r.BestBidPrice)

	// r, _ = client.GetOpenOrdersByInstrument(&models.GetOpenOrdersByInstrumentParams{
	// 	InstrumentName: "BTC-PERPETUAL",
	// 	Type:           "all",
	// })

	// for _, v := range r {
	// 	t.Logf("%#v", v)
	// }
}
