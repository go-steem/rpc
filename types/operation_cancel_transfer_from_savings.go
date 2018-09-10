package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//CancelTransferFromSavingsOperation represents cancel_transfer_from_savings operation data.
type CancelTransferFromSavingsOperation struct {
	From      string `json:"from"`
	RequestID uint32 `json:"request_id"`
}

//Type function that defines the type of operation CancelTransferFromSavingsOperation.
func (op *CancelTransferFromSavingsOperation) Type() OpType {
	return TypeCancelTransferFromSavings
}

//Data returns the operation data CancelTransferFromSavingsOperation.
func (op *CancelTransferFromSavingsOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type CancelTransferFromSavingsOperation to bytes.
func (op *CancelTransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCancelTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestID)
	return enc.Err()
}
