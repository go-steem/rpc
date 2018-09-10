package types

//CommentBenefactorRewardOperation represents comment_benefactor_reward operation data.
type CommentBenefactorRewardOperation struct {
	Benefactor string `json:"benefactor"`
	Author     string `json:"author"`
	Permlink   string `json:"permlink"`
	Reward     *Asset `json:"reward"`
}

//Type function that defines the type of operation CommentBenefactorRewardOperation.
func (op *CommentBenefactorRewardOperation) Type() OpType {
	return TypeCommentBenefactorReward
}

//Data returns the operation data CommentBenefactorRewardOperation.
func (op *CommentBenefactorRewardOperation) Data() interface{} {
	return op
}
