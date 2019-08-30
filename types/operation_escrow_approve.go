package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//EscrowApproveOperation represents escrow_approve operation data.
type EscrowApproveOperation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowID uint32 `json:"escrow_id"`
	Approve  bool   `json:"approve"`
}

//Type function that defines the type of operation EscrowApproveOperation.
func (op *EscrowApproveOperation) Type() OpType {
	return TypeEscrowApprove
}

//Data returns the operation data EscrowApproveOperation.
func (op *EscrowApproveOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type EscrowApproveOperation to bytes.
func (op *EscrowApproveOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeEscrowApprove.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Agent)
	enc.Encode(op.Who)
	enc.Encode(op.EscrowID)
	enc.EncodeBool(op.Approve)
	return enc.Err()
}
