package database

import (
	"github.com/asuleymanov/steem-go/types"
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
	BuildTestnet                   bool       `json:"STEEMIT_BUILD_TESTNET"`
	GrapheneCurrentDBVersion       string     `json:"GRAPHENE_CURRENT_DB_VERSION"`
	SbdSymbol                      *types.Int `json:"SBD_SYMBOL"`
	Percent100                     int        `json:"STEEMIT_100_PERCENT"`
	Percent1                       *types.Int `json:"STEEMIT_1_PERCENT"`
	AddressPrefix                  string     `json:"STEEMIT_ADDRESS_PREFIX"`
	AprPercentMultiplyPerBlock     string     `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_BLOCK"`
	AprPercentMultiplyPerHour      string     `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_HOUR"`
	AprPercentMultiplyPerRound     string     `json:"STEEMIT_APR_PERCENT_MULTIPLY_PER_ROUND"`
	AprPercentShiftPerBlock        *types.Int `json:"STEEMIT_APR_PERCENT_SHIFT_PER_BLOCK"`
	AprPercentShiftPerHour         *types.Int `json:"STEEMIT_APR_PERCENT_SHIFT_PER_HOUR"`
	AprPercentShiftPerRound        *types.Int `json:"STEEMIT_APR_PERCENT_SHIFT_PER_ROUND"`
	BandwidthAverageWindowSeconds  *types.Int `json:"STEEMIT_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	BandwidthPrecision             *types.Int `json:"STEEMIT_BANDWIDTH_PRECISION"`
	BlockchainPrecision            *types.Int `json:"STEEMIT_BLOCKCHAIN_PRECISION"`
	BlockchainPrecisionDigits      *types.Int `json:"STEEMIT_BLOCKCHAIN_PRECISION_DIGITS"`
	BlockchainHardforkVersion      string     `json:"STEEMIT_BLOCKCHAIN_HARDFORK_VERSION"`
	BlockchainVersion              string     `json:"STEEMIT_BLOCKCHAIN_VERSION"`
	BlockInterval                  uint       `json:"STEEMIT_BLOCK_INTERVAL"`
	BlocksPerDay                   *types.Int `json:"STEEMIT_BLOCKS_PER_DAY"`
	BlocksPerHour                  *types.Int `json:"STEEMIT_BLOCKS_PER_HOUR"`
	BlocksPerYear                  *types.Int `json:"STEEMIT_BLOCKS_PER_YEAR"`
	CashoutWindowSeconds           *types.Int `json:"STEEMIT_CASHOUT_WINDOW_SECONDS"`
	ChainID                        string     `json:"STEEMIT_CHAIN_ID"`
	ContentAprPercent              *types.Int `json:"STEEMIT_CONTENT_APR_PERCENT"`
	ConversionDelay                string     `json:"STEEMIT_CONVERSION_DELAY"`
	CurateAprPercent               *types.Int `json:"STEEMIT_CURATE_APR_PERCENT"`
	DefaultSbdInterestRate         *types.Int `json:"STEEMIT_DEFAULT_SBD_INTEREST_RATE"`
	FeedHistoryWindow              *types.Int `json:"STEEMIT_FEED_HISTORY_WINDOW"`
	FeedIntervalBlocks             *types.Int `json:"STEEMIT_FEED_INTERVAL_BLOCKS"`
	FreeTransactionsWithNewAccount *types.Int `json:"STEEMIT_FREE_TRANSACTIONS_WITH_NEW_ACCOUNT"`
	GenesisTime                    string     `json:"STEEMIT_GENESIS_TIME"`
	HardforkRequiredWitnesses      *types.Int `json:"STEEMIT_HARDFORK_REQUIRED_WITNESSES"`
	InitMinerName                  string     `json:"STEEMIT_INIT_MINER_NAME"`
	InitPublicKeyStr               string     `json:"STEEMIT_INIT_PUBLIC_KEY_STR"`
	InitSupply                     *types.Int `json:"STEEMIT_INIT_SUPPLY"`
	InitTime                       string     `json:"STEEMIT_INIT_TIME"`
	IrreversibleThreshold          *types.Int `json:"STEEMIT_IRREVERSIBLE_THRESHOLD"`
	LiquidityAprPercent            *types.Int `json:"STEEMIT_LIQUIDITY_APR_PERCENT"`
	LiquidityRewardBlocks          *types.Int `json:"STEEMIT_LIQUIDITY_REWARD_BLOCKS"`
	LiquidityRewardPeriodSec       *types.Int `json:"STEEMIT_LIQUIDITY_REWARD_PERIOD_SEC"`
	LiquidityTimeoutSec            string     `json:"STEEMIT_LIQUIDITY_TIMEOUT_SEC"`
	MaxAccountNameLength           *types.Int `json:"STEEMIT_MAX_ACCOUNT_NAME_LENGTH"`
	MaxAccountWitnessVotes         *types.Int `json:"STEEMIT_MAX_ACCOUNT_WITNESS_VOTES"`
	MaxAssetWhitelistAuthorities   *types.Int `json:"STEEMIT_MAX_ASSET_WHITELIST_AUTHORITIES"`
	MaxAuthorityMembership         *types.Int `json:"STEEMIT_MAX_AUTHORITY_MEMBERSHIP"`
	MaxBlockSize                   *types.Int `json:"STEEMIT_MAX_BLOCK_SIZE"`
	MaxCashoutWindowSeconds        *types.Int `json:"STEEMIT_MAX_CASHOUT_WINDOW_SECONDS"`
	MaxCommentDepth                *types.Int `json:"STEEMIT_MAX_COMMENT_DEPTH"`
	MaxFeedAge                     string     `json:"STEEMIT_MAX_FEED_AGE"`
	MaxInstanceId                  string     `json:"STEEMIT_MAX_INSTANCE_ID"`
	MaxMemoSize                    *types.Int `json:"STEEMIT_MAX_MEMO_SIZE"`
	MaxWitnesses                   *types.Int `json:"STEEMIT_MAX_WITNESSES"`
	MaxMinerWitnesses              *types.Int `json:"STEEMIT_MAX_MINER_WITNESSES"`
	MaxProxyRecursionDepth         *types.Int `json:"STEEMIT_MAX_PROXY_RECURSION_DEPTH"`
	MaxRationDecayRate             *types.Int `json:"STEEMIT_MAX_RATION_DECAY_RATE"`
	MaxReserveRatio                *types.Int `json:"STEEMIT_MAX_RESERVE_RATIO"`
	MaxRunnerWitnesses             *types.Int `json:"STEEMIT_MAX_RUNNER_WITNESSES"`
	MaxShareSupply                 string     `json:"STEEMIT_MAX_SHARE_SUPPLY"`
	MaxSigCheckDepth               *types.Int `json:"STEEMIT_MAX_SIG_CHECK_DEPTH"`
	MaxTimeUntilExpiration         *types.Int `json:"STEEMIT_MAX_TIME_UNTIL_EXPIRATION"`
	MaxTransactionSize             *types.Int `json:"STEEMIT_MAX_TRANSACTION_SIZE"`
	MaxUndoHistory                 *types.Int `json:"STEEMIT_MAX_UNDO_HISTORY"`
	MaxUrlLength                   *types.Int `json:"STEEMIT_MAX_URL_LENGTH"`
	MaxVoteChanges                 *types.Int `json:"STEEMIT_MAX_VOTE_CHANGES"`
	MaxVotedWitnesses              *types.Int `json:"STEEMIT_MAX_VOTED_WITNESSES"`
	MaxWithdrawRoutes              *types.Int `json:"STEEMIT_MAX_WITHDRAW_ROUTES"`
	MaxWitnessUrlLength            *types.Int `json:"STEEMIT_MAX_WITNESS_URL_LENGTH"`
	MinAccountCreationFee          *types.Int `json:"STEEMIT_MIN_ACCOUNT_CREATION_FEE"`
	MinAccountNameLength           *types.Int `json:"STEEMIT_MIN_ACCOUNT_NAME_LENGTH"`
	MinBlockSizeLimit              *types.Int `json:"STEEMIT_MIN_BLOCK_SIZE_LIMIT"`
	MinContentReward               string     `json:"STEEMIT_MIN_CONTENT_REWARD"`
	MinCurateReward                string     `json:"STEEMIT_MIN_CURATE_REWARD"`
	MinerAccount                   string     `json:"STEEMIT_MINER_ACCOUNT"`
	MinerPayPercent                *types.Int `json:"STEEMIT_MINER_PAY_PERCENT"`
	MinFeeds                       *types.Int `json:"STEEMIT_MIN_FEEDS"`
	MiningReward                   string     `json:"STEEMIT_MINING_REWARD"`
	MiningTime                     string     `json:"STEEMIT_MINING_TIME"`
	MinLiquidityReward             string     `json:"STEEMIT_MIN_LIQUIDITY_REWARD"`
	MinLiquidityRewardPeriodSec    *types.Int `json:"STEEMIT_MIN_LIQUIDITY_REWARD_PERIOD_SEC"`
	MinPayoutSbd                   string     `json:"STEEMIT_MIN_PAYOUT_SBD"`
	MinPowReward                   string     `json:"STEEMIT_MIN_POW_REWARD"`
	MinProducerReward              string     `json:"STEEMIT_MIN_PRODUCER_REWARD"`
	MinRation                      *types.Int `json:"STEEMIT_MIN_RATION"`
	MinTransactionExpirationLimit  *types.Int `json:"STEEMIT_MIN_TRANSACTION_EXPIRATION_LIMIT"`
	MinTransactionSizeLimit        *types.Int `json:"STEEMIT_MIN_TRANSACTION_SIZE_LIMIT"`
	MinUndoHistory                 *types.Int `json:"STEEMIT_MIN_UNDO_HISTORY"`
	NullAccount                    string     `json:"STEEMIT_NULL_ACCOUNT"`
	NumInitMiners                  *types.Int `json:"STEEMIT_NUM_INIT_MINERS"`
	PowAprPercent                  *types.Int `json:"STEEMIT_POW_APR_PERCENT"`
	ProducerAprPercent             *types.Int `json:"STEEMIT_PRODUCER_APR_PERCENT"`
	ProxyToSelfAccount             string     `json:"STEEMIT_PROXY_TO_SELF_ACCOUNT"`
	SbdInterestCompoundIntervalSec *types.Int `json:"STEEMIT_SBD_INTEREST_COMPOUND_INTERVAL_SEC"`
	SecondsPerYear                 *types.Int `json:"STEEMIT_SECONDS_PER_YEAR"`
	ReverseAuctionWindowSeconds    *types.Int `json:"STEEMIT_REVERSE_AUCTION_WINDOW_SECONDS"`
	StartMinerVotingBlock          *types.Int `json:"STEEMIT_START_MINER_VOTING_BLOCK"`
	StartVestingBlock              *types.Int `json:"STEEMIT_START_VESTING_BLOCK"`
	Symbol                         string     `json:"STEEMIT_SYMBOL"`
	TempAccount                    string     `json:"STEEMIT_TEMP_ACCOUNT"`
	UpvoteLockout                  *types.Int `json:"STEEMIT_UPVOTE_LOCKOUT"`
	VestingWithdrawIntervals       *types.Int `json:"STEEMIT_VESTING_WITHDRAW_INTERVALS"`
	VestingWithdrawIntervalSeconds *types.Int `json:"STEEMIT_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	VoteChangeLockoutPeriod        *types.Int `json:"STEEMIT_VOTE_CHANGE_LOCKOUT_PERIOD"`
	VoteRegenerationSeconds        int        `json:"STEEMIT_VOTE_REGENERATION_SECONDS"`
	SteemSymbol                    string     `json:"STEEM_SYMBOL"`
	VestsSymbol                    string     `json:"VESTS_SYMBOL"`
	BlockchainName                 string     `json:"BLOCKCHAIN_NAME"`
}

