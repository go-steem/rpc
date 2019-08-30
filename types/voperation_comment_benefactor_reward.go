package types

//CommentBenefactorRewardOperation represents comment_benefactor_reward operation data.
type CommentBenefactorRewardOperation struct {
	Benefactor    string `json:"benefactor"`
	Author        string `json:"author"`
	Permlink      string `json:"permlink"`
	SbdPayout     *Asset `json:"sbd_payout"`
	SteemPayout   *Asset `json:"steem_payout"`
	VestingPayout *Asset `json:"vesting_payout"`
}

//Type function that defines the type of operation CommentBenefactorRewardOperation.
func (op *CommentBenefactorRewardOperation) Type() OpType {
	return TypeCommentBenefactorReward
}

//Data returns the operation data CommentBenefactorRewardOperation.
func (op *CommentBenefactorRewardOperation) Data() interface{} {
	return op
}
