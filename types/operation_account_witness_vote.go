package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//AccountWitnessVoteOperation represents account_witness_vote operation data.
type AccountWitnessVoteOperation struct {
	Account string `json:"account"`
	Witness string `json:"witness"`
	Approve bool   `json:"approve"`
}

//Type function that defines the type of operation AccountWitnessVoteOperation.
func (op *AccountWitnessVoteOperation) Type() OpType {
	return TypeAccountWitnessVote
}

//Data returns the operation data AccountWitnessVoteOperation.
func (op *AccountWitnessVoteOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type AccountWitnessVoteOperation to bytes.
func (op *AccountWitnessVoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessVote.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Witness)
	enc.EncodeBool(op.Approve)
	return enc.Err()
}