type DynamicGlobalProperties struct {
	Time                     *types.Time  `json:"time"`
	TotalPow                 *types.Int   `json:"total_pow"`
	NumPowWitnesses          *types.Int   `json:"num_pow_witnesses"`
	CurrentReserveRatio      *types.Int   `json:"current_reserve_ratio"`
	ID                       *types.ID    `json:"id"`
	CurrentSupply            *types.Asset `json:"current_supply"`
	CurrentSBDSupply         *types.Asset `json:"current_sbd_supply"`
	MaximumBlockSize         *types.Int   `json:"maximum_block_size"`
	RecentSlotsFilled        *types.Int   `json:"recent_slots_filled"`
	CurrentWitness           string       `json:"current_witness"`
	TotalRewardShares2       *types.Int   `json:"total_reward_shares2"`
	AverageBlockSize         *types.Int   `json:"average_block_size"`
	CurrentAslot             *types.Int   `json:"current_aslot"`
	LastIrreversibleBlockNum uint32       `json:"last_irreversible_block_num"`
	TotalVestingShares       *types.Asset `json:"total_vesting_shares"`
	TotalVersingFundSteem    *types.Asset `json:"total_vesting_fund_steem"`
	HeadBlockID              string       `json:"head_block_id"`
	HeadBlockNumber          uint32       `json:"head_block_number"`
	VirtualSupply            *types.Asset `json:"virtual_supply"`
	ConfidentialSupply       *types.Asset `json:"confidential_supply"`
	ConfidentialSBDSupply    *types.Asset `json:"confidential_sbd_supply"`
	TotalRewardFundSteem     *types.Asset `json:"total_reward_fund_steem"`
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
	ID                      *types.ID              `json:"id"`
	Author                  string                 `json:"author"`
	Permlink                string                 `json:"permlink"`
	Category                string                 `json:"category"`
	ParentAuthor            string                 `json:"parent_author"`
	ParentPermlink          string                 `json:"parent_permlink"`
	Title                   string                 `json:"title"`
	Body                    string                 `json:"body"`
	JsonMetadata            *types.ContentMetadata `json:"json_metadata"`
	LastUpdate              *types.Time            `json:"last_update"`
	Created                 *types.Time            `json:"created"`
	Active                  *types.Time            `json:"active"`
	LastPayout              *types.Time            `json:"last_payout"`
	Depth                   *types.Int             `json:"depth"`
	Children                *types.Int             `json:"children"`
	ChildrenRshares2        *types.Int             `json:"children_rshares2"`
	NetRshares              *types.Int             `json:"net_rshares"`
	AbsRshares              *types.Int             `json:"abs_rshares"`
	VoteRshares             *types.Int             `json:"vote_rshares"`
	ChildrenAbsRshares      *types.Int             `json:"children_abs_rshares"`
	CashoutTime             *types.Time            `json:"cashout_time"`
	MaxCashoutTime          *types.Time            `json:"max_cashout_time"`
	TotalVoteWeight         *types.Int             `json:"total_vote_weight"`
	RewardWeight            *types.Int             `json:"reward_weight"`
	TotalPayoutValue        string                 `json:"total_payout_value"`
	CuratorPayoutValue      string                 `json:"curator_payout_value"`
	AuthorRewards           *types.Int             `json:"author_rewards"`
	NetVotes                int                    `json:"net_votes"`
	RootComment             *types.Int             `json:"root_comment"`
	Mode                    string                 `json:"mode"`
	MaxAcceptedPayout       string                 `json:"max_accepted_payout"`
	PercentSteemDollars     *types.Int             `json:"percent_steem_dollars"`
	AllowReplies            bool                   `json:"allow_replies"`
	AllowVotes              bool                   `json:"allow_votes"`
	AllowCurationRewards    bool                   `json:"allow_curation_rewards"`
	URL                     string                 `json:"url"`
	RootTitle               string                 `json:"root_title"`
	PendingPayoutValue      string                 `json:"pending_payout_value"`
	TotalPendingPayoutValue string                 `json:"total_pending_payout_value"`
	ActiveVotes             []*VoteState           `json:"active_votes"`
	Replies                 []*Content             `json:"replies"`
	AuthorReputation        *types.Int             `json:"author_reputation"`
	Promoted                string                 `json:"promoted"`
	BodyLength              *types.Int             `json:"body_length"`
	RebloggedBy             []interface{}          `json:"reblogged_by"`
}

