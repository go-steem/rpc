package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//TransferToSavingsOperation represents transfer_to_savings operation data.
type TransferToSavingsOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount *Asset `json:"amount"`
	Memo   string `json:"memo"`
}

//Type function that defines the type of operation TransferToSavingsOperation.
func (op *TransferToSavingsOperation) Type() OpType {
	return TypeTransferToSavings
}

//Data returns the operation data TransferToSavingsOperation.
func (op *TransferToSavingsOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type TransferToSavingsOperation to bytes.
func (op *TransferToSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}
