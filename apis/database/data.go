package database

import (
	// Stdlib
	"encoding/json"
	"strconv"
	"strings"

	// RPC
	"github.com/asuleymanov/golos-go/types"
)

type DiscussionQuery struct {
	Tag            string   `json:"tag"`
	Limit          uint32   `json:"limit"`
	FilterTags     []string `json:"filter_tags"`
	StartAuthor    string   `json:"start_author,omitempty"`
	StartPermlink  string   `json:"start_permlink,omitempty"`
	ParentAuthor   string   `json:"parent_author,omitempty"`
	ParentPermlink string   `json:"parent_permlink"`
}

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
	SteemitBlockInterval                  uint       `json:"STEEMIT_BLOCK_INTERVAL"`
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
	Time                     *types.Time `json:"time"`
	TotalPow                 *types.Int  `json:"total_pow"`
	NumPowWitnesses          *types.Int  `json:"num_pow_witnesses"`
	CurrentReserveRatio      *types.Int  `json:"current_reserve_ratio"`
	ID                       *types.ID   `json:"id"`
	CurrentSupply            string      `json:"current_supply"`
	CurrentSBDSupply         string      `json:"current_sbd_supply"`
	MaximumBlockSize         *types.Int  `json:"maximum_block_size"`
	RecentSlotsFilled        *types.Int  `json:"recent_slots_filled"`
	CurrentWitness           string      `json:"current_witness"`
	TotalRewardShares2       *types.Int  `json:"total_reward_shares2"`
	AverageBlockSize         *types.Int  `json:"average_block_size"`
	CurrentAslot             *types.Int  `json:"current_aslot"`
	LastIrreversibleBlockNum uint32      `json:"last_irreversible_block_num"`
	TotalVestingShares       string      `json:"total_vesting_shares"`
	TotalVersingFundSteem    string      `json:"total_vesting_fund_steem"`
	HeadBlockID              string      `json:"head_block_id"`
	HeadBlockNumber          uint32      `json:"head_block_number"`
	VirtualSupply            string      `json:"virtual_supply"`
	ConfidentialSupply       string      `json:"confidential_supply"`
	ConfidentialSBDSupply    string      `json:"confidential_sbd_supply"`
	TotalRewardFundSteem     string      `json:"total_reward_fund_steem"`
	TotalActivityFundSteem   string      `json:"total_activity_fund_steem"`
	TotalActivityFundShares  *types.Int  `json:"total_activity_fund_shares"`
	SBDInterestRate          *types.Int  `json:"sbd_interest_rate"`
	MaxVirtualBandwidth      *types.Int  `json:"max_virtual_bandwidth"`
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
	ID                      *types.ID        `json:"id"`
	Author                  string           `json:"author"`
	Permlink                string           `json:"permlink"`
	Category                string           `json:"category"`
	ParentAuthor            string           `json:"parent_author"`
	ParentPermlink          string           `json:"parent_permlink"`
	Title                   string           `json:"title"`
	Body                    string           `json:"body"`
	JsonMetadata            *ContentMetadata `json:"json_metadata"`
	LastUpdate              *types.Time      `json:"last_update"`
	Created                 *types.Time      `json:"created"`
	Active                  *types.Time      `json:"active"`
	LastPayout              *types.Time      `json:"last_payout"`
	Depth                   *types.Int       `json:"depth"`
	Children                *types.Int       `json:"children"`
	ChildrenRshares2        *types.Int       `json:"children_rshares2"`
	NetRshares              *types.Int       `json:"net_rshares"`
	AbsRshares              *types.Int       `json:"abs_rshares"`
	VoteRshares             *types.Int       `json:"vote_rshares"`
	ChildrenAbsRshares      *types.Int       `json:"children_abs_rshares"`
	CashoutTime             *types.Time      `json:"cashout_time"`
	MaxCashoutTime          *types.Time      `json:"max_cashout_time"`
	TotalVoteWeight         *types.Int       `json:"total_vote_weight"`
	RewardWeight            *types.Int       `json:"reward_weight"`
	TotalPayoutValue        string           `json:"total_payout_value"`
	CuratorPayoutValue      string           `json:"curator_payout_value"`
	AuthorRewards           *types.Int       `json:"author_rewards"`
	NetVotes                *types.Int       `json:"net_votes"`
	RootComment             *types.Int       `json:"root_comment"`
	Mode                    string           `json:"mode"`
	MaxAcceptedPayout       string           `json:"max_accepted_payout"`
	PercentSteemDollars     *types.Int       `json:"percent_steem_dollars"`
	AllowReplies            bool             `json:"allow_replies"`
	AllowVotes              bool             `json:"allow_votes"`
	AllowCurationRewards    bool             `json:"allow_curation_rewards"`
	URL                     string           `json:"url"`
	RootTitle               string           `json:"root_title"`
	PendingPayoutValue      string           `json:"pending_payout_value"`
	TotalPendingPayoutValue string           `json:"total_pending_payout_value"`
	ActiveVotes             []*VoteState     `json:"active_votes"`
	Replies                 []*Content       `json:"replies"`
	AuthorReputation        *types.Int       `json:"author_reputation"`
	Promoted                string           `json:"promoted"`
	BodyLength              *types.Int       `json:"body_length"`
	RebloggedBy             []interface{}    `json:"reblogged_by"`
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

