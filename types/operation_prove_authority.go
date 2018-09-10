package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//ProveAuthorityOperation represents prove_authority operation data.
type ProveAuthorityOperation struct {
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

//Type function that defines the type of operation ProveAuthorityOperation.
func (op *ProveAuthorityOperation) Type() OpType {
	return TypeProveAuthority
}

//Data returns the operation data ProveAuthorityOperation.
func (op *ProveAuthorityOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ProveAuthorityOperation to bytes.
func (op *ProveAuthorityOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeProveAuthority.Code()))
	enc.Encode(op.Challenged)
	enc.EncodeBool(op.RequireOwner)
	return enc.Err()
}
