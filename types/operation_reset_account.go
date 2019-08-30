package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//ResetAccountOperation represents reset_account operation data.
type ResetAccountOperation struct {
	ResetAccount      string     `json:"reset_account"`
	AccountToReset    string     `json:"account_to_reset"`
	NewOwnerAuthority *Authority `json:"new_owner_authority"`
}

//Type function that defines the type of operation ResetAccountOperation.
func (op *ResetAccountOperation) Type() OpType {
	return TypeResetAccount
}

//Data returns the operation data ResetAccountOperation.
func (op *ResetAccountOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ResetAccountOperation to bytes.
func (op *ResetAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeResetAccount.Code()))
	enc.Encode(op.ResetAccount)
	enc.Encode(op.AccountToReset)
	enc.Encode(op.NewOwnerAuthority)
	return enc.Err()
}
