package websocket

import "fmt"

const (
	CHANNEL_USER_ACCESS_LOG                  = "user.access_log"
	CHANNEL_USER_MMP_TRIGGER_PATTERN         = "user.mmp_trigger.%s"         //user.mmp_trigger.{index_name}
	CHANNEL_BOOK_GROUP_PATTERN               = "book.%s.%s.%d.%s"            // book.{instrument_name}.{group}.{depth}.{interval}
	CHANNEL_DERIBIT_VOLATILITY_INDEX_PATTERN = "deribit_volatility_index.%s" // deribit_volatility_index.{index_name}
)

const (
	DERIBIT_VOLATILITY_INDEX_NAME_BTC = "btc_usd"
	DERIBIT_VOLATILITY_INDEX_NAME_ETH = "eth_usd"
)

func ChannelUserMMPTrigger(idxname string) string {
	if idxname == "" {
		return ""
	}
	return fmt.Sprintf(CHANNEL_USER_MMP_TRIGGER_PATTERN, idxname)
}

func ChannelBookGroup(instrumentname, group string, depth int, interval string) string {
	if instrumentname == "" {
		return ""
	}

	if group == "" {
		group = "none"
	}

	if depth == 0 {
		depth = 1
	}

	if interval == "" {
		interval = "100ms"
	}

	return fmt.Sprintf(CHANNEL_BOOK_GROUP_PATTERN, instrumentname, group, depth, interval)
}

func ChannelDeribitVolatilityIndex(idxname string) string {
	if idxname == "" {
		return ""
	}

	return fmt.Sprintf(CHANNEL_DERIBIT_VOLATILITY_INDEX_PATTERN, idxname)
}
