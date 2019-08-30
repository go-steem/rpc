package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//DeclineVotingRightsOperation represents decline_voting_rights operation data.
type DeclineVotingRightsOperation struct {
	Account string `json:"account"`
	Decline bool   `json:"decline"`
}

//Type function that defines the type of operation DeclineVotingRightsOperation.
func (op *DeclineVotingRightsOperation) Type() OpType {
	return TypeDeclineVotingRights
}

//Data returns the operation data DeclineVotingRightsOperation.
func (op *DeclineVotingRightsOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type DeclineVotingRightsOperation to bytes.
func (op *DeclineVotingRightsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeclineVotingRights.Code()))
	enc.Encode(op.Account)
	enc.EncodeBool(op.Decline)
	return enc.Err()
}
