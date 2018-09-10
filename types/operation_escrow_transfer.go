package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//EscrowTransferOperation represents escrow_transfer operation data.
type EscrowTransferOperation struct {
	From                 string `json:"from"`
	To                   string `json:"to"`
	SbdAmount            *Asset `json:"sbd_amount"`
	SteemAmount          *Asset `json:"steem_amount"`
	EscrowID             uint32 `json:"escrow_id"`
	Agent                string `json:"agent"`
	Fee                  *Asset `json:"fee"`
	JSONMeta             string `json:"json_meta"`
	RatificationDeadline *Time  `json:"ratification_deadline"`
	EscrowExpiration     *Time  `json:"escrow_expiration"`
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
	enc.Encode(op.SbdAmount)
	enc.Encode(op.SteemAmount)
	enc.Encode(op.EscrowID)
	enc.Encode(op.Agent)
	enc.Encode(op.Fee)
	enc.Encode(op.JSONMeta)
	enc.Encode(op.RatificationDeadline)
	enc.Encode(op.EscrowExpiration)
	return enc.Err()
}
