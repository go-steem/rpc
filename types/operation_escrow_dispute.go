package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//EscrowDisputeOperation represents escrow_dispute operation data.
type EscrowDisputeOperation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowID uint32 `json:"escrow_id"`
}

//Type function that defines the type of operation EscrowDisputeOperation.
func (op *EscrowDisputeOperation) Type() OpType {
	return TypeEscrowDispute
}

//Data returns the operation data EscrowDisputeOperation.
func (op *EscrowDisputeOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type EscrowDisputeOperation to bytes.
func (op *EscrowDisputeOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeEscrowDispute.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Agent)
	enc.Encode(op.Who)
	enc.Encode(op.EscrowID)
	return enc.Err()
}
