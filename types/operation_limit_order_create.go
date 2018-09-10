package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//LimitOrderCreateOperation represents limit_order_create operation data.
type LimitOrderCreateOperation struct {
	Owner        string `json:"owner"`
	OrderID      uint32 `json:"orderid"`
	AmountToSell *Asset `json:"amount_to_sell"`
	MinToReceive *Asset `json:"min_to_receive"`
	FillOrKill   bool   `json:"fill_or_kill"`
	Expiration   *Time  `json:"expiration"`
}

//Type function that defines the type of operation LimitOrderCreateOperation.
func (op *LimitOrderCreateOperation) Type() OpType {
	return TypeLimitOrderCreate
}

//Data returns the operation data LimitOrderCreateOperation.
func (op *LimitOrderCreateOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type LimitOrderCreateOperation to bytes.
func (op *LimitOrderCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCreate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	enc.Encode(op.AmountToSell)
	enc.Encode(op.MinToReceive)
	enc.EncodeBool(op.FillOrKill)
	enc.Encode(op.Expiration)
	return enc.Err()
}
