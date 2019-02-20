package api

import (
	"github.com/asuleymanov/steem-go/types"
)

//network_broadcast_api

//BroadcastTransaction api request broadcast_transaction
func (api *API) BroadcastTransaction(tx *types.Transaction) error {
	return api.call("network_broadcast_api", "broadcast_transaction", []interface{}{tx}, nil)
}

//BroadcastTransactionSynchronous api request broadcast_transaction_synchronous
func (api *API) BroadcastTransactionSynchronous(tx *types.Transaction) (*BroadcastResponse, error) {
	var resp BroadcastResponse
	err := api.call("network_broadcast_api", "broadcast_transaction_synchronous", []interface{}{tx}, &resp)
	return &resp, err
}
