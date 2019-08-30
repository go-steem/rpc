package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//CustomOperation represents custom operation data.
type CustomOperation struct {
	RequiredAuths []string `json:"required_auths"`
	ID            uint16   `json:"id"`
	Datas         []byte   `json:"data"`
}

//Type function that defines the type of operation CustomOperation.
func (op *CustomOperation) Type() OpType {
	return TypeCustom
}

//Data returns the operation data CustomOperation.
func (op *CustomOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type CustomOperation to bytes.
func (op *CustomOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCustom.Code()))
	enc.Encode(op.RequiredAuths)
	enc.Encode(op.ID)
	enc.Encode(op.Datas)
	return enc.Err()
}
