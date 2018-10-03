package market_history

import (
	"fmt"

	"github.com/asuleymanov/steem-go/transports"
)

const apiID = "market_history_api"

//API plug-in structure
type API struct {
	caller transports.Caller
}

//NewAPI plug-in initialization
func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var emptyParams = []struct{}{}

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{apiID, method, params}, resp)
}

//GetTicker api request get_ticker
func (api *API) GetTicker() (*Ticker, error) {
	var resp Ticker
	err := api.call("get_ticker", emptyParams, &resp)
	return &resp, err
}

//GetVolume api request get_volume
func (api *API) GetVolume() (*Volume, error) {
	var resp Volume
	err := api.call("get_volume", emptyParams, &resp)
	return &resp, err
}

//GetOrderBook api request get_order_book
func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("%v: get_order_book -> limit must not exceed 1000", apiID)
	}
	var resp OrderBook
	err := api.call("get_order_book", []interface{}{limit}, &resp)
	return &resp, err
}

//GetTradeHistory api request get_trade_history
func (api *API) GetTradeHistory(start, end string, limit uint32) ([]*Trades, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("%v: get_order_book -> limit must not exceed 1000", apiID)
	}
	var resp []*Trades
	err := api.call("get_trade_history", []interface{}{start, end, limit}, &resp)
	return resp, err
}

//GetRecentTrades api request get_recent_trades
func (api *API) GetRecentTrades(limit uint32) ([]*Trades, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("%v: get_order_book -> limit must not exceed 1000", apiID)
	}
	var resp []*Trades
	err := api.call("get_recent_trades", []interface{}{limit}, &resp)
	return resp, err
}

//GetMarketHistory api request get_market_history
func (api *API) GetMarketHistory(bSec uint32, start, end string) ([]*MarketHistory, error) {
	var resp []*MarketHistory
	err := api.call("get_market_history", []interface{}{bSec, start, end}, &resp)
	return resp, err
}

//GetMarketHistoryBuckets api request get_market_history_buckets
func (api *API) GetMarketHistoryBuckets() ([]*uint32, error) {
	var resp []*uint32
	err := api.call("get_market_history_buckets", emptyParams, &resp)
	return resp, err
}
