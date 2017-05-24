package networkbroadcast

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/golos-go/interfaces"
	"github.com/asuleymanov/golos-go/internal/rpc"
	"github.com/asuleymanov/golos-go/types"

	// Vendor
	"github.com/pkg/errors"
)

const APIID = "network_broadcast_api"

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

/*
 * broadcast_transaction
 */

func (api *API) BroadcastTransaction(tx *types.Transaction) error {
	params := []interface{}{tx}
	return api.call("broadcast_transaction", params, nil)
}

/*
 * broadcast_transaction_synchronous
 */

func (api *API) BroadcastTransactionSynchronousRaw(tx *types.Transaction) (*json.RawMessage, error) {
	params := []interface{}{tx}

	var resp json.RawMessage
	if err := api.call("broadcast_transaction_synchronous", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type BroadcastResponse struct {
	ID       string `json:"id"`
	BlockNum uint32 `json:"block_num"`
	TrxNum   uint32 `json:"trx_num"`
	Expired  bool   `json:"expired"`
}

func (api *API) BroadcastTransactionSynchronous(tx *types.Transaction) (*BroadcastResponse, error) {
	raw, err := api.BroadcastTransactionSynchronousRaw(tx)
	if err != nil {
		return nil, err
	}

	var resp BroadcastResponse
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal BroadcastResponse: %v", string(*raw))
	}
	return &resp, nil
}
