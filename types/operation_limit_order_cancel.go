package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//LimitOrderCancelOperation represents limit_order_cancel operation data.
type LimitOrderCancelOperation struct {
	Owner   string `json:"owner"`
	OrderID uint32 `json:"orderid"`
}

//Type function that defines the type of operation LimitOrderCancelOperation.
func (op *LimitOrderCancelOperation) Type() OpType {
	return TypeLimitOrderCancel
}

//Data returns the operation data LimitOrderCancelOperation.
func (op *LimitOrderCancelOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type LimitOrderCancelOperation to bytes.
func (op *LimitOrderCancelOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCancel.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	return enc.Err()
}
