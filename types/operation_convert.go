package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//ConvertOperation represents convert operation data.
type ConvertOperation struct {
	Owner     string `json:"owner"`
	RequestID uint32 `json:"requestid"`
	Amount    *Asset `json:"amount"`
}

//Type function that defines the type of operation ConvertOperation.
func (op *ConvertOperation) Type() OpType {
	return TypeConvert
}

//Data returns the operation data ConvertOperation.
func (op *ConvertOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ConvertOperation to bytes.
func (op *ConvertOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeConvert.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.RequestID)
	enc.Encode(op.Amount)
	return enc.Err()
}
