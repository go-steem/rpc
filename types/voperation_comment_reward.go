package types

//CommentRewardOperation represents comment_reward operation data.
type CommentRewardOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Payout   *Asset `json:"payout"`
}

//Type function that defines the type of operation CommentRewardOperation.
func (op *CommentRewardOperation) Type() OpType {
	return TypeCommentReward
}

//Data returns the operation data CommentRewardOperation.
func (op *CommentRewardOperation) Data() interface{} {
	return op
}
