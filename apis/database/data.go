package database

import (
	// Stdlib
	"encoding/json"
	"strconv"
	"strings"

	// RPC
	"github.com/asuleymanov/golos-go/types"
)

type Config struct {
	SteemitBuildTestnet                   bool   `json:"STEEMIT_BUILD_TESTNET"`
	GrapheneCurrentDBVersion              string `json:"GRAPHENE_CURRENT_DB_VERSION"`
	SbdSymbol                             uint   `json:"SBD_SYMBOL"`
	Steemit100Percent                     uint   `json:"STEEMIT_100_PERCENT"`
	Steemit1Percent                       uint   `json:"STEEMIT_1_PERCENT"`
	SteemitAddressPrefix                  string `json:"STEEMIT_ADDRESS_PREFIX"`
	SteemitAprPercentMultiplyPerBlock     string `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_BLOCK"`
	SteemitAprPercentMultiplyPerHour      string `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_HOUR"`
	SteemitAprPercentMultiplyPerRound     string `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_ROUND"`
	SteemitAprPercentShiftPerBlock        uint   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_BLOCK"`
	SteemitAprPercentShiftPerHour         uint   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_HOUR"`
	SteemitAprPercentShiftPerRound        uint   `json:"STEEMIT_APR_PERCENT_SHIFT_PER_ROUND"`
	SteemitBandwidthAverageWindowSeconds  uint   `json:"STEEMIT_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	SteemitBandwidthPrecision             uint   `json:"STEEMIT_BANDWIDTH_PRECISION"`
	SteemitBlockchainPrecision            uint   `json:"STEEMIT_BLOCKCHAIN_PRECISION"`
	SteemitBlockchainPrecisionDigits      uint   `json:"STEEMIT_BLOCKCHAIN_PRECISION_DIGITS"`
	SteemitBlockchainHardforkVersion      string `json:"STEEMIT_BLOCKCHAIN_HARDFORK_VERSION"`
	SteemitBlockchainVersion              string `json:"STEEMIT_BLOCKCHAIN_VERSION"`
	SteemitBlockInterval                  uint   `json:"STEEMIT_BLOCK_INTERVAL"`
	SteemitBlocksPerDay                   uint   `json:"STEEMIT_BLOCKS_PER_DAY"`
	SteemitBlocksPerHour                  uint   `json:"STEEMIT_BLOCKS_PER_HOUR"`
	SteemitBlocksPerYear                  uint   `json:"STEEMIT_BLOCKS_PER_YEAR"`
	SteemitCashoutWindowSeconds           uint   `json:"STEEMIT_CASHOUT_WINDOW_SECONDS"`
	SteemitChainId                        string `json:"STEEMIT_CHAIN_ID"`
	SteemitContentAprPercent              uint   `json:"STEEMIT_CONTENT_APR_PERCENT"`
	SteemitConversionDelay                string `json:"STEEMIT_CONVERSION_DELAY"`
	SteemitCurateAprPercent               uint   `json:"STEEMIT_CURATE_APR_PERCENT"`
	SteemitDefaultSbdInterestRate         uint   `json:"STEEMIT_DEFAULT_SBD_INTEREST_RATE"`
	SteemitFeedHistoryWindow              uint   `json:"STEEMIT_FEED_HISTORY_WINDOW"`
	SteemitFeedIntervalBlocks             uint   `json:"STEEMIT_FEED_INTERVAL_BLOCKS"`
	SteemitFreeTransactionsWithNewAccount uint   `json:"STEEMIT_FREE_TRANSACTIONS_WITH_NEW_ACCOUNT"`
	SteemitGenesisTime                    string `json:"STEEMIT_GENESIS_TIME"`
	SteemitHardforkRequiredWitnesses      uint   `json:"STEEMIT_HARDFORK_REQUIRED_WITNESSES"`
	SteemitInitMinerName                  string `json:"STEEMIT_INIT_MINER_NAME"`
	SteemitInitPublicKeyStr               string `json:"STEEMIT_INIT_PUBLIC_KEY_STR"`
	SteemitInitSupply                     string `json:"STEEMIT_INIT_SUPPLY"`
	SteemitInitTime                       string `json:"STEEMIT_INIT_TIME"`
	SteemitIrreversibleThreshold          uint   `json:"STEEMIT_IRREVERSIBLE_THRESHOLD"`
	SteemitLiquidityAprPercent            uint   `json:"STEEMIT_LIQUIDITY_APR_PERCENT"`
	SteemitLiquidityRewardBlocks          uint   `json:"STEEMIT_LIQUIDITY_REWARD_BLOCKS"`
	SteemitLiquidityRewardPeriodSec       uint   `json:"STEEMIT_LIQUIDITY_REWARD_PERIOD_SEC"`
	SteemitLiquidityTimeoutSec            string `json:"STEEMIT_LIQUIDITY_TIMEOUT_SEC"`
	SteemitMaxAccountNameLength           uint   `json:"STEEMIT_MAX_ACCOUNT_NAME_LENGTH"`
	SteemitMaxAccountWitnessVotes         uint   `json:"STEEMIT_MAX_ACCOUNT_WITNESS_VOTES"`
	SteemitMaxAssetWhitelistAuthorities   uint   `json:"STEEMIT_MAX_ASSET_WHITELIST_AUTHORITIES"`
	SteemitMaxAuthorityMembership         uint   `json:"STEEMIT_MAX_AUTHORITY_MEMBERSHIP"`
	SteemitMaxBlockSize                   uint   `json:"STEEMIT_MAX_BLOCK_SIZE"`
	SteemitMaxCashoutWindowSeconds        uint   `json:"STEEMIT_MAX_CASHOUT_WINDOW_SECONDS"`
	SteemitMaxCommentDepth                uint   `json:"STEEMIT_MAX_COMMENT_DEPTH"`
	SteemitMaxFeedAge                     string `json:"STEEMIT_MAX_FEED_AGE"`
	SteemitMaxInstanceId                  string `json:"STEEMIT_MAX_INSTANCE_ID"`
	SteemitMaxMemoSize                    uint   `json:"STEEMIT_MAX_MEMO_SIZE"`
	SteemitMaxWitnesses                   uint   `json:"STEEMIT_MAX_WITNESSES"`
	SteemitMaxMinerWitnesses              uint   `json:"STEEMIT_MAX_MINER_WITNESSES"`
	SteemitMaxProxyRecursionDepth         uint   `json:"STEEMIT_MAX_PROXY_RECURSION_DEPTH"`
	SteemitMaxRationDecayRate             uint   `json:"STEEMIT_MAX_RATION_DECAY_RATE"`
	SteemitMaxReserveRatio                uint   `json:"STEEMIT_MAX_RESERVE_RATIO"`
	SteemitMaxRunnerWitnesses             uint   `json:"STEEMIT_MAX_RUNNER_WITNESSES"`
	SteemitMaxShareSupply                 string `json:"STEEMIT_MAX_SHARE_SUPPLY"`
	SteemitMaxSigCheckDepth               uint   `json:"STEEMIT_MAX_SIG_CHECK_DEPTH"`
	SteemitMaxTimeUntilExpiration         uint   `json:"STEEMIT_MAX_TIME_UNTIL_EXPIRATION"`
	SteemitMaxTransactionSize             uint   `json:"STEEMIT_MAX_TRANSACTION_SIZE"`
	SteemitMaxUndoHistory                 uint   `json:"STEEMIT_MAX_UNDO_HISTORY"`
	SteemitMaxUrlLength                   uint   `json:"STEEMIT_MAX_URL_LENGTH"`
	SteemitMaxVoteChanges                 uint   `json:"STEEMIT_MAX_VOTE_CHANGES"`
	SteemitMaxVotedWitnesses              uint   `json:"STEEMIT_MAX_VOTED_WITNESSES"`
	SteemitMaxWithdrawRoutes              uint   `json:"STEEMIT_MAX_WITHDRAW_ROUTES"`
	SteemitMaxWitnessUrlLength            uint   `json:"STEEMIT_MAX_WITNESS_URL_LENGTH"`
	SteemitMinAccountCreationFee          uint   `json:"STEEMIT_MIN_ACCOUNT_CREATION_FEE"`
	SteemitMinAccountNameLength           uint   `json:"STEEMIT_MIN_ACCOUNT_NAME_LENGTH"`
	SteemitMinBlockSizeLimit              uint   `json:"STEEMIT_MIN_BLOCK_SIZE_LIMIT"`
	SteemitMinContentReward               string `json:"STEEMIT_MIN_CONTENT_REWARD"`
	SteemitMinCurateReward                string `json:"STEEMIT_MIN_CURATE_REWARD"`
	SteemitMinerAccount                   string `json:"STEEMIT_MINER_ACCOUNT"`
	SteemitMinerPayPercent                uint   `json:"STEEMIT_MINER_PAY_PERCENT"`
	SteemitMinFeeds                       uint   `json:"STEEMIT_MIN_FEEDS"`
	SteemitMiningReward                   string `json:"STEEMIT_MINING_REWARD"`
	SteemitMiningTime                     string `json:"STEEMIT_MINING_TIME"`
	steemitMinLiquidityReward             string `json:"STEEMIT_MIN_LIQUIDITY_REWARD"`
	SteemitMinLiquidityRewardPeriodSec    uint   `json:"STEEMIT_MIN_LIQUIDITY_REWARD_PERIOD_SEC"`
	SteemitMinPayoutSbd                   string `json:"STEEMIT_MIN_PAYOUT_SBD"`
	SteemitMinPowReward                   string `json:"STEEMIT_MIN_POW_REWARD"`
	SteemitMinProducerReward              string `json:"STEEMIT_MIN_PRODUCER_REWARD"`
	SteemitMinRation                      uint   `json:"STEEMIT_MIN_RATION"`
	SteemitMinTransactionExpirationLimit  uint   `json:"STEEMIT_MIN_TRANSACTION_EXPIRATION_LIMIT"`
	SteemitMinTransactionSizeLimit        uint   `json:"STEEMIT_MIN_TRANSACTION_SIZE_LIMIT"`
	SteemitMinUndoHistory                 uint   `json:"STEEMIT_MIN_UNDO_HISTORY"`
	SteemitNullAccount                    string `json:"STEEMIT_NULL_ACCOUNT"`
	SteemitNumInitMiners                  uint   `json:"STEEMIT_NUM_INIT_MINERS"`
	SteemitPowAprPercent                  uint   `json:"STEEMIT_POW_APR_PERCENT"`
	SteemitProducerAprPercent             uint   `json:"STEEMIT_PRODUCER_APR_PERCENT"`
	SteemitProxyToSelfAccount             string `json:"STEEMIT_PROXY_TO_SELF_ACCOUNT"`
	SteemitSbdInterestCompoundIntervalSec uint   `json:"STEEMIT_SBD_INTEREST_COMPOUND_INTERVAL_SEC"`
	SteemitSecondsPerYear                 uint   `json:"STEEMIT_SECONDS_PER_YEAR"`
	SteemitReverseAuctionWindowSeconds    uint   `json:"STEEMIT_REVERSE_AUCTION_WINDOW_SECONDS"`
	SteemitStartMinerVotingBlock          uint   `json:"STEEMIT_START_MINER_VOTING_BLOCK"`
	SteemitStartVestingBlock              uint   `json:"STEEMIT_START_VESTING_BLOCK"`
	SteemitSymbol                         string `json:"STEEMIT_SYMBOL"`
	SteemitTempAccount                    string `json:"STEEMIT_TEMP_ACCOUNT"`
	SteemitUpvoteLockout                  uint   `json:"STEEMIT_UPVOTE_LOCKOUT"`
	SteemitVestingWithdrawIntervals       uint   `json:"STEEMIT_VESTING_WITHDRAW_INTERVALS"`
	SteemitVestingWithdrawIntervalSeconds uint   `json:"STEEMIT_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	SteemitVoteChangeLockoutPeriod        uint   `json:"STEEMIT_VOTE_CHANGE_LOCKOUT_PERIOD"`
	SteemitVoteRegenerationSeconds        uint   `json:"STEEMIT_VOTE_REGENERATION_SECONDS"`
	SteemSymbol                           string `json:"STEEM_SYMBOL"`
	VestsSymbol                           string `json:"VESTS_SYMBOL"`
	BlockchainName                        string `json:"BLOCKCHAIN_NAME"`
}

