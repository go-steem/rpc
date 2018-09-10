package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//TransferOperation represents transfer operation data.
type TransferOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount *Asset `json:"amount"`
	Memo   string `json:"memo"`
}

//Type function that defines the type of operation TransferOperation.
func (op *TransferOperation) Type() OpType {
	return TypeTransfer
}

//Data returns the operation data TransferOperation.
func (op *TransferOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type TransferOperation to bytes.
func (op *TransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransfer.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}