type AccountKeys struct {
	WeightThreshold *types.Int    `json:"weight_threshold"`
	AccountAuths    []interface{} `json:"account_auths"`
	KeyAuths        []interface{} `json:"key_auths"`
}

type Account struct {
	ID                            *types.Int    `json:"id"`
	Name                          string        `json:"name"`
	Owner                         *AccountKeys  `json:"owner"`
	Active                        *AccountKeys  `json:"active"`
	Posting                       *AccountKeys  `json:"posting"`
	MemoKey                       string        `json:"memo_key"`
	JSONMetadata                  string        `json:"json_metadata"`
	Proxy                         string        `json:"proxy"`
	LastOwnerUpdate               *types.Time   `json:"last_owner_update"`
	LastAccountUpdate             *types.Time   `json:"last_account_update"`
	Created                       *types.Time   `json:"created"`
	Mined                         bool          `json:"mined"`
	OwnerChallenged               bool          `json:"owner_challenged"`
	ActiveChallenged              bool          `json:"active_challenged"`
	LastOwnerProved               *types.Time   `json:"last_owner_proved"`
	LastActiveProved              *types.Time   `json:"last_active_proved"`
	RecoveryAccount               string        `json:"recovery_account"`
	LastAccountRecovery           *types.Time   `json:"last_account_recovery"`
	ResetAccount                  string        `json:"reset_account"`
	CommentCount                  *types.Int    `json:"comment_count"`
	LifetimeVoteCount             *types.Int    `json:"lifetime_vote_count"`
	PostCount                     *types.Int    `json:"post_count"`
	CanVote                       bool          `json:"can_vote"`
	VotingPower                   *types.Int    `json:"voting_power"`
	LastVoteTime                  *types.Time   `json:"last_vote_time"`
	Balance                       string        `json:"balance"`
	SavingsBalance                string        `json:"savings_balance"`
	SbdBalance                    string        `json:"sbd_balance"`
	SbdSeconds                    string        `json:"sbd_seconds"`
	SbdSecondsLastUpdate          *types.Time   `json:"sbd_seconds_last_update"`
	SbdLastInterestPayment        *types.Time   `json:"sbd_last_interest_payment"`
	SavingsSbdBalance             string        `json:"savings_sbd_balance"`
	SavingsSbdSeconds             string        `json:"savings_sbd_seconds"`
	SavingsSbdSecondsLastUpdate   *types.Time   `json:"savings_sbd_seconds_last_update"`
	SavingsSbdLastInterestPayment *types.Time   `json:"savings_sbd_last_interest_payment"`
	SavingsWithdrawRequests       *types.Int    `json:"savings_withdraw_requests"`
	VestingShares                 string        `json:"vesting_shares"`
	VestingWithdrawRate           string        `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal         *types.Time   `json:"next_vesting_withdrawal"`
	Withdrawn                     *types.Int    `json:"withdrawn"`
	ToWithdraw                    *types.Int    `json:"to_withdraw"`
	WithdrawRoutes                *types.Int    `json:"withdraw_routes"`
	CurationRewards               *types.Int    `json:"curation_rewards"`
	PostingRewards                *types.Int    `json:"posting_rewards"`
	ProxiedVsfVotes               []*types.Int  `json:"proxied_vsf_votes"`
	WitnessesVotedFor             *types.Int    `json:"witnesses_voted_for"`
	AverageBandwidth              *types.Int    `json:"average_bandwidth"`
	LifetimeBandwidth             string        `json:"lifetime_bandwidth"`
	LastBandwidthUpdate           *types.Time   `json:"last_bandwidth_update"`
	AverageMarketBandwidth        *types.Int    `json:"average_market_bandwidth"`
	LastMarketBandwidthUpdate     *types.Time   `json:"last_market_bandwidth_update"`
	LastPost                      *types.Time   `json:"last_post"`
	LastRootPost                  *types.Time   `json:"last_root_post"`
	PostBandwidth                 *types.Int    `json:"post_bandwidth"`
	NewAverageBandwidth           string        `json:"new_average_bandwidth"`
	NewAverageMarketBandwidth     *types.Int64  `json:"new_average_market_bandwidth"`
	VestingBalance                string        `json:"vesting_balance"`
	Reputation                    string        `json:"reputation"`
	TransferHistory               []interface{} `json:"transfer_history"`
	MarketHistory                 []interface{} `json:"market_history"`
	PostHistory                   []interface{} `json:"post_history"`
	VoteHistory                   []interface{} `json:"vote_history"`
	OtherHistory                  []interface{} `json:"other_history"`
	WitnessVotes                  []string      `json:"witness_votes"`
	TagsUsage                     []interface{} `json:"tags_usage"`
	GuestBloggers                 []interface{} `json:"guest_bloggers"`
	BlogCategory                  interface{}   `json:"blog_category"`
}

type WitnessSchedule struct {
	ID                            *types.Int       `json:"id"`
	CurrentVirtualTime            string           `json:"current_virtual_time"`
	NextShuffleBlockNum           *types.Int       `json:"next_shuffle_block_num"`
	CurrentShuffledWitnesses      string           `json:"current_shuffled_witnesses"`
	NumScheduledWitnesses         *types.Int       `json:"num_scheduled_witnesses"`
	Top19Weight                   *types.Int       `json:"top19_weight"`
	TimeshareWeight               *types.Int       `json:"timeshare_weight"`
	MinerWeight                   *types.Int       `json:"miner_weight"`
	WitnessPayNormalizationFactor *types.Int       `json:"witness_pay_normalization_factor"`
	MedianProps                   *ChainProperties `json:"median_props"`
	MajorityVersion               string           `json:"majority_version"`
}

type FeedHistory struct {
	ID                   *types.Int                   `json:"id"`
	CurrentMedianHistory *CurrentMedianHistoryPrice   `json:"current_median_history"`
	PriceHistory         []*CurrentMedianHistoryPrice `json:"price_history"`
}

type Witness struct {
	ID                    *types.Int                 `json:"id"`
	Owner                 string                     `json:"owner"`
	Created               *types.Time                `json:"created"`
	URL                   string                     `json:"url"`
	Votes                 string                     `json:"votes"`
	VirtualLastUpdate     string                     `json:"virtual_last_update"`
	VirtualPosition       string                     `json:"virtual_position"`
	VirtualScheduledTime  string                     `json:"virtual_scheduled_time"`
	TotalMissed           *types.Int                 `json:"total_missed"`
	LastAslot             *types.Int                 `json:"last_aslot"`
	LastConfirmedBlockNum *types.Int                 `json:"last_confirmed_block_num"`
	PowWorker             *types.Int                 `json:"pow_worker"`
	SigningKey            string                     `json:"signing_key"`
	Props                 *ChainProperties           `json:"props"`
	SbdExchangeRate       *CurrentMedianHistoryPrice `json:"sbd_exchange_rate"`
	LastSbdExchangeUpdate *types.Time                `json:"last_sbd_exchange_update"`
	LastWork              string                     `json:"last_work"`
	RunningVersion        string                     `json:"running_version"`
	HardforkVersionVote   string                     `json:"hardfork_version_vote"`
	HardforkTimeVote      *types.Time                `json:"hardfork_time_vote"`
}

type SavingsWithdraw struct {
	ID        *types.ID   `json:"id"`
	From      string      `json:"from"`
	To        string      `json:"to"`
	Memo      string      `json:"memo"`
	RequestID *types.Int  `json:"request_id"`
	Amount    string      `json:"amount"`
	Complete  *types.Time `json:"complete"`
}

type TrendingTags struct {
	Name                  string     `json:"name"`
	TotalChildrenRshares2 string     `json:"total_children_rshares2"`
	TotalPayouts          string     `json:"total_payouts"`
	NetVotes              *types.Int `json:"net_votes"`
	TopPosts              *types.Int `json:"top_posts"`
	Comments              *types.Int `json:"comments"`
}

type Categories struct {
	ID           *types.Int `json:"id"`
	Name         string     `json:"name"`
	AbsRshares   string     `json:"abs_rshares"`
	TotalPayouts string     `json:"total_payouts"`
	Discussions  *types.Int `json:"discussions"`
	LastUpdate   string     `json:"last_update"`
}
