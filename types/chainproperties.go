package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//ChainPropertiesOLD is an additional structure used by other structures.
type ChainProperties struct {
	AccountCreationFee *Asset `json:"account_creation_fee"`
	MaximumBlockSize   uint32 `json:"maximum_block_size"`
	SBDInterestRate    uint16 `json:"sbd_interest_rate"`
}

//MarshalTransaction is a function of converting type ChainPropertiesOLD to bytes.
func (cp *ChainProperties) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.Encode(cp.AccountCreationFee)
	enc.Encode(cp.MaximumBlockSize)
	enc.Encode(cp.SBDInterestRate)
	return enc.Err()
}
