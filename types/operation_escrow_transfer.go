package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//EscrowTransferOperation represents escrow_transfer operation data.
type EscrowTransferOperation struct {
	From                 string `json:"from"`
	To                   string `json:"to"`
	Agent                string `json:"agent"`
	EscrowID             uint32 `json:"escrow_id"`
	SbdAmount            *Asset `json:"sbd_amount"`
	SteemAmount          *Asset `json:"steem_amount"`
	Fee                  *Asset `json:"fee"`
	RatificationDeadline *Time  `json:"ratification_deadline"`
	EscrowExpiration     *Time  `json:"escrow_expiration"`
	JSONMeta             string `json:"json_meta"`
}

//Type function that defines the type of operation EscrowTransferOperation.
func (op *EscrowTransferOperation) Type() OpType {
	return TypeEscrowTransfer
}

//Data returns the operation data EscrowTransferOperation.
func (op *EscrowTransferOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type EscrowTransferOperation to bytes.
func (op *EscrowTransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeEscrowTransfer.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Agent)
	enc.Encode(op.EscrowID)
	enc.Encode(op.SbdAmount)
	enc.Encode(op.SteemAmount)
	enc.Encode(op.Fee)
	enc.Encode(op.RatificationDeadline)
	enc.Encode(op.EscrowExpiration)
	enc.Encode(op.JSONMeta)
	return enc.Err()
}
