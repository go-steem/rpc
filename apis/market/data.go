package market

type OrderBook struct {
	Ask []*OrderBookAB `json:"asks"`
	Bid []*OrderBookAB `json:"bids"`
}

type OrderBookAB struct {
	OrderPrice *OrderPrice `json:"order_price"`
	RealPrice  string      `json:"real_price"`
	Steem      *types.Int  `json:"steem"`
	Sbd        *types.Int  `json:"sbd"`
	Created    string      `json:"created"`
}
