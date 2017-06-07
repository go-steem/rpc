package types

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/golos-go/encoding/transaction"

	// Vendor
	"github.com/pkg/errors"
)

// FC_REFLECT( steemit::chain::report_over_production_operation,
//             (reporter)
//             (first_block)
//             (second_block) )

type ReportOverProductionOperation struct {
	Reporter string `json:"reporter"`
}

func (op *ReportOverProductionOperation) Type() OpType {
	return TypeReportOverProduction
}

func (op *ReportOverProductionOperation) Data() interface{} {
	return op
}

// FC_REFLECT( steemit::chain::convert_operation,
//             (owner)
//             (requestid)
//             (amount) )

type ConvertOperation struct {
	Owner     string `json:"owner"`
	RequestID uint32 `json:"requestid"`
	Amount    string `json:"amount"`
}

func (op *ConvertOperation) Type() OpType {
	return TypeConvert
}

func (op *ConvertOperation) Data() interface{} {
	return op
}

// FC_REFLECT( steemit::chain::feed_publish_operation,
//             (publisher)
//             (exchange_rate) )

type FeedPublishOperation struct {
	Publisher    string `json:"publisher"`
	ExchangeRate struct {
		Base  string `json:"base"`
		Quote string `json:"quote"`
	} `json:"exchange_rate"`
}

func (op *FeedPublishOperation) Type() OpType {
	return TypeFeedPublish
}

func (op *FeedPublishOperation) Data() interface{} {
	return op
}

// FC_REFLECT( steemit::chain::pow,
//             (worker)
//             (input)
//             (signature)
//             (work) )

type POW struct {
	Worker    string `json:"worker"`
	Input     string `json:"input"`
	Signature string `json:"signature"`
	Work      string `json:"work"`
}

// FC_REFLECT( steemit::chain::chain_properties,
//             (account_creation_fee)
//             (maximum_block_size)
//             (sbd_interest_rate) );

type ChainProperties struct {
	AccountCreationFee string `json:"account_creation_fee"`
	MaximumBlockSize   uint32 `json:"maximum_block_size"`
	SBDInterestRate    uint16 `json:"sbd_interest_rate"`
}

// FC_REFLECT( steemit::chain::pow_operation,
//             (worker_account)
//             (block_id)
//             (nonce)
//             (work)
//             (props) )

type POWOperation struct {
	WorkerAccount string           `json:"worker_account"`
	BlockID       string           `json:"block_id"`
	Nonce         *Int             `json:"nonce"`
	Work          *POW             `json:"work"`
	Props         *ChainProperties `json:"props"`
}

func (op *POWOperation) Type() OpType {
	return TypePOW
}

func (op *POWOperation) Data() interface{} {
	return op
}

// FC_REFLECT( steemit::chain::account_create_operation,
//             (fee)
//             (creator)
//             (new_account_name)
//             (owner)
//             (active)
//             (posting)
//             (memo_key)
//             (json_metadata) )

type AccountCreateOperation struct {
	Fee            string     `json:"fee"`
	Creator        string     `json:"creator"`
	NewAccountName string     `json:"new_account_name"`
	Owner          *Authority `json:"owner"`
	Active         *Authority `json:"active"`
	Posting        *Authority `json:"posting"`
	MemoKey        string     `json:"memo_key"`
	JsonMetadata   string     `json:"json_metadata"`
}

func (op *AccountCreateOperation) Type() OpType {
	return TypeAccountCreate
}

func (op *AccountCreateOperation) Data() interface{} {
	return op
}

func (op *AccountCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountCreate.Code()))
	enc.Encode(op.Fee)
	enc.Encode(op.Creator)
	enc.Encode(op.NewAccountName)
	enc.Encode(op.Owner)
	enc.Encode(op.Active)
	enc.Encode(op.Posting)
	enc.Encode(op.MemoKey)
	enc.Encode(op.JsonMetadata)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::account_update_operation,
//             (account)
//             (owner)
//             (active)
//             (posting)
//             (memo_key)
//             (json_metadata) )

