package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//AccountWitnessProxyOperation represents account_witness_proxy operation data.
type AccountWitnessProxyOperation struct {
	Account string `json:"account"`
	Proxy   string `json:"proxy"`
}

//Type function that defines the type of operation AccountWitnessProxyOperation.
func (op *AccountWitnessProxyOperation) Type() OpType {
	return TypeAccountWitnessProxy
}

//Data returns the operation data AccountWitnessProxyOperation.
func (op *AccountWitnessProxyOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AccountWitnessProxyOperation to bytes.
func (op *AccountWitnessProxyOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessProxy.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Proxy)
	return enc.Err()
}
