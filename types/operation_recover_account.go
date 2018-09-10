package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//RecoverAccountOperation represents recover_account operation data.
type RecoverAccountOperation struct {
	AccountToRecover     string        `json:"account_to_recover"`
	NewOwnerAuthority    *Authority    `json:"new_owner_authority"`
	RecentOwnerAuthority *Authority    `json:"recent_owner_authority"`
	Extensions           []interface{} `json:"extensions"`
}

//Type function that defines the type of operation RecoverAccountOperation.
func (op *RecoverAccountOperation) Type() OpType {
	return TypeRecoverAccount
}

//Data returns the operation data RecoverAccountOperation.
func (op *RecoverAccountOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type RecoverAccountOperation to bytes.
func (op *RecoverAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeRecoverAccount.Code()))
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewOwnerAuthority)
	enc.Encode(op.RecentOwnerAuthority)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