type AccountUpdateOperation struct {
	Account      string     `json:"account"`
	Owner        *Authority `json:"owner"`
	Active       *Authority `json:"active"`
	Posting      *Authority `json:"posting"`
	MemoKey      string     `json:"memo_key"`
	JsonMetadata string     `json:"json_metadata"`
}

func (op *AccountUpdateOperation) Type() OpType {
	return TypeAccountUpdate
}

func (op *AccountUpdateOperation) Data() interface{} {
	return op
}

// FC_REFLECT( steemit::chain::transfer_operation,
//             (from)
//             (to)
//             (amount)
//             (memo) )

type TransferOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
	Memo   string `json:"memo"`
}

func (op *TransferOperation) Type() OpType {
	return TypeTransfer
}

func (op *TransferOperation) Data() interface{} {
	return op
}

func (op *TransferOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransfer.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::transfer_to_vesting_operation,
//             (from)
//             (to)
//             (amount) )

type TransferToVestingOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

func (op *TransferToVestingOperation) Type() OpType {
	return TypeTransferToVesting
}

func (op *TransferToVestingOperation) Data() interface{} {
	return op
}

func (op *TransferToVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToVesting.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.Encode(op.Amount)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::withdraw_vesting_operation,
//             (account)
//             (vesting_shares) )

type WithdrawVestingOperation struct {
	Account       string `json:"account"`
	VestingShares string `json:"vesting_shares"`
}

func (op *WithdrawVestingOperation) Type() OpType {
	return TypeWithdrawVesting
}

func (op *WithdrawVestingOperation) Data() interface{} {
	return op
}

// FC_REFLECT( steemit::chain::set_withdraw_vesting_route_operation,
//             (from_account)
//             (to_account)
//             (percent)
//             (auto_vest) )

// FC_REFLECT( steemit::chain::witness_update_operation,
//             (owner)
//             (url)
//             (block_signing_key)
//             (props)
//             (fee) )

// FC_REFLECT( steemit::chain::account_witness_vote_operation,
//             (account)
//             (witness)(approve) )

type AccountWitnessVoteOperation struct {
	Account string `json:"account"`
	Witness string `json:"witness"`
	Approve bool   `json:"approve"`
}

func (op *AccountWitnessVoteOperation) Type() OpType {
	return TypeAccountWitnessVote
}

func (op *AccountWitnessVoteOperation) Data() interface{} {
	return op
}

// FC_REFLECT( steemit::chain::account_witness_proxy_operation,
//             (account)
//             (proxy) )

type AccountWitnessProxyOperation struct {
	Account string `json:"account"`
	Proxy   string `json:"proxy"`
}

func (op *AccountWitnessProxyOperation) Type() OpType {
	return TypeAccountWitnessProxy
}

func (op *AccountWitnessProxyOperation) Data() interface{} {
	return op
}

// FC_REFLECT( steemit::chain::comment_operation,
//             (parent_author)
//             (parent_permlink)
//             (author)
//             (permlink)
//             (title)
//             (body)
//             (json_metadata) )

// CommentOperation represents either a new post or a comment.
//
// In case Title is filled in and ParentAuthor is empty, it is a new post.
// The post category can be read from ParentPermlink.
type CommentOperation struct {
	Author         string `json:"author"`
	Title          string `json:"title"`
	Permlink       string `json:"permlink"`
	ParentAuthor   string `json:"parent_author"`
	ParentPermlink string `json:"parent_permlink"`
	Body           string `json:"body"`
}

func (op *CommentOperation) Type() OpType {
	return TypeComment
}

func (op *CommentOperation) Data() interface{} {
	return op
}

func (op *CommentOperation) IsStoryOperation() bool {
	return op.ParentAuthor == ""
}

func (op *CommentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeComment.Code()))
	enc.Encode(op.ParentAuthor)
	enc.Encode(op.ParentPermlink)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Title)
	enc.Encode(op.Body)
	enc.Encode(op.JMeta)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::vote_operation,
//             (voter)
//             (author)
//             (permlink)
//             (weight) )

