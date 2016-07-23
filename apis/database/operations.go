package database

import (
	"encoding/json"
	"reflect"

	"github.com/go-steem/rpc/apis/types"

	"github.com/pkg/errors"
)

const (
	OpTypeConvert             = "convert"
	OpTypeFeedPublish         = "feed_publish"
	OpTypePow                 = "pow"
	OpTypeAccountCreate       = "account_create"
	OpTypeAccountUpdate       = "account_update"
	OpTypeTransfer            = "transfer"
	OpTypeTransferToVesting   = "transfer_to_vesting"
	OpTypeWithdrawVesting     = "withdraw_vesting"
	OpTypeAccountWitnessVote  = "account_witness_vote"
	OpTypeAccountWitnessProxy = "account_witness_proxy"
	OpTypeComment             = "comment"
	OpTypeVote                = "vote"
	OpTypeLimitOrderCreate    = "limit_order_create"
	OpTypeFillOrder           = "fill_order"
	OpTypeLimitOrderCancel    = "limit_order_cancel"
	OpTypeDeleteComment       = "delete_comment"
	OpTypeCommentOptions      = "comment_options"
)

var opBodyObjects = map[string]interface{}{
	OpTypeConvert:             &ConvertOperation{},
	OpTypeFeedPublish:         &FeedPublishOperation{},
	OpTypePow:                 &PowOperation{},
	OpTypeAccountCreate:       &AccountCreateOperation{},
	OpTypeAccountUpdate:       &AccountUpdateOperation{},
	OpTypeTransfer:            &TransferOperation{},
	OpTypeTransferToVesting:   &TransferToVestingOperation{},
	OpTypeWithdrawVesting:     &WithdrawVestingOperation{},
	OpTypeAccountWitnessVote:  &AccountWitnessVoteOperation{},
	OpTypeAccountWitnessProxy: &AccountWitnessProxyOperation{},
	OpTypeComment:             &CommentOperation{},
	OpTypeVote:                &VoteOperation{},
	OpTypeLimitOrderCreate:    &LimitOrderCreateOperation{},
	OpTypeFillOrder:           &FillOrderOperation{},
	OpTypeLimitOrderCancel:    &LimitOrderCancelOperation{},
	OpTypeDeleteComment:       &DeleteCommentOperation{},
	OpTypeCommentOptions:      &CommentOptionsOperation{},
}

// FC_REFLECT( steemit::chain::report_over_production_operation,
//             (reporter)
//             (first_block)
//             (second_block) )

type ReportOverProductionOperation struct {
	Reporter string `json:"reporter"`
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

// FC_REFLECT( steemit::chain::pow,
//             (worker)
//             (input)
//             (signature)
//             (work) )

type Pow struct {
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

type PowOperation struct {
	WorkerAccount string           `json:"worker_account"`
	BlockID       string           `json:"block_id"`
	Nonce         uint64           `json:"nonce"`
	Work          *Pow             `json:"work"`
	Props         *ChainProperties `json:"props"`
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

// FC_REFLECT( steemit::chain::transfer_to_vesting_operation,
//             (from)
//             (to)
//             (amount) )

type TransferToVestingOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

// FC_REFLECT( steemit::chain::withdraw_vesting_operation,
//             (account)
//             (vesting_shares) )

type WithdrawVestingOperation struct {
	Account       string `json:"account"`
	VestingShares string `json:"vesting_shares"`
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

// FC_REFLECT( steemit::chain::account_witness_proxy_operation,
//             (account)
//             (proxy) )

type AccountWitnessProxyOperation struct {
	Account string `json:"account"`
	Proxy   string `json:"proxy"`
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

func (op *CommentOperation) IsStoryOperation() bool {
	return op.ParentAuthor == ""
}

// FC_REFLECT( steemit::chain::vote_operation,
//             (voter)
//             (author)
//             (permlink)
//             (weight) )

type VoteOperation struct {
	Voter    string     `json:"voter"`
	Author   string     `json:"author"`
	Permlink string     `json:"permlink"`
	Weight   *types.Int `json:"weight"`
}

// FC_REFLECT( steemit::chain::custom_operation,
//             (required_auths)
//             (id)
//             (data) )

// FC_REFLECT( steemit::chain::custom_json_operation,
//             (required_auths)
//             (required_posting_auths)
//             (id)
//             (json) )

// FC_REFLECT( steemit::chain::limit_order_create_operation,
//             (owner)
//             (orderid)
//             (amount_to_sell)
//             (min_to_receive)
//             (fill_or_kill)
//             (expiration) )

type LimitOrderCreateOperation struct {
	Owner        string      `json:"owner"`
	OrderID      uint32      `json:"orderid"`
	AmountToSell string      `json:"amount_to_sell"`
	MinToReceive string      `json:"min_to_receive"`
	FillOrKill   bool        `json:"fill_or_kill"`
	Expiration   *types.Time `json:"expiration"`
}

// FC_REFLECT( steemit::chain::fill_order_operation,
//             (owner)
//             (orderid)
//             (pays)
//             (receives) );

type FillOrderOperation struct {
	Owner    string `json:"owner"`
	OrderID  string `json:"orderid"`
	Pays     string `json:"pays"`
	Receives string `json:"receives"`
}

// FC_REFLECT( steemit::chain::limit_order_cancel_operation,
//             (owner)
//             (orderid) )

type LimitOrderCancelOperation struct {
	Owner   string `json:"owner"`
	OrderID uint32 `json:"orderid"`
}

// FC_REFLECT( steemit::chain::delete_comment_operation,
//             (author)
//             (permlink) )

type DeleteCommentOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
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

// Operation represents an operation stored in a transaction.
type Operation struct {
	// Type contains the string representation of the operation.
	Type string
	// Body contains one of the operation objects depending on the type.
	Body interface{}
}

func (op *Operation) UnmarshalJSON(data []byte) error {
	// The operation object is [opType, opBody].
	raw := make([]json.RawMessage, 2)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if len(raw) != 2 {
		return errors.Errorf("invalid operation object: %v\n", string(data))
	}

	// Unmarshal opType.
	var opType string
	if err := json.Unmarshal(raw[0], &opType); err != nil {
		return errors.Wrapf(err, "failed to unmarshal Operation.Type: %v", string(raw[0]))
	}

	// Unmarshal opBody.
	bodyTemplate, ok := opBodyObjects[opType]
	if !ok {
		bodyTemplate = map[string]interface{}{}
	}
	opBody := reflect.New(reflect.Indirect(reflect.ValueOf(bodyTemplate)).Type()).Interface()
	if err := json.Unmarshal(raw[1], opBody); err != nil {
		return errors.Wrapf(err, "failed to unmarshal Operation.Body: %v", string(raw[1]))
	}

	// Update fields.
	op.Type = opType
	op.Body = opBody
	return nil
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
