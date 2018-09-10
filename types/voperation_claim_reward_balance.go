package types

//ClaimRewardBalanceOperation represents claim_reward_balance operation data.
type ClaimRewardBalanceOperation struct {
	Account     string `json:"account"`
	RewardSteem string `json:"reward_steem"`
	RewardSbd   string `json:"reward_sbd"`
	RewardVests string `json:"reward_vests"`
}

//Type function that defines the type of operation ClaimRewardBalanceOperation.
func (op *ClaimRewardBalanceOperation) Type() OpType {
	return TypeClaimRewardBalance
}

//Data returns the operation data ClaimRewardBalanceOperation.
func (op *ClaimRewardBalanceOperation) Data() interface{} {
	return op
}
