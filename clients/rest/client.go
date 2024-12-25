package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	restmodels "github.com/BestNathan/deribit-api/clients/rest/models"
	"github.com/BestNathan/deribit-api/pkg/deribit"
	"github.com/BestNathan/deribit-api/pkg/models"
	"github.com/sirupsen/logrus"

	"net/http"

	"github.com/shopspring/decimal"
)

const (
	StdDepth    = 10
	BtcTickSize = 0.5
)

type DeribitRestClient struct {
	Client      *http.Client
	ClientID    string
	ApiSecret   string
	BaseURL     string
	AccessToken *string
	Logger      *logrus.Logger
}

func NewDeribitRestClient(cfg *deribit.Configuration) *DeribitRestClient {
	return &DeribitRestClient{
		Client:      &http.Client{},
		ClientID:    cfg.ApiKey,
		ApiSecret:   cfg.SecretKey,
		BaseURL:     cfg.RestAddr,
		AccessToken: nil,
		Logger:      cfg.Logger,
	}
}

func (d *DeribitRestClient) GetAuthToken() (string, error) {
	baseURL := strings.Replace(d.BaseURL, "/ws", "", 1)
	d.Logger.Debugf("Original Base URL: %s", baseURL)

	baseURL = strings.TrimSuffix(baseURL, "/")
	d.Logger.Debugf("URL after trim: %s", baseURL)

	authURL := baseURL + "/public/auth"
	d.Logger.Debugf("Final auth URL: %s", authURL)

	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", d.ClientID)
	params.Set("client_secret", d.ApiSecret)

	fullURL := authURL + "?" + params.Encode()
	d.Logger.Debugf("Full URL with params: %s", fullURL)

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := d.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.Logger.Errorf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	d.Logger.Debugf("Response status: %s", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	d.Logger.Debugf("Response body: %s", string(body))

	var result struct {
		Result struct {
			AccessToken string `json:"access_token"`
		} `json:"result"`
		Error *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if result.Error != nil {
		return "", fmt.Errorf("API error %d: %s", result.Error.Code, result.Error.Message)
	}

	d.AccessToken = &result.Result.AccessToken
	return *d.AccessToken, nil
}

func (d *DeribitRestClient) request(method string, params map[string]interface{}, private bool) (map[string]interface{}, error) {
	baseURL := strings.Replace(d.BaseURL, "/ws", "", 1)
	baseURL = strings.TrimSuffix(baseURL, "/")
	fullURL := baseURL + "/" + method

	d.Logger.Debugf("Request URL: %s", fullURL)
	d.Logger.Debugf("Request params: %+v", params)

	urlValues := url.Values{}
	for key, value := range params {
		urlValues.Set(key, fmt.Sprintf("%v", value))
	}

	fullURL = fullURL + "?" + urlValues.Encode()

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if private && d.AccessToken != nil {
		req.Header.Set("Authorization", "Bearer "+*d.AccessToken)
	}

	resp, err := d.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.Logger.Errorf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	d.Logger.Debugf("Response body: %s", string(body))

	var result struct {
		JSONRPC string                 `json:"jsonrpc"`
		ID      int                    `json:"id"`
		Result  map[string]interface{} `json:"result"`
		Error   *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("API error %d: %s", result.Error.Code, result.Error.Message)
	}

	return result.Result, nil
}

// requestInterface makes a request to the Deribit API and returns the result as interface{}
func (d *DeribitRestClient) requestInterface(method string, params map[string]interface{}, private bool) (interface{}, error) {
	baseURL := strings.Replace(d.BaseURL, "/ws", "", 1)
	baseURL = strings.TrimSuffix(baseURL, "/")
	fullURL := baseURL + "/" + method

	d.Logger.Debugf("Request URL: %s", fullURL)
	d.Logger.Debugf("Request params: %+v", params)

	urlValues := url.Values{}
	for key, value := range params {
		urlValues.Set(key, fmt.Sprintf("%v", value))
	}

	fullURL = fullURL + "?" + urlValues.Encode()

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if private && d.AccessToken != nil {
		req.Header.Set("Authorization", "Bearer "+*d.AccessToken)
	}

	resp, err := d.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.Logger.Errorf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	d.Logger.Debugf("Response body: %s", string(body))

	var result struct {
		JSONRPC string      `json:"jsonrpc"`
		ID      int         `json:"id"`
		Result  interface{} `json:"result"`
		Error   *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("API error %d: %s", result.Error.Code, result.Error.Message)
	}

	return result.Result, nil
}

func (d *DeribitRestClient) GetOrderbook(instrument string, depth *int) (restmodels.OrderBook, error) {
	depthValue := StdDepth
	if depth != nil {
		depthValue = *depth
	}

	params := map[string]interface{}{
		"instrument_name": instrument,
		"depth":           depthValue,
	}

	d.Logger.Debugf("Getting orderbook for %s with depth %d", instrument, depthValue)

	result, err := d.request("public/get_order_book", params, false)
	if err != nil {
		return restmodels.OrderBook{}, fmt.Errorf("failed to get orderbook: %w", err)
	}

	var orderBook restmodels.OrderBook
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return restmodels.OrderBook{}, fmt.Errorf("failed to marshal result: %w", err)
	}

	d.Logger.Debugf("Orderbook raw data: %s", string(resultBytes))

	if err := json.Unmarshal(resultBytes, &orderBook); err != nil {
		return restmodels.OrderBook{}, fmt.Errorf("failed to unmarshal order book: %w", err)
	}

	d.Logger.Debugf("Parsed orderbook: %+v", orderBook)
	return orderBook, nil
}

func (d *DeribitRestClient) GetTicker(instrument string) (decimal.Decimal, error) {
	params := map[string]interface{}{
		"instrument_name": instrument,
	}

	response, err := d.request("public/ticker", params, false)
	if err != nil {
		return decimal.Zero, err
	}

	lastPrice, exists := response["last_price"]
	if !exists {
		return decimal.Zero, errors.New("missing last_price field in response")
	}

	switch v := lastPrice.(type) {
	case float64:
		return decimal.NewFromFloat(v), nil
	case string:
		return decimal.NewFromString(v)
	default:
		return decimal.Zero, errors.New("last price is neither a number nor a string")
	}
}

func (d *DeribitRestClient) PlaceLimitOrder(instrument string,
	price decimal.Decimal,
	amount decimal.Decimal,
	direction restmodels.Direction) (restmodels.Order, error) {
	priceToTick := roundToTickSize(price, decimal.NewFromFloat(BtcTickSize))
	method := "private/sell"
	if direction == "buy" {
		method = "private/buy"
	}

	params := map[string]interface{}{
		"instrument_name": instrument,
		"price":           priceToTick.String(),
		"amount":          amount.String(),
		"type":            models.OrderTypeLimit,
		"post_only":       true,
	}

	orderResult, err := d.request(method, params, true)
	if err != nil {
		return restmodels.Order{}, err
	}

	order := restmodels.Order{}
	orderData, _ := json.Marshal(orderResult["order"])
	if err := json.Unmarshal(orderData, &order); err != nil {
		return restmodels.Order{}, err
	}

	return order, nil
}

func (d *DeribitRestClient) GetRecentTrades(instrument string, count int) ([]models.Trade, error) {
	params := map[string]interface{}{
		"instrument_name": instrument,
		"count":           count,
	}

	d.Logger.Debugf("Getting recent trades for %s with count %d", instrument, count)

	result, err := d.request("public/get_last_trades_by_instrument", params, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get trades: %w", err)
	}

	trades, ok := result["trades"].([]interface{})
	if !ok {
		return nil, errors.New("invalid trades data format")
	}

	tradesData, err := json.Marshal(trades)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal trades: %w", err)
	}

	d.Logger.Debugf("Trades raw data: %s", string(tradesData))

	var tradesList []models.Trade
	if err := json.Unmarshal(tradesData, &tradesList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal trades: %w", err)
	}

	d.Logger.Debugf("Parsed trades: %+v", tradesList)
	return tradesList, nil
}

