package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//CreateClaimedAccountOperation represents create_claimed_account operation data.
type CreateClaimedAccountOperation struct {
	Creator        string        `json:"creator"`
	NewAccountName string        `json:"new_account_name"`
	Owner          *Authority    `json:"owner"`
	Active         *Authority    `json:"active"`
	Posting        *Authority    `json:"posting"`
	MemoKey        string        `json:"memo_key"`
	JsonMetadata   string        `json:"json_metadata"`
	Extensions     []interface{} `json:"extensions"`
}

//Type function that defines the type of operation CreateClaimedAccountOperation.
func (op *CreateClaimedAccountOperation) Type() OpType {
	return TypeCreateClaimedAccount
}

//Data returns the operation data CreateClaimedAccountOperation.
func (op *CreateClaimedAccountOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type CreateClaimedAccountOperation to bytes.
func (op *CreateClaimedAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCreateClaimedAccount.Code()))
	enc.Encode(op.Creator)
	enc.Encode(op.NewAccountName)
	enc.Encode(op.Owner)
	enc.Encode(op.Active)
	enc.Encode(op.Posting)
	enc.Encode(op.MemoKey)
	enc.Encode(op.JsonMetadata)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
