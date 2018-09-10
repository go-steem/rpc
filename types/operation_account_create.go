package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//AccountCreateOperation represents account_create operation data.
type AccountCreateOperation struct {
	Fee            *Asset           `json:"fee"`
	Creator        string           `json:"creator"`
	NewAccountName string           `json:"new_account_name"`
	Owner          *Authority       `json:"owner"`
	Active         *Authority       `json:"active"`
	Posting        *Authority       `json:"posting"`
	MemoKey        string           `json:"memo_key"`
	JSONMetadata   *AccountMetadata `json:"json_metadata"`
}

//Type function that defines the type of operation AccountCreateOperation.
func (op *AccountCreateOperation) Type() OpType {
	return TypeAccountCreate
}

//Data returns the operation data AccountCreateOperation.
func (op *AccountCreateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AccountCreateOperation to bytes.
func (op *AccountCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountCreate.Code()))
	enc.Encode(op.Fee)
	enc.EncodeString(op.Creator)
	enc.EncodeString(op.NewAccountName)
	enc.Encode(op.Owner)
	enc.Encode(op.Active)
	enc.Encode(op.Posting)
	enc.EncodePubKey(op.MemoKey)
	enc.Encode(op.JSONMetadata)
	return enc.Err()
}
