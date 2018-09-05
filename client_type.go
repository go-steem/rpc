package rpc

import (
	"github.com/asuleymanov/steem-go/types"
)

var (
	//OptionsTrue variable for type PCOptions
	OptionsTrue bool = true
	//OptionsFalse variable for type PCOptions
	OptionsFalse bool = false
)

//Keys is used as a keystroke for a specific user.
//Only a few keys can be set.
type Keys struct {
	PKey []string
	AKey []string
	OKey []string
	MKey []string
}

//BResp of response when sending a transaction.
type BResp struct {
	ID       string
	BlockNum uint32
	TrxNum   uint32
	Expired  bool
}

//OperResp type is returned when the operation is performed.
type OperResp struct {
	NameOper string
	PermLink string
	Bresp    *BResp
}

//PCOptions for the Comment and Post functions.
//Sets the water to receive payment for a comment or post.
type PCOptions struct {
	Percent              uint16
	AllowVotes           *bool
	AllowCurationRewards *bool
	BeneficiarieList     *[]types.Beneficiarie
}

//ArrTransfer for MultiTransfer function
type ArrTransfer struct {
	To      string
	Memo    string
	Ammount string
}

//ArrVote for MultiVote function
type ArrVote struct {
	User   string
	Weight int
}
