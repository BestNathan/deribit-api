package main

import (
	"github.com/BestNathan/deribit-api/clients/rest"
	"github.com/BestNathan/deribit-api/pkg/deribit"
)

func main() {

	cfg := deribit.GetConfig()
	client := rest.NewDeribitRestClient(cfg)

	_, err := client.GetAuthToken()
	if err != nil {
		cfg.Logger.Errorf("Failed to get auth token: %v", err)
		return
	}

	//depth := 5
	//orderbook, err := client.GetOrderbook("BTC-PERPETUAL", &depth)
	//if err != nil {
	//	cfg.Logger.Errorf("Failed to get orderbook: %v", err)
	//	return
	//}
	//cfg.Logger.Infof("Orderbook: %v", orderbook)

	// Get Funding Rate
	//fundingRate, err := client.GetCurrentFundingRate("BTC-PERPETUAL")
	//if err != nil {
	//	cfg.Logger.Errorf("Failed to get funding rate: %v", err)
	//	return
	//} else {
	//	cfg.Logger.Infof("Funding Rate: %v", fundingRate)
	//}

	// Get Trades
	//trades, err := client.GetRecentTrades("BTC-PERPETUAL", 10)
	//if err != nil {
	//	cfg.Logger.Errorf("Failed to get trades: %v", err)
	//	return
	//} else {
	//	cfg.Logger.Infof("Trades: %v", trades)
	//}

	// Get GetBookSummaryByInstrument
	summary, err := client.GetBookSummary("BTC-PERPETUAL")
	if err != nil {
		cfg.Logger.Errorf("Failed to get trades: %v", err)
		return
	} else {
		cfg.Logger.Infof("Summary: %v", summary)
	}
}
