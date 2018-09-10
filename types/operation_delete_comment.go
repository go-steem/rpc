package types

import (
	"github.com/asuleymanov/steem-go/encoding/transaction"
)

//DeleteCommentOperation represents delete_comment operation data.
type DeleteCommentOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

//Type function that defines the type of operation DeleteCommentOperation.
func (op *DeleteCommentOperation) Type() OpType {
	return TypeDeleteComment
}

//Data returns the operation data DeleteCommentOperation.
func (op *DeleteCommentOperation) Data() interface{} {
	return op
}

//MarshalTransaction is a function of converting type DeleteCommentOperation to bytes.
func (op *DeleteCommentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeleteComment.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	return enc.Err()
}
