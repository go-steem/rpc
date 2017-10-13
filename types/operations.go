package types

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/golos-go/encoding/transaction"
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

func (op *ConvertOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeConvert.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.RequestID)
	enc.EncodeMoney(op.Amount)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::feed_publish_operation,
//             (publisher)
//             (exchange_rate) )

type FeedPublishOperation struct {
	Publisher    string   `json:"publisher"`
	ExchangeRate ExchRate `json:"exchange_rate"`
}

type ExchRate struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
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
	enc.EncodeMoney(op.Amount)
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
	enc.EncodeMoney(op.Amount)
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

func (op *WithdrawVestingOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeWithdrawVesting.Code()))
	enc.Encode(op.Account)
	enc.EncodeMoney(op.VestingShares)
	return enc.Err()
}

// FC_REFLECT( steemit::chain::set_withdraw_vesting_route_operation,
//             (from_account)
//             (to_account)
//             (percent)
//             (auto_vest) )

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

func (op *AccountWitnessVoteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessVote.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Witness)
	enc.EncodeBool(op.Approve)
	return enc.Err()
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

func (op *AccountWitnessProxyOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeAccountWitnessProxy.Code()))
	enc.Encode(op.Account)
	enc.Encode(op.Proxy)
	return enc.Err()
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
	ParentAuthor   string `json:"parent_author"`
	ParentPermlink string `json:"parent_permlink"`
	Author         string `json:"author"`
	Permlink       string `json:"permlink"`
	Title          string `json:"title"`
	Body           string `json:"body"`
	JsonMetadata   string `json:"json_metadata"`
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
	if !op.IsStoryOperation() {
		enc.Encode(op.ParentAuthor)
	} else {
		enc.Encode(byte(0))
	}
	enc.Encode(op.ParentPermlink)
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.Encode(op.Title)
	enc.Encode(op.Body)
	enc.Encode(op.JsonMetadata)
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

func (op *LimitOrderCreateOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCreate.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	enc.EncodeMoney(op.AmountToSell)
	enc.EncodeMoney(op.MinToReceive)
	enc.EncodeBool(op.FillOrKill)
	enc.Encode(op.Expiration)
	return enc.Err()
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

func (op *LimitOrderCancelOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeLimitOrderCancel.Code()))
	enc.Encode(op.Owner)
	enc.Encode(op.OrderID)
	return enc.Err()
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

func (op *CommentOptionsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCommentOptions.Code()))
	enc.Encode(op.Author)
	enc.Encode(op.Permlink)
	enc.EncodeMoney(op.MaxAcceptedPayout)
	enc.Encode(op.PercentSteemDollars)
	enc.EncodeBool(op.AllowVotes)
	enc.EncodeBool(op.AllowCurationRewards)
	enc.Encode(byte(0))
	return enc.Err()
}

type Authority struct {
	AccountAuths    StringInt64Map `json:"account_auths"`
	KeyAuths        StringInt64Map `json:"key_auths"`
	WeightThreshold uint32         `json:"weight_threshold"`
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

// FC_REFLECT( steemit::chain::witness_update_operation,
//             (owner)
//             (url)
//             (block_signing_key)
//             (props)
//             (fee) )
type WitnessUpdateOperation struct {
	Owner           string           `json:"owner"`
	Url             string           `json:"url"`
	BlockSigningKey string           `json:"block_signing_key"`
	Props           *ChainProperties `json:"props"`
	Fee             string           `json:"fee"`
}

func (op *WitnessUpdateOperation) Type() OpType {
	return TypeWitnessUpdate
}

func (op *WitnessUpdateOperation) Data() interface{} {
	return op
}

type CustomOperation struct {
	RequiredAuths []string `json:"required_auths"`
	Id            uint16   `json:"id"`
	Datas         []byte   `json:"data"`
}

func (op *CustomOperation) Type() OpType {
	return TypeCustom
}

func (op *CustomOperation) Data() interface{} {
	return op
}

type SetWithdrawVestingRouteOperation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Percent     uint16 `json:"percent"`
	AutoVest    bool   `json:"auto_vest"`
}

func (op *SetWithdrawVestingRouteOperation) Type() OpType {
	return TypeSetWithdrawVestingRoute
}

func (op *SetWithdrawVestingRouteOperation) Data() interface{} {
	return op
}

func (op *SetWithdrawVestingRouteOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeSetWithdrawVestingRoute.Code()))
	enc.Encode(op.FromAccount)
	enc.Encode(op.ToAccount)
	enc.Encode(op.Percent)
	enc.EncodeBool(op.AutoVest)
	return enc.Err()
}

