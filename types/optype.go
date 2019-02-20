package types

// OpType represents a Golos operation type, i.e. vote, comment, pow and so on.
type OpType string

// Code returns the operation code associated with the given operation type.
func (kind OpType) Code() uint16 {
	return opCodes[kind]
}

const (
	TypeVote                        OpType = "vote"
	TypeComment                     OpType = "comment"
	TypeTransfer                    OpType = "transfer"
	TypeTransferToVesting           OpType = "transfer_to_vesting"
	TypeWithdrawVesting             OpType = "withdraw_vesting"
	TypeLimitOrderCreate            OpType = "limit_order_create"
	TypeLimitOrderCancel            OpType = "limit_order_cancel"
	TypeFeedPublish                 OpType = "feed_publish"
	TypeConvert                     OpType = "convert"
	TypeAccountCreate               OpType = "account_create"
	TypeAccountUpdate               OpType = "account_update"
	TypeWitnessUpdate               OpType = "witness_update"
	TypeAccountWitnessVote          OpType = "account_witness_vote"
	TypeAccountWitnessProxy         OpType = "account_witness_proxy"
	TypePOW                         OpType = "pow"
	TypeCustom                      OpType = "custom"
	TypeReportOverProduction        OpType = "report_over_production"
	TypeDeleteComment               OpType = "delete_comment"
	TypeCustomJSON                  OpType = "custom_json"
	TypeCommentOptions              OpType = "comment_options"
	TypeSetWithdrawVestingRoute     OpType = "set_withdraw_vesting_route"
	TypeLimitOrderCreate2           OpType = "limit_order_create2"
	TypeClaimAccount                OpType = "claim_account"
	TypeCreateClaimedAccount        OpType = "create_claimed_account"
	TypeRequestAccountRecovery      OpType = "request_account_recovery"
	TypeRecoverAccount              OpType = "recover_account"
	TypeChangeRecoveryAccount       OpType = "change_recovery_account"
	TypeEscrowTransfer              OpType = "escrow_transfer"
	TypeEscrowDispute               OpType = "escrow_dispute"
	TypeEscrowRelease               OpType = "escrow_release"
	TypePOW2                        OpType = "pow2"
	TypeEscrowApprove               OpType = "escrow_approve"
	TypeTransferToSavings           OpType = "transfer_to_savings"
	TypeTransferFromSavings         OpType = "transfer_from_savings"
	TypeCancelTransferFromSavings   OpType = "cancel_transfer_from_savings"
	TypeCustomBinary                OpType = "custom_binary"
	TypeDeclineVotingRights         OpType = "decline_voting_rights"
	TypeResetAccount                OpType = "reset_account"
	TypeSetResetAccount             OpType = "set_reset_account"
	TypeClaimRewardBalance          OpType = "claim_reward_balance"
	TypeDelegateVestingShares       OpType = "delegate_vesting_shares"
	TypeAccountCreateWithDelegation OpType = "account_create_with_delegation"
	TypeFillConvertRequest          OpType = "fill_convert_request"       //Virtual Operation
	TypeAuthorReward                OpType = "author_reward"              //Virtual Operation
	TypeCurationReward              OpType = "curation_reward"            //Virtual Operation
	TypeCommentReward               OpType = "comment_reward"             //Virtual Operation
	TypeLiquidityReward             OpType = "liquidity_reward"           //Virtual Operation
	TypeInterest                    OpType = "interest"                   //Virtual Operation
	TypeFillVestingWithdraw         OpType = "fill_vesting_withdraw"      //Virtual Operation
	TypeFillOrder                   OpType = "fill_order"                 //Virtual Operation
	TypeShutdownWitness             OpType = "shutdown_witness"           //Virtual Operation
	TypeFillTransferFromSavings     OpType = "fill_transfer_from_savings" //Virtual Operation
	TypeHardfork                    OpType = "hardfork"                   //Virtual Operation
	TypeCommentPayoutUpdate         OpType = "comment_payout_update"      //Virtual Operation
	TypeReturnVestingDelegation     OpType = "return_vesting_delegation"  //Virtual Operation
	TypeCommentBenefactorReward     OpType = "comment_benefactor_reward"  //Virtual Operation
)

var opTypes = [...]OpType{
	TypeVote,
	TypeComment,
	TypeTransfer,
	TypeTransferToVesting,
	TypeWithdrawVesting,
	TypeLimitOrderCreate,
	TypeLimitOrderCancel,
	TypeFeedPublish,
	TypeConvert,
	TypeAccountCreate,
	TypeAccountUpdate,
	TypeWitnessUpdate,
	TypeAccountWitnessVote,
	TypeAccountWitnessProxy,
	TypePOW,
	TypeCustom,
	TypeReportOverProduction,
	TypeDeleteComment,
	TypeCustomJSON,
	TypeCommentOptions,
	TypeSetWithdrawVestingRoute,
	TypeLimitOrderCreate2,
	TypeClaimAccount,
	TypeCreateClaimedAccount,
	TypeRequestAccountRecovery,
	TypeRecoverAccount,
	TypeChangeRecoveryAccount,
	TypeEscrowTransfer,
	TypeEscrowDispute,
	TypeEscrowRelease,
	TypePOW2,
	TypeEscrowApprove,
	TypeTransferToSavings,
	TypeTransferFromSavings,
	TypeCancelTransferFromSavings,
	TypeCustomBinary,
	TypeDeclineVotingRights,
	TypeResetAccount,
	TypeSetResetAccount,
	TypeClaimRewardBalance,
	TypeDelegateVestingShares,
	TypeAccountCreateWithDelegation,
	TypeFillConvertRequest,      //Virtual Operation
	TypeAuthorReward,            //Virtual Operation
	TypeCurationReward,          //Virtual Operation
	TypeCommentReward,           //Virtual Operation
	TypeLiquidityReward,         //Virtual Operation
	TypeInterest,                //Virtual Operation
	TypeFillVestingWithdraw,     //Virtual Operation
	TypeFillOrder,               //Virtual Operation
	TypeShutdownWitness,         //Virtual Operation
	TypeFillTransferFromSavings, //Virtual Operation
	TypeHardfork,                //Virtual Operation
	TypeCommentPayoutUpdate,     //Virtual Operation
	TypeReturnVestingDelegation, //Virtual Operation
	TypeCommentBenefactorReward, //Virtual Operation
}

// opCodes keeps mapping operation type -> operation code.
var opCodes map[OpType]uint16

func init() {
	opCodes = make(map[OpType]uint16, len(opTypes))
	for i, opType := range opTypes {
		opCodes[opType] = uint16(i)
	}
}
