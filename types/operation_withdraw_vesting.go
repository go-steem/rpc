package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//WithdrawVestingOperation represents withdraw_vesting operation data.
type WithdrawVestingOperation struct {
	Account       string `json:"account"`
	VestingShares *Asset `json:"vesting_shares"`
}

//Type function that defines the type of operation WithdrawVestingOperation.
func (op *WithdrawVestingOperation) Type() OpType {
	return TypeWithdrawVesting
}

//Data returns the operation data WithdrawVestingOperation.
func (op *WithdrawVestingOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type WithdrawVestingOperation to bytes.
func (op *WithdrawVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWithdrawVesting.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.VestingShares)
	return enc.Err()
}
