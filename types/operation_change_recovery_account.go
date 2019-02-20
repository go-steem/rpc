package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//ChangeRecoveryAccountOperation represents change_recovery_account operation data.
type ChangeRecoveryAccountOperation struct {
	AccountToRecover   string        `json:"account_to_recover"`
	NewRecoveryAccount string        `json:"new_recovery_account"`
	Extensions         []interface{} `json:"extensions"`
}

//Type function that defines the type of operation ChangeRecoveryAccountOperation.
func (op *ChangeRecoveryAccountOperation) Type() OpType {
	return TypeChangeRecoveryAccount
}

//Data returns the operation data ChangeRecoveryAccountOperation.
func (op *ChangeRecoveryAccountOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ChangeRecoveryAccountOperation to bytes.
func (op *ChangeRecoveryAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChangeRecoveryAccount.Code()))
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewRecoveryAccount)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
