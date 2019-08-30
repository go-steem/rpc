package types

//CurationRewardOperation represents curation_reward operation data.
type CurationRewardOperation struct {
	Curator         string `json:"curator"`
	Reward          *Asset `json:"reward"`
	CommentAuthor   string `json:"comment_author"`
	CommentPermlink string `json:"comment_permlink"`
}

//Type function that defines the type of operation CurationRewardOperation.
func (op *CurationRewardOperation) Type() OpType {
	return TypeCurationReward
}

//Data returns the operation data CurationRewardOperation.
func (op *CurationRewardOperation) Data() interface{} {
	return op
}
