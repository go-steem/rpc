package types

//CommentPayoutUpdateOperation represents comment_payout_update operation data.
type CommentPayoutUpdateOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

//Type function that defines the type of operation CommentPayoutUpdateOperation.
func (op *CommentPayoutUpdateOperation) Type() OpType {
	return TypeCommentPayoutUpdate
}

//Data returns the operation data CommentPayoutUpdateOperation.
func (op *CommentPayoutUpdateOperation) Data() interface{} {
	return op
}
