package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//TransferFromSavingsOperation represents transfer_from_savings operation data.
type TransferFromSavingsOperation struct {
	From      string `json:"from"`
	RequestID uint32 `json:"request_id"`
	To        string `json:"to"`
	Amount    *Asset `json:"amount"`
	Memo      string `json:"memo"`
}

//Type function that defines the type of operation TransferFromSavingsOperation.
func (op *TransferFromSavingsOperation) Type() OpType {
	return TypeTransferFromSavings
}

//Data returns the operation data TransferFromSavingsOperation.
func (op *TransferFromSavingsOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type TransferFromSavingsOperation to bytes.
func (op *TransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestID)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}