type LimitOrderCreate2Operation struct {
	Qwner        string   `json:"owner"`
	Orderid      uint32   `json:"orderid"`
	AmountToSell string   `json:"amount_to_sell"`
	ExchangeRate ExchRate `json:"exchange_rate"`
	FillOrKill   bool     `json:"fill_or_kill"`
	Expiration   uint32   `json:"expiration"`
}

func (op *LimitOrderCreate2Operation) Type() OpType {
	return TypeLimitOrderCreate2
}

func (op *LimitOrderCreate2Operation) Data() interface{} {
	return op
}

type ChallengeAuthorityOperation struct {
	Challenger   string `json:"challenger"`
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

func (op *ChallengeAuthorityOperation) Type() OpType {
	return TypeChallengeAuthority
}

func (op *ChallengeAuthorityOperation) Data() interface{} {
	return op
}

type ProveAuthorityOperation struct {
	Challenged   string `json:"challenged"`
	RequireOwner bool   `json:"require_owner"`
}

func (op *ProveAuthorityOperation) Type() OpType {
	return TypeProveAuthority
}

func (op *ProveAuthorityOperation) Data() interface{} {
	return op
}

type RequestAccountRecoveryOperation struct {
	RecoveryAccount   string        `json:"recovery_account"`
	AccountToRecover  string        `json:"account_to_recover"`
	NewOwnerAuthority []interface{} `json:"new_owner_authority"`
	Extensions        []interface{} `json:"extensions"`
}

func (op *RequestAccountRecoveryOperation) Type() OpType {
	return TypeRequestAccountRecovery
}

func (op *RequestAccountRecoveryOperation) Data() interface{} {
	return op
}

type RecoverAccountOperation struct {
	AccountToRecover     string        `json:"account_to_recover"`
	NewOwnerAuthority    string        `json:"new_owner_authority"`
	RecentOwnerAuthority string        `json:"recent_owner_authority"`
	Extensions           []interface{} `json:"extensions"`
}

func (op *RecoverAccountOperation) Type() OpType {
	return TypeRecoverAccount
}

func (op *RecoverAccountOperation) Data() interface{} {
	return op
}

type ChangeRecoveryAccountOperation struct {
	AccountToRecover   string        `json:"account_to_recover"`
	NewRecoveryAccount string        `json:"new_recovery_account"`
	Extensions         []interface{} `json:"extensions"`
}

func (op *ChangeRecoveryAccountOperation) Type() OpType {
	return TypeChangeRecoveryAccount
}

func (op *ChangeRecoveryAccountOperation) Data() interface{} {
	return op
}

func (op *ChangeRecoveryAccountOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeChangeRecoveryAccount.Code()))
	enc.Encode(op.AccountToRecover)
	enc.Encode(op.NewRecoveryAccount)
	enc.Encode(byte(0))
	return enc.Err()
}

type EscrowTransferOperation struct {
	From                 string `json:"from"`
	To                   string `json:"to"`
	SbdAmount            string `json:"sbd_amount"`
	SteemAmount          string `json:"steem_amount"`
	EscrowId             uint32 `json:"escrow_id"`
	Agent                string `json:"agent"`
	Fee                  string `json:"fee"`
	JsonMeta             string `json:"json_meta"`
	RatificationDeadline uint32 `json:"ratification_deadline"`
	EscrowExpiration     uint32 `json:"escrow_expiration"`
}

func (op *EscrowTransferOperation) Type() OpType {
	return TypeEscrowTransfer
}

func (op *EscrowTransferOperation) Data() interface{} {
	return op
}

type EscrowDisputeOperation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowId uint32 `json:"escrow_id"`
}

func (op *EscrowDisputeOperation) Type() OpType {
	return TypeEscrowDispute
}

func (op *EscrowDisputeOperation) Data() interface{} {
	return op
}

type EscrowReleaseOperation struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Agent       string `json:"agent"`
	Who         string `json:"who"`
	Receiver    string `json:"receiver"`
	EscrowId    uint32 `json:"escrow_id"`
	SbdAmount   string `json:"sbd_amount"`
	SteemAmount string `json:"steem_amount"`
}

func (op *EscrowReleaseOperation) Type() OpType {
	return TypeEscrowRelease
}

func (op *EscrowReleaseOperation) Data() interface{} {
	return op
}

type POW2Operation struct {
	Input      *POW2Input `json:"input"`
	PowSummary uint32     `json:"pow_summary"`
}

type POW2Input struct {
	WorkerAccount string `json:"worker_account"`
	PrevBlock     []byte `json:"prev_block"`
	Nonce         uint64 `json:"nonce"`
}

func (op *POW2Operation) Type() OpType {
	return TypePOW2
}

func (op *POW2Operation) Data() interface{} {
	return op
}

type EscrowApproveOperation struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Agent    string `json:"agent"`
	Who      string `json:"who"`
	EscrowId uint32 `json:"escrow_id"`
	Approve  bool   `json:"approve"`
}

