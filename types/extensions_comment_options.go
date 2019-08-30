package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//Beneficiary is an additional structure used by other structures.
type Beneficiary struct {
	Account string `json:"account"`
	Weight  uint16 `json:"weight"`
}

//CommentPayoutBeneficiaries is an additional structure used by other structures.
type CommentPayoutBeneficiaries struct {
	Beneficiaries []Beneficiary `json:"beneficiaries"`
}

//MarshalTransaction is a function of converting type CommentPayoutBeneficiaries to bytes.
func (cp *CommentPayoutBeneficiaries) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(byte(0))
	enc.EncodeUVarint(uint64(len(cp.Beneficiaries)))
	for _, val := range cp.Beneficiaries {
		enc.Encode(val.Account)
		enc.Encode(val.Weight)
	}
	return enc.Err()
}

//AllowedVoteAssets
