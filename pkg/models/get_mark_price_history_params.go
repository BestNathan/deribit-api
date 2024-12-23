package models

type GetMarkPriceHistoryParams struct {
	InstrumentName string `json:"instrument_name"`

	// Both timestamp are milliseconds since the UNIX epoch
	StartTimestamp int64 `json:"start_timestamp"`
	EndTimestamp   int64 `json:"end_timestamp"`
}
