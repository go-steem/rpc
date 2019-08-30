package types

//AuthorRewardOperation represents author_reward operation data.
type AuthorRewardOperation struct {
	Author        string `json:"author"`
	Permlink      string `json:"permlink"`
	SbdPayout     *Asset `json:"sbd_payout"`
	SteemPayout   *Asset `json:"steem_payout"`
	VestingPayout *Asset `json:"vesting_payout"`
}

//Type function that defines the type of operation AuthorRewardOperation.
func (op *AuthorRewardOperation) Type() OpType {
	return TypeAuthorReward
}

//Data returns the operation data AuthorRewardOperation.
func (op *AuthorRewardOperation) Data() interface{} {
	return op
}
