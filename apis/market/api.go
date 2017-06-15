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
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	return call.Raw(api.caller, "get_order_book", []interface{}{limit})
}

func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	var resp *OrderBook
	if err := api.caller.Call("get_order_book", []interface{}{limit}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