type DynamicGlobalProperties struct {
	Time                     *types.Time  `json:"time"`
	TotalPow                 *types.Int   `json:"total_pow"`
	NumPowWitnesses          *types.Int   `json:"num_pow_witnesses"`
	CurrentReserveRatio      *types.Int   `json:"current_reserve_ratio"`
	ID                       *types.ID    `json:"id"`
	CurrentSupply            string       `json:"current_supply"`
	CurrentSBDSupply         string       `json:"current_sbd_supply"`
	MaximumBlockSize         *types.Int   `json:"maximum_block_size"`
	RecentSlotsFilled        *types.Int   `json:"recent_slots_filled"`
	CurrentWitness           string       `json:"current_witness"`
	TotalRewardShares2       *types.Int   `json:"total_reward_shares2"`
	AverageBlockSize         *types.Int   `json:"average_block_size"`
	CurrentAslot             *types.Int   `json:"current_aslot"`
	LastIrreversibleBlockNum uint32       `json:"last_irreversible_block_num"`
	TotalVestingShares       string       `json:"total_vesting_shares"`
	TotalVersingFundSteem    string       `json:"total_vesting_fund_steem"`
	HeadBlockID              string       `json:"head_block_id"`
	HeadBlockNumber          types.UInt32 `json:"head_block_number"`
	VirtualSupply            string       `json:"virtual_supply"`
	ConfidentialSupply       string       `json:"confidential_supply"`
	ConfidentialSBDSupply    string       `json:"confidential_sbd_supply"`
	TotalRewardFundSteem     string       `json:"total_reward_fund_steem"`
	TotalActivityFundSteem   string       `json:"total_activity_fund_steem"`
	TotalActivityFundShares  *types.Int   `json:"total_activity_fund_shares"`
	SBDInterestRate          *types.Int   `json:"sbd_interest_rate"`
	MaxVirtualBandwidth      *types.Int   `json:"max_virtual_bandwidth"`
}