func (op *EscrowApproveOperation) Type() OpType {
	return TypeEscrowApprove
}

func (op *EscrowApproveOperation) Data() interface{} {
	return op
}

type TransferToSavingsOperation struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
	Memo   string `json:"memo"`
}

func (op *TransferToSavingsOperation) Type() OpType {
	return TypeTransferToSavings
}

func (op *TransferToSavingsOperation) Data() interface{} {
	return op
}

func (op *TransferToSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferToSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

type TransferFromSavingsOperation struct {
	From      string `json:"from"`
	RequestId uint32 `json:"request_id"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	Memo      string `json:"memo"`
}

func (op *TransferFromSavingsOperation) Type() OpType {
	return TypeTransferFromSavings
}

func (op *TransferFromSavingsOperation) Data() interface{} {
	return op
}

func (op *TransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestId)
	enc.Encode(op.To)
	enc.EncodeMoney(op.Amount)
	enc.Encode(op.Memo)
	return enc.Err()
}

type CancelTransferFromSavingsOperation struct {
	From      string `json:"from"`
	RequestId uint32 `json:"request_id"`
}

func (op *CancelTransferFromSavingsOperation) Type() OpType {
	return TypeCancelTransferFromSavings
}

func (op *CancelTransferFromSavingsOperation) Data() interface{} {
	return op
}

func (op *CancelTransferFromSavingsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCancelTransferFromSavings.Code()))
	enc.Encode(op.From)
	enc.Encode(op.RequestId)
	return enc.Err()
}

type CustomBinaryOperation struct {
	RequiredOwnerAuths   []string `json:"required_owner_auths"`
	RequiredActiveAuths  []string `json:"required_active_auths"`
	RequiredPostingAuths []string `json:"required_posting_auths"`
	RequiredAuths        []string `json:"required_auths"`
	Id                   string   `json:"id"`
	Datas                []byte   `json:"data"`
}

func (op *CustomBinaryOperation) Type() OpType {
	return TypeCustomBinary
}

func (op *CustomBinaryOperation) Data() interface{} {
	return op
}

type DeclineVotingRightsOperation struct {
	Account string `json:"account"`
	Decline bool   `json:"decline"`
}

func (op *DeclineVotingRightsOperation) Type() OpType {
	return TypeDeclineVotingRights
}

func (op *DeclineVotingRightsOperation) Data() interface{} {
	return op
}

func (op *DeclineVotingRightsOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeDeclineVotingRights.Code()))
	enc.Encode(op.Account)
	enc.EncodeBool(op.Decline)
	return enc.Err()
}

type ResetAccountOperation struct {
	ResetAccount      string `json:"reset_account"`
	AccountToReset    string `json:"Account_to_reset"`
	NewOwnerAuthority string `json:"new_owner_authority"`
}

func (op *ResetAccountOperation) Type() OpType {
	return TypeResetAccount
}

func (op *ResetAccountOperation) Data() interface{} {
	return op
}

type SetResetAccountOperation struct {
	Account             string `json:"account"`
	CurrentResetAccount string `json:"current_reset_account"`
	ResetAccount        string `json:"reset_account"`
}

func (op *SetResetAccountOperation) Type() OpType {
	return TypeSetResetAccount
}

func (op *SetResetAccountOperation) Data() interface{} {
	return op
}

type ClaimRewardBalanceOperation struct {
	Account     string `json:"account"`
	RewardSteem string `json:"reward_steem"`
	RewardSbd   string `json:"reward_sbd"`
	RewardVests string `json:"reward_vests"`
}

func (op *ClaimRewardBalanceOperation) Type() OpType {
	return TypeClaimRewardBalance
}

func (op *ClaimRewardBalanceOperation) Data() interface{} {
	return op
}

type DelegateVestingSharesOperation struct {
	Delegator     string `json:"delegator"`
	Delegatee     string `json:"delegatee"`
	VestingShares string `json:"vesting_shares"`
}

func (op *DelegateVestingSharesOperation) Type() OpType {
	return TypeDelegateVestingShares
}

func (op *DelegateVestingSharesOperation) Data() interface{} {
	return op
}

type AccountCreateWithDelegationOperation struct {
	Fee            string        `json:"fee"`
	Delegation     string        `json:"delegation"`
	Creator        string        `json:"creator"`
	NewAccountName string        `json:"new_account_name"`
	Owner          string        `json:"owner"`
	Active         string        `json:"active"`
	Posting        string        `json:"posting"`
	MemoKey        string        `json:"memo_key"`
	JsonMetadata   string        `json:"json_metadata"`
	Extensions     []interface{} `json:"extensions"`
}

func (op *AccountCreateWithDelegationOperation) Type() OpType {
	return TypeAccountCreateWithDelegation
}

func (op *AccountCreateWithDelegationOperation) Data() interface{} {
	return op
}

type FillConvertRequestOperation struct {
	Owner     string `json:"owner"`
	Requestid uint32 `json:"requestid"`
	AmountIn  string `json:"amount_in"`
	AmountOut string `json:"amount_out"`
}

func (op *FillConvertRequestOperation) Type() OpType {
	return TypeFillConvertRequest
}

func (op *FillConvertRequestOperation) Data() interface{} {
	return op
}

type AuthorRewardOperation struct {
	Author        string `json:"author"`
	Permlink      string `json:"permlink"`
	SbdPayout     string `json:"sbd_payout"`
	SteemPayout   string `json:"steem_payout"`
	VestingPayout string `json:"vesting_payout"`
}

func (op *AuthorRewardOperation) Type() OpType {
	return TypeAuthorReward
}

func (op *AuthorRewardOperation) Data() interface{} {
	return op
}

type CurationRewardOperation struct {
	Curator         string `json:"curator"`
	Reward          string `json:"reward"`
	CommentAuthor   string `json:"comment_author"`
	CommentPermlink string `json:"comment_permlink"`
}

func (op *CurationRewardOperation) Type() OpType {
	return TypeCurationReward
}

func (op *CurationRewardOperation) Data() interface{} {
	return op
}

type CommentRewardOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Payout   string `json:"payout"`
}

func (op *CommentRewardOperation) Type() OpType {
	return TypeCommentReward
}

func (op *CommentRewardOperation) Data() interface{} {
	return op
}

type LiquidityRewardOperation struct {
	Owner  string `json:"owner"`
	Payout string `json:"payout"`
}

func (op *LiquidityRewardOperation) Type() OpType {
	return TypeLiquidityReward
}

func (op *LiquidityRewardOperation) Data() interface{} {
	return op
}

type InterestOperation struct {
	Owner    string `json:"owner"`
	Interest string `json:"interest"`
}

func (op *InterestOperation) Type() OpType {
	return TypeInterest
}

func (op *InterestOperation) Data() interface{} {
	return op
}

type FillVestingWithdrawOperation struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
	Withdrawn   string `json:"withdrawn"`
	Deposited   string `json:"deposited"`
}

func (op *FillVestingWithdrawOperation) Type() OpType {
	return TypeFillVestingWithdraw
}

func (op *FillVestingWithdrawOperation) Data() interface{} {
	return op
}

type FillOrderOperation struct {
	CurrentOwner   string `json:"current_owner"`
	CurrentOrderid uint32 `json:"current_orderid"`
	CurrentPays    string `json:"current_pays"`
	OpenOwner      string `json:"open_owner"`
	OpenOrderid    uint32 `json:"open_orderid"`
	OpenPays       string `json:"open_pays"`
}

func (op *FillOrderOperation) Type() OpType {
	return TypeFillOrder
}

func (op *FillOrderOperation) Data() interface{} {
	return op
}

type ShutdownWitnessOperation struct {
	Owner string `json:"owner"`
}

func (op *ShutdownWitnessOperation) Type() OpType {
	return TypeShutdownWitness
}

func (op *ShutdownWitnessOperation) Data() interface{} {
	return op
}

type FillTransferFromSavingsOperation struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	RequestId uint32 `json:"request_id"`
	Memo      string `json:"memo"`
}

func (op *FillTransferFromSavingsOperation) Type() OpType {
	return TypeFillTransferFromSavings
}

func (op *FillTransferFromSavingsOperation) Data() interface{} {
	return op
}

type HardforkOperation struct {
	HardforkId uint32 `json:"hardfork_id"`
}

func (op *HardforkOperation) Type() OpType {
	return TypeHardfork
}

func (op *HardforkOperation) Data() interface{} {
	return op
}

type CommentPayoutUpdateOperation struct {
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

func (op *CommentPayoutUpdateOperation) Type() OpType {
	return TypeCommentPayoutUpdate
}

func (op *CommentPayoutUpdateOperation) Data() interface{} {
	return op
}

type ReturnVestingDelegationOperation struct {
	Account       string `json:"account"`
	VestingShares string `json:"vesting_shares"`
}

func (op *ReturnVestingDelegationOperation) Type() OpType {
	return TypeReturnVestingDelegation
}

func (op *ReturnVestingDelegationOperation) Data() interface{} {
	return op
}

type CommentBenefactorRewardOperation struct {
	Benefactor string `json:"benefactor"`
	Author     string `json:"author"`
	Permlink   string `json:"permlink"`
	Reward     string `json:"reward"`
}

func (op *CommentBenefactorRewardOperation) Type() OpType {
	return TypeCommentBenefactorReward
}

func (op *CommentBenefactorRewardOperation) Data() interface{} {
	return op
}
