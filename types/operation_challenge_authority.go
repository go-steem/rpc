package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//ChallengeAuthorityOperation represents challenge_authority operation data.
type ChallengeAuthorityOperation struct {
	Challenger   string `json:"challenger"`
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

//Type function that defines the type of operation ChallengeAuthorityOperation.
func (op *ChallengeAuthorityOperation) Type() OpType {
	return TypeChallengeAuthority
}

//Data returns the operation data ChallengeAuthorityOperation.
func (op *ChallengeAuthorityOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type ChallengeAuthorityOperation to bytes.
func (op *ChallengeAuthorityOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChallengeAuthority.Code()))
	enc.Encode(op.Challenger)
	enc.Encode(op.Challenged)
	enc.EncodeBool(op.RequireOwner)
	return enc.Err()
}
