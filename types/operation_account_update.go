package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//AccountUpdateOperation represents account_update operation data.
type AccountUpdateOperation struct {
	Account      string           `json:"account"`
	Owner        *Authority       `json:"owner,omitempty"`
	Active       *Authority       `json:"active,omitempty"`
	Posting      *Authority       `json:"posting,omitempty"`
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
		enc.Encode(byte(1))
		enc.Encode(op.Owner)
	} else {
		enc.Encode(byte(0))
	}
	if op.Active != nil {
		enc.Encode(byte(1))
		enc.Encode(op.Active)
	} else {
		enc.Encode(byte(0))
	}
	if op.Posting != nil {
		enc.Encode(byte(1))
		enc.Encode(op.Posting)
	} else {
		enc.Encode(byte(0))
	}
	enc.EncodePubKey(op.MemoKey)
	enc.Encode(op.JSONMetadata)
	return enc.Err()
}
