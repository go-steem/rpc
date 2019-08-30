package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//SetWithdrawVestingRouteOperation represents set_withdraw_vesting_route operation data.
type SetWithdrawVestingRouteOperation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Percent     uint16 `json:"percent"`
	AutoVest    bool   `json:"auto_vest"`
}

//Type function that defines the type of operation SetWithdrawVestingRouteOperation.
func (op *SetWithdrawVestingRouteOperation) Type() OpType {
	return TypeSetWithdrawVestingRoute
}

//Data returns the operation data SetWithdrawVestingRouteOperation.
func (op *SetWithdrawVestingRouteOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type SetWithdrawVestingRouteOperation to bytes.
func (op *SetWithdrawVestingRouteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetWithdrawVestingRoute.Code()))
	enc.Encode(op.FromAccount)
	enc.Encode(op.ToAccount)
	enc.Encode(op.Percent)
	enc.EncodeBool(op.AutoVest)
	return enc.Err()
}
