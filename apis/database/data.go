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
	SteemitBuildTestnet                   bool       `json:"STEEMIT_BUILD_TESTNET"`
	GrapheneCurrentDBVersion              string     `json:"GRAPHENE_CURRENT_DB_VERSION"`
	SbdSymbol                             *types.Int `json:"SBD_SYMBOL"`
	Steemit100Percent                     *types.Int `json:"STEEMIT_100_PERCENT"`
	Steemit1Percent                       *types.Int `json:"STEEMIT_1_PERCENT"`
	SteemitAddressPrefix                  string     `json:"STEEMIT_ADDRESS_PREFIX"`
	SteemitAprPercentMultiplyPerBlock     string     `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_BLOCK"`
	SteemitAprPercentMultiplyPerHour      string     `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_HOUR"`
	SteemitAprPercentMultiplyPerRound     string     `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_ROUND"`
	SteemitAprPercentShiftPerBlock        *types.Int `json:"STEEMIT_APR_PERCENT_SHIFT_PER_BLOCK"`
	SteemitAprPercentShiftPerHour         *types.Int `json:"STEEMIT_APR_PERCENT_SHIFT_PER_HOUR"`
	SteemitAprPercentShiftPerRound        *types.Int `json:"STEEMIT_APR_PERCENT_SHIFT_PER_ROUND"`
	SteemitBandwidthAverageWindowSeconds  *types.Int `json:"STEEMIT_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	SteemitBandwidthPrecision             *types.Int `json:"STEEMIT_BANDWIDTH_PRECISION"`
	SteemitBlockchainPrecision            *types.Int `json:"STEEMIT_BLOCKCHAIN_PRECISION"`
	SteemitBlockchainPrecisionDigits      *types.Int `json:"STEEMIT_BLOCKCHAIN_PRECISION_DIGITS"`
	SteemitBlockchainHardforkVersion      string     `json:"STEEMIT_BLOCKCHAIN_HARDFORK_VERSION"`
	SteemitBlockchainVersion              string     `json:"STEEMIT_BLOCKCHAIN_VERSION"`
	SteemitBlockInterval                  *types.Int `json:"STEEMIT_BLOCK_INTERVAL"`
	SteemitBlocksPerDay                   *types.Int `json:"STEEMIT_BLOCKS_PER_DAY"`
	SteemitBlocksPerHour                  *types.Int `json:"STEEMIT_BLOCKS_PER_HOUR"`
	SteemitBlocksPerYear                  *types.Int `json:"STEEMIT_BLOCKS_PER_YEAR"`
	SteemitCashoutWindowSeconds           *types.Int `json:"STEEMIT_CASHOUT_WINDOW_SECONDS"`
	SteemitChainId                        string     `json:"STEEMIT_CHAIN_ID"`
	SteemitContentAprPercent              *types.Int `json:"STEEMIT_CONTENT_APR_PERCENT"`
	SteemitConversionDelay                string     `json:"STEEMIT_CONVERSION_DELAY"`
	SteemitCurateAprPercent               *types.Int `json:"STEEMIT_CURATE_APR_PERCENT"`
	SteemitDefaultSbdInterestRate         *types.Int `json:"STEEMIT_DEFAULT_SBD_INTEREST_RATE"`
	SteemitFeedHistoryWindow              *types.Int `json:"STEEMIT_FEED_HISTORY_WINDOW"`
	SteemitFeedIntervalBlocks             *types.Int `json:"STEEMIT_FEED_INTERVAL_BLOCKS"`
	SteemitFreeTransactionsWithNewAccount *types.Int `json:"STEEMIT_FREE_TRANSACTIONS_WITH_NEW_ACCOUNT"`
	SteemitGenesisTime                    string     `json:"STEEMIT_GENESIS_TIME"`
	SteemitHardforkRequiredWitnesses      *types.Int `json:"STEEMIT_HARDFORK_REQUIRED_WITNESSES"`
	SteemitInitMinerName                  string     `json:"STEEMIT_INIT_MINER_NAME"`
	SteemitInitPublicKeyStr               string     `json:"STEEMIT_INIT_PUBLIC_KEY_STR"`
	SteemitInitSupply                     string     `json:"STEEMIT_INIT_SUPPLY"`
	SteemitInitTime                       string     `json:"STEEMIT_INIT_TIME"`
	SteemitIrreversibleThreshold          *types.Int `json:"STEEMIT_IRREVERSIBLE_THRESHOLD"`
	SteemitLiquidityAprPercent            *types.Int `json:"STEEMIT_LIQUIDITY_APR_PERCENT"`
	SteemitLiquidityRewardBlocks          *types.Int `json:"STEEMIT_LIQUIDITY_REWARD_BLOCKS"`
	SteemitLiquidityRewardPeriodSec       *types.Int `json:"STEEMIT_LIQUIDITY_REWARD_PERIOD_SEC"`
	SteemitLiquidityTimeoutSec            string     `json:"STEEMIT_LIQUIDITY_TIMEOUT_SEC"`
	SteemitMaxAccountNameLength           *types.Int `json:"STEEMIT_MAX_ACCOUNT_NAME_LENGTH"`
	SteemitMaxAccountWitnessVotes         *types.Int `json:"STEEMIT_MAX_ACCOUNT_WITNESS_VOTES"`
	SteemitMaxAssetWhitelistAuthorities   *types.Int `json:"STEEMIT_MAX_ASSET_WHITELIST_AUTHORITIES"`
	SteemitMaxAuthorityMembership         *types.Int `json:"STEEMIT_MAX_AUTHORITY_MEMBERSHIP"`
	SteemitMaxBlockSize                   *types.Int `json:"STEEMIT_MAX_BLOCK_SIZE"`
	SteemitMaxCashoutWindowSeconds        *types.Int `json:"STEEMIT_MAX_CASHOUT_WINDOW_SECONDS"`
	SteemitMaxCommentDepth                *types.Int `json:"STEEMIT_MAX_COMMENT_DEPTH"`
	SteemitMaxFeedAge                     string     `json:"STEEMIT_MAX_FEED_AGE"`
	SteemitMaxInstanceId                  string     `json:"STEEMIT_MAX_INSTANCE_ID"`
	SteemitMaxMemoSize                    *types.Int `json:"STEEMIT_MAX_MEMO_SIZE"`
	SteemitMaxWitnesses                   *types.Int `json:"STEEMIT_MAX_WITNESSES"`
	SteemitMaxMinerWitnesses              *types.Int `json:"STEEMIT_MAX_MINER_WITNESSES"`
	SteemitMaxProxyRecursionDepth         *types.Int `json:"STEEMIT_MAX_PROXY_RECURSION_DEPTH"`
	SteemitMaxRationDecayRate             *types.Int `json:"STEEMIT_MAX_RATION_DECAY_RATE"`
	SteemitMaxReserveRatio                *types.Int `json:"STEEMIT_MAX_RESERVE_RATIO"`
	SteemitMaxRunnerWitnesses             *types.Int `json:"STEEMIT_MAX_RUNNER_WITNESSES"`
	SteemitMaxShareSupply                 string     `json:"STEEMIT_MAX_SHARE_SUPPLY"`
	SteemitMaxSigCheckDepth               *types.Int `json:"STEEMIT_MAX_SIG_CHECK_DEPTH"`
	SteemitMaxTimeUntilExpiration         *types.Int `json:"STEEMIT_MAX_TIME_UNTIL_EXPIRATION"`
	SteemitMaxTransactionSize             *types.Int `json:"STEEMIT_MAX_TRANSACTION_SIZE"`
	SteemitMaxUndoHistory                 *types.Int `json:"STEEMIT_MAX_UNDO_HISTORY"`
	SteemitMaxUrlLength                   *types.Int `json:"STEEMIT_MAX_URL_LENGTH"`
	SteemitMaxVoteChanges                 *types.Int `json:"STEEMIT_MAX_VOTE_CHANGES"`
	SteemitMaxVotedWitnesses              *types.Int `json:"STEEMIT_MAX_VOTED_WITNESSES"`
	SteemitMaxWithdrawRoutes              *types.Int `json:"STEEMIT_MAX_WITHDRAW_ROUTES"`
	SteemitMaxWitnessUrlLength            *types.Int `json:"STEEMIT_MAX_WITNESS_URL_LENGTH"`
	SteemitMinAccountCreationFee          *types.Int `json:"STEEMIT_MIN_ACCOUNT_CREATION_FEE"`
	SteemitMinAccountNameLength           *types.Int `json:"STEEMIT_MIN_ACCOUNT_NAME_LENGTH"`
	SteemitMinBlockSizeLimit              *types.Int `json:"STEEMIT_MIN_BLOCK_SIZE_LIMIT"`
	SteemitMinContentReward               string     `json:"STEEMIT_MIN_CONTENT_REWARD"`
	SteemitMinCurateReward                string     `json:"STEEMIT_MIN_CURATE_REWARD"`
	SteemitMinerAccount                   string     `json:"STEEMIT_MINER_ACCOUNT"`
	SteemitMinerPayPercent                *types.Int `json:"STEEMIT_MINER_PAY_PERCENT"`
	SteemitMinFeeds                       *types.Int `json:"STEEMIT_MIN_FEEDS"`
	SteemitMiningReward                   string     `json:"STEEMIT_MINING_REWARD"`
	SteemitMiningTime                     string     `json:"STEEMIT_MINING_TIME"`
	steemitMinLiquidityReward             string     `json:"STEEMIT_MIN_LIQUIDITY_REWARD"`
	SteemitMinLiquidityRewardPeriodSec    *types.Int `json:"STEEMIT_MIN_LIQUIDITY_REWARD_PERIOD_SEC"`
	SteemitMinPayoutSbd                   string     `json:"STEEMIT_MIN_PAYOUT_SBD"`
	SteemitMinPowReward                   string     `json:"STEEMIT_MIN_POW_REWARD"`
	SteemitMinProducerReward              string     `json:"STEEMIT_MIN_PRODUCER_REWARD"`
	SteemitMinRation                      *types.Int `json:"STEEMIT_MIN_RATION"`
	SteemitMinTransactionExpirationLimit  *types.Int `json:"STEEMIT_MIN_TRANSACTION_EXPIRATION_LIMIT"`
	SteemitMinTransactionSizeLimit        *types.Int `json:"STEEMIT_MIN_TRANSACTION_SIZE_LIMIT"`
	SteemitMinUndoHistory                 *types.Int `json:"STEEMIT_MIN_UNDO_HISTORY"`
	SteemitNullAccount                    string     `json:"STEEMIT_NULL_ACCOUNT"`
	SteemitNumInitMiners                  *types.Int `json:"STEEMIT_NUM_INIT_MINERS"`
	SteemitPowAprPercent                  *types.Int `json:"STEEMIT_POW_APR_PERCENT"`
	SteemitProducerAprPercent             *types.Int `json:"STEEMIT_PRODUCER_APR_PERCENT"`
	SteemitProxyToSelfAccount             string     `json:"STEEMIT_PROXY_TO_SELF_ACCOUNT"`
	SteemitSbdInterestCompoundIntervalSec *types.Int `json:"STEEMIT_SBD_INTEREST_COMPOUND_INTERVAL_SEC"`
	SteemitSecondsPerYear                 *types.Int `json:"STEEMIT_SECONDS_PER_YEAR"`
	SteemitReverseAuctionWindowSeconds    *types.Int `json:"STEEMIT_REVERSE_AUCTION_WINDOW_SECONDS"`
	SteemitStartMinerVotingBlock          *types.Int `json:"STEEMIT_START_MINER_VOTING_BLOCK"`
	SteemitStartVestingBlock              *types.Int `json:"STEEMIT_START_VESTING_BLOCK"`
	SteemitSymbol                         string     `json:"STEEMIT_SYMBOL"`
	SteemitTempAccount                    string     `json:"STEEMIT_TEMP_ACCOUNT"`
	SteemitUpvoteLockout                  *types.Int `json:"STEEMIT_UPVOTE_LOCKOUT"`
	SteemitVestingWithdrawIntervals       *types.Int `json:"STEEMIT_VESTING_WITHDRAW_INTERVALS"`
	SteemitVestingWithdrawIntervalSeconds *types.Int `json:"STEEMIT_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	SteemitVoteChangeLockoutPeriod        *types.Int `json:"STEEMIT_VOTE_CHANGE_LOCKOUT_PERIOD"`
	SteemitVoteRegenerationSeconds        *types.Int `json:"STEEMIT_VOTE_REGENERATION_SECONDS"`
	SteemSymbol                           string     `json:"STEEM_SYMBOL"`
	VestsSymbol                           string     `json:"VESTS_SYMBOL"`
	BlockchainName                        string     `json:"BLOCKCHAIN_NAME"`
}

