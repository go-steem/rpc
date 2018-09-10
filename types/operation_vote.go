package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//VoteOperation represents vote operation data.
type VoteOperation struct {
	Voter    string `json:"voter"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Weight   Int16  `json:"weight"`
}

//Type function that defines the type of operation VoteOperation.
func (op *VoteOperation) Type() OpType {
	return TypeVote
}

//Data returns the operation data VoteOperation.
func (op *VoteOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type VoteOperation to bytes.
func (op *VoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeVote.Code()))
	enc.Encode(op.Voter)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Weight)
	return enc.Err()
}
