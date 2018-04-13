package market

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/rpc/transports"

	// Vendor
	"github.com/pkg/errors"
)

const apiID = "market_history_api"

type API struct {
	caller transports.Caller
}

func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var emptyParams = struct{}{}

func (api *API) raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{apiID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "steem-go: %v: failed to call %v\n", apiID, method)
	}
	return &resp, nil
}

//get_ticker
func (api *API) GetTicker() (*Ticker, error) {
	raw, err := api.raw("get_ticker", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp *Ticker
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: market_history_api: failed to unmarshal get_ticker response")
	}
	return resp, nil
}

//get_volume
func (api *API) GetVolume() (*Volume, error) {
	raw, err := api.raw("get_volume", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp *Volume
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: market_history_api: failed to unmarshal get_volume response")
	}
	return resp, nil
}

//get_order_book
func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	if limit > 1000 {
		return nil, errors.New("steem-go: market_history_api: get_order_book -> limit must not exceed 1000")
	}
	raw, err := api.raw("get_order_book", []interface{}{limit})
	if err != nil {
		return nil, err
	}
	var resp *OrderBook
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: market_history_api: failed to unmarshal get_order_book response")
	}
	return resp, nil
}

//get_trade_history
func (api *API) GetTradeHistory(start, end string, limit uint32) ([]*Trades, error) {
	if limit > 1000 {
		return nil, errors.New("steem-go: market_history_api: get_order_book -> limit must not exceed 1000")
	}
	raw, err := api.raw("get_trade_history", []interface{}{start, end, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Trades
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: market_history_api: failed to unmarshal get_trade_history response")
	}
	return resp, nil
}

//get_recent_trades
func (api *API) GetRecentTrades(limit uint32) ([]*Trades, error) {
	if limit > 1000 {
		return nil, errors.New("steem-go: market_history_api: get_order_book -> limit must not exceed 1000")
	}
	raw, err := api.raw("get_recent_trades", []interface{}{limit})
	if err != nil {
		return nil, err
	}
	var resp []*Trades
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: market_history_api: failed to unmarshal get_recent_trades response")
	}
	return resp, nil
}

//get_market_history
func (api *API) GetMarketHistory(b_sec uint32, start, end string) ([]*MarketHistory, error) {
	raw, err := api.raw("get_market_history", []interface{}{b_sec, start, end})
	if err != nil {
		return nil, err
	}
	var resp []*MarketHistory
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: market_history_api: failed to unmarshal get_market_history response")
	}
	return resp, nil
}

//get_market_history_buckets
func (api *API) GetMarketHistoryBuckets() ([]uint32, error) {
	raw, err := api.raw("get_market_history_buckets", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp []uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: market_history_api: failed to unmarshal get_market_history_buckets response")
	}
	return resp, nil
}