type DynamicGlobalProperties struct {
	Time                     *types.Time   `json:"time"`
	TotalPow                 *types.Int    `json:"total_pow"`
	NumPowWitnesses          *types.Int    `json:"num_pow_witnesses"`
	CurrentReserveRatio      *types.Int    `json:"current_reserve_ratio"`
	ID                       *types.ID     `json:"id"`
	CurrentSupply            string        `json:"current_supply"`
	CurrentSBDSupply         string        `json:"current_sbd_supply"`
	MaximumBlockSize         *types.Int    `json:"maximum_block_size"`
	RecentSlotsFilled        *types.Int    `json:"recent_slots_filled"`
	CurrentWitness           string        `json:"current_witness"`
	TotalRewardShares2       *types.Int    `json:"total_reward_shares2"`
	AverageBlockSize         *types.Int    `json:"average_block_size"`
	CurrentAslot             *types.Int    `json:"current_aslot"`
	LastIrreversibleBlockNum *types.UInt32 `json:"last_irreversible_block_num"`
	TotalVestingShares       string        `json:"total_vesting_shares"`
	TotalVersingFundSteem    string        `json:"total_vesting_fund_steem"`
	HeadBlockID              string        `json:"head_block_id"`
	HeadBlockNumber          *types.UInt32 `json:"head_block_number"`
	VirtualSupply            string        `json:"virtual_supply"`
	ConfidentialSupply       string        `json:"confidential_supply"`
	ConfidentialSBDSupply    string        `json:"confidential_sbd_supply"`
	TotalRewardFundSteem     string        `json:"total_reward_fund_steem"`
	TotalActivityFundSteem   string        `json:"total_activity_fund_steem"`
	TotalActivityFundShares  *types.Int    `json:"total_activity_fund_shares"`
	SBDInterestRate          *types.Int    `json:"sbd_interest_rate"`
	MaxVirtualBandwidth      *types.Int    `json:"max_virtual_bandwidth"`
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

type ChainProperties struct {
	AccountCreationFee string     `json:"account_creation_fee"`
	MaximumBlockSize   *types.Int `json:"maximum_block_size"`
	SbdInterestRate    *types.Int `json:"sbd_interest_rate"`
}

type NextScheduledHardfork struct {
	HfVersion string      `json:"hf_version"`
	LiveTime  *types.Time `json:"live_time"`
}

type CurrentMedianHistoryPrice struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
}

