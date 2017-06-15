package market

type OrderBook struct {
	Ask []*OrderBookAB `json:"asks"`
	Bid []*OrderBookAB `json:"bids"`
}

type OrderBookAB struct {
	OrderPrice *OrderPrice `json:"order_price"`
	RealPrice  string      `json:"real_price"`
	Steem      uint        `json:"steem"`
	Sbd        uint        `json:"sbd"`
	Created    string      `json:"created"`
}

type OrderPrice struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
}
