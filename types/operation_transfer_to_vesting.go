package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//TransferToVestingOperation represents transfer_to_vesting operation data.
type TransferToVestingOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount *Asset `json:"amount"`
}

//Type function that defines the type of operation TransferToVestingOperation.
func (op *TransferToVestingOperation) Type() OpType {
	return TypeTransferToVesting
}

//Data returns the operation data TransferToVestingOperation.
func (op *TransferToVestingOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type TransferToVestingOperation to bytes.
func (op *TransferToVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToVesting.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	return enc.Err()
}