func (content *Content) IsStory() bool {
	return content.ParentAuthor == ""
}

type VoteState struct {
	Voter   string       `json:"voter"`
	Weight  *types.Int   `json:"weight"`
	Rshares *types.Int64 `json:"rshares"`
	Percent int          `json:"percent"`
	Time    *types.Time  `json:"time"`
}

type NextScheduledHardfork struct {
	HfVersion string      `json:"hf_version"`
	LiveTime  *types.Time `json:"live_time"`
}

type CurrentMedianHistoryPrice struct {
	Base  *types.Asset `json:"base"`
	Quote *types.Asset `json:"quote"`
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

type Account struct {
	ID                            *types.Int             `json:"id"`
	Name                          string                 `json:"name"`
	Owner                         *types.Authority       `json:"owner"`
	Active                        *types.Authority       `json:"active"`
	Posting                       *types.Authority       `json:"posting"`
	MemoKey                       string                 `json:"memo_key"`
	JSONMetadata                  *types.AccountMetadata `json:"json_metadata"`
	Proxy                         string                 `json:"proxy"`
	LastOwnerUpdate               *types.Time            `json:"last_owner_update"`
	LastAccountUpdate             *types.Time            `json:"last_account_update"`
	Created                       *types.Time            `json:"created"`
	Mined                         bool                   `json:"mined"`
	OwnerChallenged               bool                   `json:"owner_challenged"`
	ActiveChallenged              bool                   `json:"active_challenged"`
	LastOwnerProved               *types.Time            `json:"last_owner_proved"`
	LastActiveProved              *types.Time            `json:"last_active_proved"`
	RecoveryAccount               string                 `json:"recovery_account"`
	LastAccountRecovery           *types.Time            `json:"last_account_recovery"`
	ResetAccount                  string                 `json:"reset_account"`
	CommentCount                  *types.Int             `json:"comment_count"`
	LifetimeVoteCount             *types.Int             `json:"lifetime_vote_count"`
	PostCount                     *types.Int             `json:"post_count"`
	CanVote                       bool                   `json:"can_vote"`
	VotingPower                   int                    `json:"voting_power"`
	LastVoteTime                  *types.Time            `json:"last_vote_time"`
	Balance                       *types.Asset           `json:"balance"`
	SavingsBalance                *types.Asset           `json:"savings_balance"`
	SbdBalance                    *types.Asset           `json:"sbd_balance"`
	SbdSeconds                    string                 `json:"sbd_seconds"`
	SbdSecondsLastUpdate          *types.Time            `json:"sbd_seconds_last_update"`
	SbdLastInterestPayment        *types.Time            `json:"sbd_last_interest_payment"`
	SavingsSbdBalance             *types.Asset           `json:"savings_sbd_balance"`
	SavingsSbdSeconds             string                 `json:"savings_sbd_seconds"`
	SavingsSbdSecondsLastUpdate   *types.Time            `json:"savings_sbd_seconds_last_update"`
	SavingsSbdLastInterestPayment *types.Time            `json:"savings_sbd_last_interest_payment"`
	SavingsWithdrawRequests       *types.Int             `json:"savings_withdraw_requests"`
	VestingShares                 *types.Asset           `json:"vesting_shares"`
	VestingWithdrawRate           *types.Asset           `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal         *types.Time            `json:"next_vesting_withdrawal"`
	Withdrawn                     *types.Int             `json:"withdrawn"`
	ToWithdraw                    *types.Int             `json:"to_withdraw"`
	WithdrawRoutes                *types.Int             `json:"withdraw_routes"`
	CurationRewards               *types.Int             `json:"curation_rewards"`
	PostingRewards                *types.Int             `json:"posting_rewards"`
	ProxiedVsfVotes               []*types.Int           `json:"proxied_vsf_votes"`
	WitnessesVotedFor             *types.Int             `json:"witnesses_voted_for"`
	AverageBandwidth              *types.Int             `json:"average_bandwidth"`
	LifetimeBandwidth             *types.Int             `json:"lifetime_bandwidth"`
	LastBandwidthUpdate           *types.Time            `json:"last_bandwidth_update"`
	AverageMarketBandwidth        *types.Int             `json:"average_market_bandwidth"`
	LastMarketBandwidthUpdate     *types.Time            `json:"last_market_bandwidth_update"`
	LastPost                      *types.Time            `json:"last_post"`
	LastRootPost                  *types.Time            `json:"last_root_post"`
	PostBandwidth                 int                    `json:"post_bandwidth"`
	NewAverageBandwidth           string                 `json:"new_average_bandwidth"`
	NewAverageMarketBandwidth     *types.Int64           `json:"new_average_market_bandwidth"`
	VestingBalance                string                 `json:"vesting_balance"`
	Reputation                    *types.Int64           `json:"reputation"`
	TransferHistory               []interface{}          `json:"transfer_history"`
	MarketHistory                 []interface{}          `json:"market_history"`
	PostHistory                   []interface{}          `json:"post_history"`
	VoteHistory                   []interface{}          `json:"vote_history"`
	OtherHistory                  []interface{}          `json:"other_history"`
	WitnessVotes                  []string               `json:"witness_votes"`
	TagsUsage                     []interface{}          `json:"tags_usage"`
	GuestBloggers                 []interface{}          `json:"guest_bloggers"`
	BlogCategory                  interface{}            `json:"blog_category"`
}

type WitnessSchedule struct {
	ID                            *types.Int             `json:"id"`
	CurrentVirtualTime            string                 `json:"current_virtual_time"`
	NextShuffleBlockNum           *types.Int             `json:"next_shuffle_block_num"`
	CurrentShuffledWitnesses      string                 `json:"current_shuffled_witnesses"`
	NumScheduledWitnesses         *types.Int             `json:"num_scheduled_witnesses"`
	Top19Weight                   *types.Int             `json:"top19_weight"`
	TimeshareWeight               *types.Int             `json:"timeshare_weight"`
	MinerWeight                   *types.Int             `json:"miner_weight"`
	WitnessPayNormalizationFactor *types.Int             `json:"witness_pay_normalization_factor"`
	MedianProps                   *types.ChainProperties `json:"median_props"`
	MajorityVersion               string                 `json:"majority_version"`
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
	Votes                 *types.Int                 `json:"votes"`
	VirtualLastUpdate     *types.Int                 `json:"virtual_last_update"`
	VirtualPosition       *types.Int                 `json:"virtual_position"`
	VirtualScheduledTime  *types.Int                 `json:"virtual_scheduled_time"`
	TotalMissed           *types.Int                 `json:"total_missed"`
	LastAslot             *types.Int                 `json:"last_aslot"`
	LastConfirmedBlockNum *types.Int                 `json:"last_confirmed_block_num"`
	PowWorker             *types.Int                 `json:"pow_worker"`
	SigningKey            string                     `json:"signing_key"`
	Props                 *types.ChainProperties     `json:"props"`
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