type Block struct {
	Number                uint32               `json:"-"`
	Timestamp             *types.Time          `json:"timestamp"`
	Witness               string               `json:"witness"`
	WitnessSignature      string               `json:"witness_signature"`
	TransactionMerkleRoot string               `json:"transaction_merkle_root"`
	Previous              string               `json:"previous"`
	Extensions            [][]interface{}      `json:"extensions"`
	Transactions          []*types.Transaction `json:"transactions"`
}

type Content struct {
	Id                      *types.ID        `json:"id"`
	RootTitle               string           `json:"root_title"`
	Active                  *types.Time      `json:"active"`
	AbsRshares              *types.Int       `json:"abs_rshares"`
	PendingPayoutValue      string           `json:"pending_payout_value"`
	TotalPendingPayoutValue string           `json:"total_pending_payout_value"`
	Category                string           `json:"category"`
	Title                   string           `json:"title"`
	LastUpdate              *types.Time      `json:"last_update"`
	Stats                   string           `json:"stats"`
	Body                    string           `json:"body"`
	Created                 *types.Time      `json:"created"`
	Replies                 []*Content       `json:"replies"`
	Permlink                string           `json:"permlink"`
	JsonMetadata            *ContentMetadata `json:"json_metadata"`
	Children                *types.Int       `json:"children"`
	NetRshares              *types.Int       `json:"net_rshares"`
	URL                     string           `json:"url"`
	ActiveVotes             []*VoteState     `json:"active_votes"`
	ParentPermlink          string           `json:"parent_permlink"`
	CashoutTime             *types.Time      `json:"cashout_time"`
	TotalPayoutValue        string           `json:"total_payout_value"`
	ParentAuthor            string           `json:"parent_author"`
	ChildrenRshares2        *types.Int       `json:"children_rshares2"`
	Author                  string           `json:"author"`
	Depth                   *types.Int       `json:"depth"`
	TotalVoteWeight         *types.Int       `json:"total_vote_weight"`
}

func (content *Content) IsStory() bool {
	return content.ParentAuthor == ""
}

type ContentMetadata struct {
	Flag  bool
	Users []string
	Tags  []string
	Image []string
}

type ContentMetadataRaw struct {
	Users types.StringSlice `json:"users"`
	Tags  types.StringSlice `json:"tags"`
	Image types.StringSlice `json:"image"`
}

func (metadata *ContentMetadata) UnmarshalJSON(data []byte) error {
	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	switch unquoted {
	case "true":
		metadata.Flag = true
		return nil
	case "false":
		metadata.Flag = false
		return nil
	}

	if len(unquoted) == 0 {
		var value ContentMetadata
		metadata = &value
		return nil
	}

	var raw ContentMetadataRaw
	if err := json.NewDecoder(strings.NewReader(unquoted)).Decode(&raw); err != nil {
		return err
	}

	metadata.Users = raw.Users
	metadata.Tags = raw.Tags
	metadata.Image = raw.Image

	return nil
}

type VoteState struct {
	Voter   string      `json:"voter"`
	Weight  *types.Int  `json:"weight"`
	Rshares *types.Int  `json:"rshares"`
	Percent *types.Int  `json:"percent"`
	Time    *types.Time `json:"time"`
}
