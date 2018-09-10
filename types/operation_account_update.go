package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//AccountUpdateOperation represents account_update operation data.
type AccountUpdateOperation struct {
	Account      string           `json:"account"`
	Owner        *Authority       `json:"owner"`
	Active       *Authority       `json:"active"`
	Posting      *Authority       `json:"posting"`
	MemoKey      string           `json:"memo_key"`
	JSONMetadata *AccountMetadata `json:"json_metadata"`
}

//Type function that defines the type of operation AccountUpdateOperation.
func (op *AccountUpdateOperation) Type() OpType {
	return TypeAccountUpdate
}

//Data returns the operation data AccountUpdateOperation.
func (op *AccountUpdateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AccountUpdateOperation to bytes.
func (op *AccountUpdateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountUpdate.Code()))
	enc.EncodeString(op.Account)
	if op.Owner != nil {
		enc.Encode(op.Owner)
	}
	if op.Active != nil {
		enc.Encode(op.Active)
	}
	if op.Posting != nil {
		enc.Encode(op.Posting)
	}
	enc.EncodePubKey(op.MemoKey)
	enc.Encode(op.JSONMetadata)
	return enc.Err()
}
