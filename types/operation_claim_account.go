package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//ClaimAccountOperation represents claim_account operation data.
type ClaimAccountOperation struct {
	Creator    string        `json:"creator"`
	Fee        *Asset        `json:"fee"`
	Extensions []interface{} `json:"extensions"`
}

//Type function that defines the type of operation ClaimAccountOperation.
func (op *ClaimAccountOperation) Type() OpType {
	return TypeClaimAccount
}

//Data returns the operation data ClaimAccountOperation.
func (op *ClaimAccountOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ClaimAccountOperation to bytes.
func (op *ClaimAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeClaimAccount.Code()))
	enc.Encode(op.Creator)
	enc.Encode(op.Fee)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
