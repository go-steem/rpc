package types

import (
	"encoding/json"

	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//CommentOptionsOperation represents comment_options operation data.
type CommentOptionsOperation struct {
	Author               string        `json:"author"`
	Permlink             string        `json:"permlink"`
	MaxAcceptedPayout    *Asset        `json:"max_accepted_payout"`
	PercentSteemDollars  uint16        `json:"percent_steem_dollars"`
	AllowVotes           bool          `json:"allow_votes"`
	AllowCurationRewards bool          `json:"allow_curation_rewards"`
	Extensions           []interface{} `json:"extensions"`
}

//Type function that defines the type of operation CommentOptionsOperation.
func (op *CommentOptionsOperation) Type() OpType {
	return TypeCommentOptions
}

//Data returns the operation data CommentOptionsOperation.
func (op *CommentOptionsOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type CommentOptionsOperation to bytes.
func (op *CommentOptionsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCommentOptions.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.MaxAcceptedPayout)
	enc.Encode(op.PercentSteemDollars)
	enc.EncodeBool(op.AllowVotes)
	enc.EncodeBool(op.AllowCurationRewards)
	if len(op.Extensions) > 0 {
		//Parse Beneficiaries
		z, _ := json.Marshal(op.Extensions[0])
		var l []interface{}
		_ = json.Unmarshal(z, &l)
		z1, _ := json.Marshal(l[1])
		var d CommentPayoutBeneficiaries
		_ = json.Unmarshal(z1, &d)

		enc.Encode(byte(1))
		enc.Encode(byte(0))
		enc.EncodeUVarint(uint64(len(d.Beneficiaries)))
		for _, val := range d.Beneficiaries {
			enc.Encode(val.Account)
			enc.Encode(val.Weight)
		}
	} else {
		enc.Encode(byte(0))
	}
	return enc.Err()
}
