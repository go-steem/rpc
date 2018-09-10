package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//SetResetAccountOperation represents set_reset_account operation data.
type SetResetAccountOperation struct {
	Account             string `json:"account"`
	CurrentResetAccount string `json:"current_reset_account"`
	ResetAccount        string `json:"reset_account"`
}

//Type function that defines the type of operation SetResetAccountOperation.
func (op *SetResetAccountOperation) Type() OpType {
	return TypeSetResetAccount
}

//Data returns the operation data SetResetAccountOperation.
func (op *SetResetAccountOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type SetResetAccountOperation to bytes.
func (op *SetResetAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetResetAccount.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.CurrentResetAccount)
	enc.Encode(op.ResetAccount)
	return enc.Err()
}