type VoteOperation struct {
	Voter    string `json:"voter"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Weight   Int16  `json:"weight"`
}

func (op *VoteOperation) Type() OpType {
	return TypeVote
}

func (op *VoteOperation) Data() interface{} {
	return op
}

func (op *VoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeVote.Code()))
	enc.Encode(op.Voter)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Weight)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::custom_operation,
//             (required_auths)
//             (id)
//             (data) )

// FC_REFLECT( steemit::chain::limit_order_create_operation,
//             (owner)
//             (orderid)
//             (amount_to_sell)
//             (min_to_receive)
//             (fill_or_kill)
//             (expiration) )

type LimitOrderCreateOperation struct {
	Owner        string `json:"owner"`
	OrderID      uint32 `json:"orderid"`
	AmountToSell string `json:"amount_to_sell"`
	MinToReceive string `json:"min_to_receive"`
	FillOrKill   bool   `json:"fill_or_kill"`
	Expiration   *Time  `json:"expiration"`
}

func (op *LimitOrderCreateOperation) Type() OpType {
	return TypeLimitOrderCreate
}

func (op *LimitOrderCreateOperation) Data() interface{} {
	return op
}

// FC_REFLECT( steemit::chain::limit_order_cancel_operation,
//             (owner)
//             (orderid) )

type LimitOrderCancelOperation struct {
	Owner   string `json:"owner"`
	OrderID uint32 `json:"orderid"`
}

func (op *LimitOrderCancelOperation) Type() OpType {
	return TypeLimitOrderCancel
}

func (op *LimitOrderCancelOperation) Data() interface{} {
	return op
}

// FC_REFLECT( steemit::chain::delete_comment_operation,
//             (author)
//             (permlink) )

type DeleteCommentOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

func (op *DeleteCommentOperation) Type() OpType {
	return TypeDeleteComment
}

func (op *DeleteCommentOperation) Data() interface{} {
	return op
}

func (op *DeleteCommentOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeleteComment.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::comment_options_operation,
//             (author)
//             (permlink)
//             (max_accepted_payout)
//             (percent_steem_dollars)
//             (allow_votes)
//             (allow_curation_rewards)
//             (extensions) )

type CommentOptionsOperation struct {
	Author               string        `json:"author"`
	Permlink             string        `json:"permlink"`
	MaxAcceptedPayout    string        `json:"max_accepted_payout"`
	PercentSteemDollars  uint16        `json:"percent_steem_dollars"`
	AllowVotes           bool          `json:"allow_votes"`
	AllowCurationRewards bool          `json:"allow_curation_rewards"`
	Extensions           []interface{} `json:"extensions"`
}

func (op *CommentOptionsOperation) Type() OpType {
	return TypeCommentOptions
}

func (op *CommentOptionsOperation) Data() interface{} {
	return op
}

type Authority struct {
	AccountAuths    []*Auth `json:"account_auths"`
	KeyAuths        []*Auth `json:"key_auths"`
	WeightThreshold uint32  `json:"weight_threshold"`
}

// XXX: Not sure about the struct field names.
type Auth struct {
	Key   string
	Check uint32
}

func (auth *Auth) UnmarshalJSON(data []byte) error {
	// The auth object is [key, check].
	raw := make([]json.RawMessage, 2)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if len(raw) != 2 {
		return errors.Errorf("invalid auth object: %v", string(data))
	}

	// Unmarshal Key.
	var key string
	if err := json.Unmarshal(raw[0], &key); err != nil {
		return errors.Wrapf(err, "failed to unmarshal Auth.Key: %v", string(raw[0]))
	}

	// Unmarshal Check.
	var check uint32
	if err := json.Unmarshal(raw[1], &check); err != nil {
		return errors.Wrapf(err, "failed to unmarshal Auth.Check: %v", string(raw[1]))
	}

	// Update fields.
	auth.Key = key
	auth.Check = check
	return nil
}

type UnknownOperation struct {
	kind OpType
	data *json.RawMessage
}

func (op *UnknownOperation) Type() OpType {
	return op.kind
}

func (op *UnknownOperation) Data() interface{} {
	return op.data
}
