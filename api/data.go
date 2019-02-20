package api

import (
	"github.com/asuleymanov/steem-go/types"
)

//BroadcastResponse structure for the BroadcastTransactionSynchronous function
type BroadcastResponse struct {
	ID       string `json:"id"`
	BlockNum int32  `json:"block_num"`
	TrxNum   int32  `json:"trx_num"`
	Expired  bool   `json:"expired"`
}

//Config structure for the GetConfig function.
type Config struct {
	Percent100                       int           `json:"CHAIN_100_PERCENT"`
	Percent1                         *types.Int    `json:"CHAIN_1_PERCENT"`
	AddressPrefix                    string        `json:"CHAIN_ADDRESS_PREFIX"`
	BandwidthAverageWindowSeconds    *types.Int    `json:"CHAIN_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	BandwidthPrecision               *types.Int    `json:"CHAIN_BANDWIDTH_PRECISION"`
	ConsensusBandwidthReservePercent *types.Int    `json:"CONSENSUS_BANDWIDTH_RESERVE_PERCENT"`
	ConsensusBandwidthReserveBelow   *types.Int    `json:"CONSENSUS_BANDWIDTH_RESERVE_BELOW"`
	HardforkVersion                  string        `json:"CHAIN_HARDFORK_VERSION"`
	Version                          string        `json:"CHAIN_VERSION"`
	BlockInterval                    uint          `json:"CHAIN_BLOCK_INTERVAL"`
	BlocksPerDay                     *types.Int    `json:"CHAIN_BLOCKS_PER_DAY"`
	BlocksPerYear                    *types.Int    `json:"CHAIN_BLOCKS_PER_YEAR"`
	CashoutWindowSeconds             *types.Int    `json:"CHAIN_CASHOUT_WINDOW_SECONDS"`
	ChainID                          string        `json:"CHAIN_ID"`
	HardforkRequiredWitnesses        *types.Int    `json:"CHAIN_HARDFORK_REQUIRED_WITNESSES"`
	InitiatorName                    string        `json:"CHAIN_INITIATOR_NAME"`
	InitiatorPublicKey               string        `json:"CHAIN_INITIATOR_PUBLIC_KEY_STR"`
	InitSupply                       *types.UInt32 `json:"CHAIN_INIT_SUPPLY"`
	CommitteeAccount                 string        `json:"CHAIN_COMMITTEE_ACCOUNT"`
	CommitteePublicKey               string        `json:"CHAIN_COMMITTEE_PUBLIC_KEY_STR"`
	IrreversibleThreshold            *types.Int    `json:"CHAIN_IRREVERSIBLE_THRESHOLD"`
	MaxAccountNameLength             *types.Int    `json:"CHAIN_MAX_ACCOUNT_NAME_LENGTH"`
	MaxAccountWitnessVotes           *types.Int    `json:"CHAIN_MAX_ACCOUNT_WITNESS_VOTES"`
	BlockSize                        *types.Int    `json:"CHAIN_BLOCK_SIZE"`
	MaxCommentDepth                  *types.Int    `json:"CHAIN_MAX_COMMENT_DEPTH"`
	MaxMemoSize                      *types.Int    `json:"CHAIN_MAX_MEMO_SIZE"`
	MaxWitnesses                     *types.Int    `json:"CHAIN_MAX_WITNESSES"`
	MaxProxyRecursionDepth           *types.Int    `json:"CHAIN_MAX_PROXY_RECURSION_DEPTH"`
	MaxReserveRatio                  *types.Int    `json:"CHAIN_MAX_RESERVE_RATIO"`
	MaxSupportWitnesses              *types.Int    `json:"CHAIN_MAX_SUPPORT_WITNESSES"`
	MaxShareSupply                   string        `json:"CHAIN_MAX_SHARE_SUPPLY"`
	MaxSigCheckDepth                 *types.Int    `json:"CHAIN_MAX_SIG_CHECK_DEPTH"`
	MaxTimeUntilExpiration           *types.Int    `json:"CHAIN_MAX_TIME_UNTIL_EXPIRATION"`
	MaxTransactionSize               *types.Int    `json:"CHAIN_MAX_TRANSACTION_SIZE"`
	MaxUndoHistory                   *types.Int    `json:"CHAIN_MAX_UNDO_HISTORY"`
	MaxVoteChanges                   *types.Int    `json:"CHAIN_MAX_VOTE_CHANGES"`
	MaxTopWitnesses                  *types.Int    `json:"CHAIN_MAX_TOP_WITNESSES"`
	MaxWithdrawRoutes                *types.Int    `json:"CHAIN_MAX_WITHDRAW_ROUTES"`
	MaxWitnessURLLength              *types.Int    `json:"CHAIN_MAX_WITNESS_URL_LENGTH"`
	MinAccountCreationFee            *types.Int    `json:"CHAIN_MIN_ACCOUNT_CREATION_FEE"`
	MinAccountNameLength             *types.Int    `json:"CHAIN_MIN_ACCOUNT_NAME_LENGTH"`
	MinBlockSizeLimit                *types.Int    `json:"CHAIN_MIN_BLOCK_SIZE_LIMIT"`
	MaxBlockSizeLimit                *types.Int    `json:"CHAIN_MAX_BLOCK_SIZE_LIMIT"`
	NullAccount                      string        `json:"CHAIN_NULL_ACCOUNT"`
	NumInitiators                    *types.Int    `json:"CHAIN_NUM_INITIATORS"`
	ProxyToSelfAccount               string        `json:"CHAIN_PROXY_TO_SELF_ACCOUNT"`
	SecondsPerYear                   *types.Int    `json:"CHAIN_SECONDS_PER_YEAR"`
	VestingWithdrawIntervals         *types.Int    `json:"CHAIN_VESTING_WITHDRAW_INTERVALS"`
	VestingWithdrawIntervalSeconds   *types.Int    `json:"CHAIN_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	EnergyRegenerationSeconds        int           `json:"CHAIN_ENERGY_REGENERATION_SECONDS"`
	TokenSymbol                      *types.Int    `json:"TOKEN_SYMBOL"`
	SharesSymbol                     *types.Int    `json:"SHARES_SYMBOL"`
	ChainName                        string        `json:"CHAIN_NAME"`
}

