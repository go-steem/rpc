package market

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/golos-go/interfaces"
	"github.com/asuleymanov/golos-go/internal/rpc"

	// Vendor
	"github.com/pkg/errors"
)

const APIID = "market_history_api"

var EmptyParams = []string{}

type API struct {
	id     int
	caller interfaces.Caller
}

func NewAPI(caller interfaces.Caller) (*API, error) {
	id, err := rpc.GetNumericAPIID(caller, APIID)
	if err != nil {
		return nil, err
	}
	return &API{id, caller}, nil
}

func (api *API) Raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{api.id, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to call %v\n", APIID, method)
	}
	return &resp, nil
}

//get_ticker
func (api *API) GetTicker() (*Ticker, error) {
	raw, err := api.Raw("get_ticker", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp *Ticker
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: market_history_api: failed to unmarshal get_ticker response")
	}
	return resp, nil
}

//get_volume
func (api *API) GetVolume() (*Volume, error) {
	raw, err := api.Raw("get_volume", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp *Volume
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: market_history_api: failed to unmarshal get_volume response")
	}
	return resp, nil
}

//get_order_book
func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	if limit > 1000 {
		return nil, errors.New("golos-go: market_history_api: get_order_book -> limit must not exceed 1000")
	}
	raw, err := api.Raw("get_order_book", []interface{}{limit})
	if err != nil {
		return nil, err
	}
	var resp *OrderBook
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: market_history_api: failed to unmarshal get_order_book response")
	}
	return resp, nil
}

//get_trade_history
func (api *API) GetTradeHistory(start, end string, limit uint32) ([]*Trades, error) {
	if limit > 1000 {
		return nil, errors.New("golos-go: market_history_api: get_order_book -> limit must not exceed 1000")
	}
	raw, err := api.Raw("get_trade_history", []interface{}{start, end, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Trades
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: market_history_api: failed to unmarshal get_trade_history response")
	}
	return resp, nil
}

//get_recent_trades
func (api *API) GetRecentTrades(limit uint32) ([]*Trades, error) {
	if limit > 1000 {
		return nil, errors.New("golos-go: market_history_api: get_order_book -> limit must not exceed 1000")
	}
	raw, err := api.Raw("get_recent_trades", []interface{}{limit})
	if err != nil {
		return nil, err
	}
	var resp []*Trades
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: market_history_api: failed to unmarshal get_recent_trades response")
	}
	return resp, nil
}

//get_market_history
func (api *API) GetMarketHistory(b_sec uint32, start, end string) ([]*MarketHistory, error) {
	raw, err := api.Raw("get_market_history", []interface{}{b_sec, start, end})
	if err != nil {
		return nil, err
	}
	var resp []*MarketHistory
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: market_history_api: failed to unmarshal get_market_history response")
	}
	return resp, nil
}

//get_market_history_buckets
func (api *API) GetMarketHistoryBuckets() ([]uint32, error) {
	raw, err := api.Raw("get_market_history_buckets", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp []uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: market_history_api: failed to unmarshal get_market_history_buckets response")
	}
	return resp, nil
}
