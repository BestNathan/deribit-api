package websocket

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	websocketmodels "github.com/BestNathan/deribit-api/clients/websocket/models"
	"github.com/BestNathan/deribit-api/pkg/deribit"
	"github.com/BestNathan/deribit-api/pkg/models"
	"github.com/shopspring/decimal"

	"github.com/stretchr/testify/assert"
)

func newClient() *DeribitWSClient {
	cfg := &deribit.Configuration{
		Credential: deribit.Credential{
			ApiKey:    os.Getenv("DERIBIT_KEY"),
			SecretKey: os.Getenv("DERIBIT_SECRET"),
		},
		WebsocketConfiguration: &deribit.WebsocketConfiguration{
			Url:           deribit.TestWSURL,
			AutoReconnect: true,
		},
		Debug: true,
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

	// r, err := client.Buy(&models.BuyParams{
	// 	InstrumentName: "BTC-PERPETUAL",
	// 	Amount:         decimal.NewFromInt(10),
	// 	Type:           "limit",
	// 	Price:          decimal.NewFromFloat(95683.0),
	// 	Label:          "BestNathan-test",
	// })

	// t.Logf("%#v %v", r, err)

	r, err := client.GetMarkPriceHistory(&models.GetMarkPriceHistoryParams{
		InstrumentName: "BTC-PERPETUAL",
		StartTimestamp: time.Now().Add(-10 * time.Minute).UnixMilli(),
		EndTimestamp:   time.Now().UnixMilli(),
	})

	t.Logf("%#v %v", r, err)

	// r, err := client.Buy(&models.BuyParams{
	// 	InstrumentName: "BTC-PERPETUAL",
	// 	Amount:         decimal.NewFromInt(10),
	// 	Type:           "limit",
	// 	Price:          decimal.NewFromFloat(95683.0),
	// 	Label:          "BestNathan",
	// })

	// t.Logf("%#v %v", r, err)

	// sr, err := client.CancellByLabel(&models.CancelByLabelParams{
	// 	Label:    "BestNathan",
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