//DynamicGlobalProperties structure for the GetDynamicGlobalProperties function.
type DynamicGlobalProperties struct {
	ID                         *types.Int   `json:"id"`
	HeadBlockNumber            uint32       `json:"head_block_number"`
	HeadBlockID                string       `json:"head_block_id"`
	GenesisTime                *types.Time  `json:"genesis_time"`
	Time                       *types.Time  `json:"time"`
	CurrentWitness             string       `json:"current_witness"`
	CommitteeFund              *types.Asset `json:"committee_fund"`
	CommitteeRequests          uint32       `json:"committee_requests"`
	CurrentSupply              *types.Asset `json:"current_supply"`
	TotalVersingFund           *types.Asset `json:"total_vesting_fund"`
	TotalVestingShares         *types.Asset `json:"total_vesting_shares"`
	TotalRewardFund            *types.Asset `json:"total_reward_fund"`
	TotalRewardShares          *types.Int64 `json:"total_reward_shares"`
	InflationCalcBlockNum      uint32       `json:"inflation_calc_block_num"`
	InflationWitnessPercent    int16        `json:"inflation_witness_percent"`
	InflationRatio             int16        `json:"inflation_ratio"`
	AverageBlockSize           uint32       `json:"average_block_size"`
	MaximumBlockSize           uint32       `json:"maximum_block_size"`
	CurrentAslot               uint64       `json:"current_aslot"`
	RecentSlotsFilled          *types.Int64 `json:"recent_slots_filled"`
	ParticipationCount         uint8        `json:"participation_count"`
	LastIrreversibleBlockNum   uint32       `json:"last_irreversible_block_num"`
	MaxVirtualBandwidth        uint64       `json:"max_virtual_bandwidth"`
	CurrentReserveRatio        uint64       `json:"current_reserve_ratio"`
	VoteRegenerationPerDay     uint32       `json:"vote_regeneration_per_day"`
	BandwidthReserveCandidates uint32       `json:"bandwidth_reserve_candidates"`
}

//BlockHeader structure for the GetBlockHeader and SetBlockAppliedCallback functions
type BlockHeader struct {
	Number                uint32        `json:"-"`
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	Witness               string        `json:"witness"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Extensions            []interface{} `json:"extensions"`
}

//Block structure for the GetBlock function
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
