package network_broadcast

import (
	"github.com/asuleymanov/steem-go/transports"
	"github.com/asuleymanov/steem-go/types"
)

const apiID = "network_broadcast_api"

//API plug-in structure
type API struct {
	caller transports.Caller
}

//NewAPI plug-in initialization
func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{apiID, method, params}, resp)
}

//BroadcastTransaction api request broadcast_transaction
func (api *API) BroadcastTransaction(tx *types.Transaction) error {
	return api.call("broadcast_transaction", []interface{}{tx}, nil)
}

//BroadcastTransactionSynchronous api request broadcast_transaction_synchronous
func (api *API) BroadcastTransactionSynchronous(tx *types.Transaction) (*BroadcastResponse, error) {
	var resp BroadcastResponse
	err := api.call("broadcast_transaction_synchronous", []interface{}{tx}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