type ConversionRequests struct {
	ID             *types.Int  `json:"id"`
	Owner          string      `json:"owner"`
	Requestid      *types.Int  `json:"requestid"`
	Amount         string      `json:"amount"`
	ConversionDate *types.Time `json:"conversion_date"`
}

type Votes struct {
	Authorperm string      `json:"authorperm"`
	Weight     *types.Int  `json:"weight"`
	Rshares    *types.Int  `json:"rshares"`
	Percent    uint        `json:"percent"`
	Time       *types.Time `json:"time"`
}

type BlockHeader struct {
	Number                uint32        `json:"-"`
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	Witness               string        `json:"witness"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Extensions            []interface{} `json:"extensions"`
}

type OrderBook struct {
	Ask []*OrderBookAB `json:"asks"`
	Bid []*OrderBookAB `json:"bids"`
}

type OrderBookAB struct {
	OrderPrice *OrderPrice `json:"order_price"`
	RealPrice  string      `json:"real_price"`
	Steem      *types.Int  `json:"steem"`
	Sbd        *types.Int  `json:"sbd"`
	Created    string      `json:"created"`
}

type OrderPrice struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
}

type OpenOrders struct {
	ID         *types.ID   `json:"id"`
	Created    types.Time  `json:"created"`
	Expiration types.Time  `json:"expiration"`
	Seller     string      `json:"seller"`
	Orderid    *types.Int  `json:"orderid"`
	ForSale    *types.Int  `json:"for_sale"`
	SellPrice  *OrderPrice `json:"sell_price"`
	RealPrice  string      `json:"real_price"`
	Rewarded   bool        `json:"rewarded"`
}
