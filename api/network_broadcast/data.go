package network_broadcast

//BroadcastResponse structure for the BroadcastTransactionSynchronous function
type BroadcastResponse struct {
	ID       string `json:"id"`
	BlockNum uint32 `json:"block_num"`
	TrxNum   uint32 `json:"trx_num"`
	Expired  bool   `json:"expired"`
}
