package market

import (
	"github.com/asuleymanov/golos-go/types"
)

type Ticker struct {
	Latest        string `json:"latest"`
	LowestAsk     string `json:"lowest_ask"`
	HighestBid    string `json:"highest_bid"`
	PercentChange string `json:"percent_change"`
	SteemVolume   string `json:"steem_volume"`
	SbdVolume     string `json:"sbd_volume"`
}

type Volume struct {
	SteemVolume string `json:"steem_volume"`
	SbdVolume   string `json:"sbd_volume"`
}

type Trades struct {
	Date        *types.Time `json:"date"`
	CurrentPays string      `json:"current_pays"`
	OpenPays    string      `json:"open_pays"`
}

type OrderBook struct {
	Ask []*OrderBookAB `json:"asks"`
	Bid []*OrderBookAB `json:"bids"`
}

type OrderBookAB struct {
	OrderPrice *OrderPrice `json:"order_price"`
	RealPrice  string      `json:"real_price"`
	Steem      uint        `json:"steem"`
	Sbd        uint        `json:"sbd"`
	Created    *types.Time `json:"created"`
}

type OrderPrice struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
}

type MarketHistory struct {
	ID          uint32      `json:"id"`
	Open        *types.Time `json:"open"`
	Seconds     int         `json:"seconds"`
	HighSteem   int         `json:"high_steem"`
	HighSbd     int         `json:"high_sbd"`
	LowSteem    int         `json:"low_steem"`
	LowSbd      int         `json:"low_sbd"`
	OpenSteem   int         `json:"open_steem"`
	OpenSbd     int         `json:"open_sbd"`
	CloseSteem  int         `json:"close_steem"`
	CloseSbd    int         `json:"close_sbd"`
	SteemVolume int         `json:"steem_volume"`
	SbdVolume   int         `json:"sbd_volume"`
}