// GetFundingRate retrieves the funding rate for a perpetual instrument
func (d *DeribitRestClient) GetFundingRate(instrument string, startTime, endTime time.Time) (models.FundingRatePoint, error) {
	params := map[string]interface{}{
		"instrument_name": instrument,
		"start_timestamp": startTime.UnixNano() / 1e6,
		"end_timestamp":   endTime.UnixNano() / 1e6,
	}

	d.Logger.Debugf("Getting funding rate for %s between %v and %v", instrument, startTime, endTime)

	result, err := d.requestInterface("public/get_funding_rate_value", params, false)
	if err != nil {
		return models.FundingRatePoint{}, fmt.Errorf("failed to get funding rate: %w", err)
	}

	rate, ok := result.(float64)
	if !ok {
		return models.FundingRatePoint{}, fmt.Errorf("unexpected funding rate format in response: %T", result)
	}

	fundingRate := models.FundingRatePoint{
		Timestamp: startTime.Unix(),
		Rate:      rate,
	}

	d.Logger.Debugf("Parsed funding rate: %+v", fundingRate)

	return fundingRate, nil
}

// GetCurrentFundingRate is a convenience method that gets the most recent funding rate
func (d *DeribitRestClient) GetCurrentFundingRate(instrument string) (models.FundingRatePoint, error) {
	now := time.Now()
	// Get the rate for the last hour
	startTime := now.Add(-1 * time.Hour)

	return d.GetFundingRate(instrument, startTime, now)
}

// GetBookSummary retrieves the book summary for a specific instrument
func (d *DeribitRestClient) GetBookSummary(instrument string) ([]models.BookSummary, error) {
	params := map[string]interface{}{
		"instrument_name": instrument,
	}

	d.Logger.Debugf("Getting book summary for %s", instrument)

	result, err := d.requestInterface("public/get_book_summary_by_instrument", params, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get book summary: %w", err)
	}

	// The result is an array of book summaries
	summariesData, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format: %T", result)
	}

	// Convert to JSON to properly unmarshal into BookSummary structs
	jsonData, err := json.Marshal(summariesData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal book summaries: %w", err)
	}

	var summaries []models.BookSummary
	if err := json.Unmarshal(jsonData, &summaries); err != nil {
		return nil, fmt.Errorf("failed to unmarshal book summaries: %w", err)
	}

	d.Logger.Debugf("Retrieved %d book summaries for %s", len(summaries), instrument)
	return summaries, nil
}
