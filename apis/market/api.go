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

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{api.id, method, params}, resp)
}

func (api *API) GetOrderBookRaw(limit uint32) (*json.RawMessage, error) {
	var resp json.RawMessage
	if limit > 1000 {
		return nil, errors.New("golos-go: market_history_api: get_order_book -> limit must not exceed 1000")
	}
	if err := api.caller.Call("get_order_book", []interface{}{limit}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	raw, err := api.GetOrderBookRaw(limit)
	if err != nil {
		return nil, err
	}
	var resp *OrderBook
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: market_history_api: failed to unmarshal get_order_book response")
	}
	return resp, nil
}
