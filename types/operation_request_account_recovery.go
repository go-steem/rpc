package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//RequestAccountRecoveryOperation represents request_account_recovery operation data.
type RequestAccountRecoveryOperation struct {
	RecoveryAccount   string        `json:"recovery_account"`
	AccountToRecover  string        `json:"account_to_recover"`
	NewOwnerAuthority *Authority    `json:"new_owner_authority"`
	Extensions        []interface{} `json:"extensions"`
}

//Type function that defines the type of operation RequestAccountRecoveryOperation.
func (op *RequestAccountRecoveryOperation) Type() OpType {
	return TypeRequestAccountRecovery
}

//Data returns the operation data RequestAccountRecoveryOperation.
func (op *RequestAccountRecoveryOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type RequestAccountRecoveryOperation to bytes.
func (op *RequestAccountRecoveryOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeRequestAccountRecovery.Code()))
	enc.Encode(op.RecoveryAccount)
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewOwnerAuthority)
	//enc.Encode(op.Extensions)
	enc.Encode(byte(0))
	return enc.Err()
}
