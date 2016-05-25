package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/powerman/rpc-codec/jsonrpc2"
	"golang.org/x/net/websocket"
)

const (
	OpTypeVote                 = "vote"
	OpTypeComment              = "comment"
	OpTypeTransfer             = "transfer"
	OpTypeTransferToVesting    = "transfer_to_vesting"
	OpTypeWithdrawVesting      = "withdraw_vesting"
	OpTypeLimitOrderCreate     = "limit_order_create"
	OpTypeLimitOrderCancel     = "limit_order_cancel"
	OpTypeFeedPublish          = "feed_publish"
	OpTypeConvert              = "convert"
	OpTypeAccountCreate        = "account_create"
	OpTypeAccountUpdate        = "account_update"
	OpTypeWitnessUpdate        = "witness_update"
	OpTypeAccountWitnessVote   = "account_witness_vote"
	OpTypeAccountWitnessProxy  = "account_witness_proxy"
	OpTypePow                  = "pow"
	OpTypeCustom               = "custom"
	OpTypeReportOverProduction = "report_over_production"
	OpTypeFullConvertRequest   = "fill_convert_request"
	OpTypeCommentReward        = "comment_reward"
	OpTypeCurateReward         = "curate_reward"
	OpTypeLiquidityReward      = "liquidity_reward"
	OpTypeInterest             = "interest"
	OpTypeFillVestingWithdraw  = "fill_vesting_withdraw"
	OpTypeFillOrder            = "fill_order"
)

type Time struct {
	*time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	parsed, err := time.Parse(`"2006-01-02T15:04:05"`, string(data))
	if err != nil {
		return err
	}
	t.Time = &parsed
	return nil
}

type Number struct {
	ValueInt    int64
	ValueString string
}

func (num *Number) UnmarshalJSON(data []byte) error {
	dataString := string(data)

	if data[0] == '"' {
		num.ValueString = dataString
		return nil
	}

	value, err := strconv.ParseInt(dataString, 0, 64)
	if err != nil {
		return err
	}
	num.ValueInt = value
	return nil
}

func (num *Number) String() string {
	if v := num.ValueString; v != "" {
		return v
	}
	return strconv.FormatInt(num.ValueInt, 10)
}

/*
Data:

STEEMIT_APR_PERCENT_SHIFT_PER_ROUND : 83
STEEMIT_BLOCKS_PER_DAY : 28800
STEEMIT_FEED_INTERVAL_BLOCKS : 1200
STEEMIT_LIQUIDITY_REWARD_PERIOD_SEC : 3600
STEEMIT_MAX_PROXY_RECURSION_DEPTH : 4
STEEMIT_MAX_TIME_UNTIL_EXPIRATION : 3600
STEEMIT_MAX_URL_LENGTH : 127
STEEMIT_MIN_ACCOUNT_CREATION_FEE : 1
STEEMIT_MIN_POW_REWARD : 1.000 STEEM
STEEMIT_SBD_INTEREST_COMPOUND_INTERVAL_SEC : 2.592e+06
STEEMIT_APR_PERCENT_MULTIPLY_PER_BLOCK : 102035135585887
STEEMIT_BLOCKCHAIN_PRECISION : 1000
STEEMIT_MAX_RESERVE_RATIO : 20000
STEEMIT_MAX_SHARE_SUPPLY : 1000000000000000
STEEMIT_MAX_UNDO_HISTORY : 10000
STEEMIT_MIN_LIQUIDITY_REWARD : 1200.000 STEEM
IS_TEST_NET : false
STEEMIT_BANDWIDTH_PRECISION : 1e+06
STEEMIT_BLOCK_INTERVAL : 3
STEEMIT_MINING_TIME : 2016-03-24T17:00:00
STEEMIT_BLOCKCHAIN_PRECISION_DIGITS : 3
STEEMIT_IRREVERSIBLE_THRESHOLD : 5100
STEEMIT_SYMBOL : STEEM
STEEMIT_TEMP_ACCOUNT : temp
STEEMIT_DEFAULT_SBD_INTEREST_RATE : 1000
STEEMIT_GENESIS_TIME : 2016-03-24T16:00:00
STEEMIT_MAX_ACCOUNT_NAME_LENGTH : 16
STEEMIT_APR_PERCENT_MULTIPLY_PER_HOUR : 119577151364285
STEEMIT_BANDWIDTH_AVERAGE_WINDOW_SECONDS : 604800
STEEMIT_INIT_MINER_NAME : initminer
STEEMIT_INIT_PUBLIC_KEY_STR : STM8GC13uCZbP44HzMLV6zPZGwVQ8Nt4Kji8PapsPiNq1BK153XTX
STEEMIT_LIQUIDITY_TIMEOUT_SEC : 604800000000
STEEMIT_MINER_ACCOUNT : miners
STEEMIT_APR_PERCENT_MULTIPLY_PER_ROUND : 133921203762304
STEEMIT_CURATE_APR_PERCENT : 3875
STEEMIT_VESTING_WITHDRAW_INTERVALS : 104
STEEMIT_ADDRESS_PREFIX : STM
STEEMIT_MAX_BLOCK_SIZE : 7.86432e+08
STEEMIT_MAX_FEED_AGE : 604800000000
STEEMIT_MAX_INSTANCE_ID : 281474976710655
STEEMIT_MIN_TRANSACTION_SIZE_LIMIT : 1024
STEEMIT_FEED_HISTORY_WINDOW : 168
STEEMIT_LIQUIDITY_REWARD_BLOCKS : 1200
STEEMIT_MAX_SIG_CHECK_DEPTH : 2
STEEMIT_MIN_BLOCK_SIZE_LIMIT : 131072
STEEMIT_MINER_PAY_PERCENT : 100
STEEMIT_POW_APR_PERCENT : 750
STEEMIT_PRODUCER_APR_PERCENT : 750
STEEMIT_PROXY_TO_SELF_ACCOUNT :
STEEMIT_START_VESTING_BLOCK : 201600
SBD_SYMBOL : 1.145197315e+09
STEEMIT_FREE_TRANSACTIONS_WITH_NEW_ACCOUNT : 100
STEEMIT_MAX_MEMO_SIZE : 2048
STEEMIT_MIN_FEEDS : 7
STEEMIT_VESTING_WITHDRAW_INTERVAL_SECONDS : 604800
GRAPHENE_CURRENT_DB_VERSION : GPH2.4
STEEMIT_100_PERCENT : 10000
STEEMIT_BLOCKS_PER_HOUR : 1200
STEEMIT_CONVERSION_DELAY : 604800000000
STEEMIT_INIT_SUPPLY : 0
STEEMIT_MAX_TRANSACTION_SIZE : 131072
STEEMIT_MIN_UNDO_HISTORY : 10
STEEMIT_VOTE_REGENERATION_SECONDS : 86400
STEEMIT_APR_PERCENT_SHIFT_PER_BLOCK : 87
STEEMIT_CONTENT_APR_PERCENT : 3875
STEEMIT_LIQUIDITY_APR_PERCENT : 750
STEEMIT_START_MINER_VOTING_BLOCK : 864000
VESTS_SYMBOL : 91621639411206
STEEMIT_APR_PERCENT_SHIFT_PER_HOUR : 77
STEEMIT_CASHOUT_WINDOW_SECONDS : 86400
STEEMIT_CHAIN_ID : 0000000000000000000000000000000000000000000000000000000000000000
STEEMIT_MAX_AUTHORITY_MEMBERSHIP : 10
STEEMIT_MAX_MINERS : 21
STEEMIT_MIN_CONTENT_REWARD : 1.000 STEEM
STEEMIT_MIN_CURATE_REWARD : 1.000 STEEM
STEEMIT_SECONDS_PER_YEAR : 3.1536e+07
STEEMIT_FIRST_CASHOUT_TIME : 2016-07-04T00:00:00
STEEMIT_MAX_RATION_DECAY_RATE : 1e+06
STEEMIT_MIN_LIQUIDITY_REWARD_PERIOD_SEC : 6e+07
STEEMIT_MIN_RATION : 100000
STEEMIT_MIN_TRANSACTION_EXPIRATION_LIMIT : 15
STEEMIT_NULL_ACCOUNT : null
STEEMIT_BLOCKS_PER_YEAR : 1.0512e+07
STEEMIT_MAX_COMMENT_DEPTH : 6
STEEMIT_MINING_REWARD : 1.000 STEEM
STEEMIT_MIN_PAYOUT_SBD : 0.020 SBD
STEEMIT_NUM_INIT_MINERS : 1
STEEM_SYMBOL : 84959911236355
STEEMIT_1_PERCENT : 100
STEEMIT_INIT_TIME : 1970-01-01T00:00:00
STEEMIT_MAX_ASSET_WHITELIST_AUTHORITIES : 10
STEEMIT_MIN_ACCOUNT_NAME_LENGTH : 3
STEEMIT_MIN_PRODUCER_REWARD : 1.000 STEEM
*/

type Config struct {
	SteemitAprPercentShiftPerRound int `json:"STEEMIT_APR_PERCENT_SHIFT_PER_ROUND"`
	SteemitBlocksPerDay            int `json:"STEEMIT_BLOCKS_PER_DAY"`
	SteemitFeedIntervalBlocks      int `json:"STEEMIT_FEED_INTERVAL_BLOCKS"`
	SteemitBlockInterval           int `json:"STEEMIT_BLOCK_INTERVAL"`
}

/*
STEEMIT_LIQUIDITY_REWARD_PERIOD_SEC : 3600
STEEMIT_MAX_PROXY_RECURSION_DEPTH : 4
STEEMIT_MAX_TIME_UNTIL_EXPIRATION : 3600
STEEMIT_MAX_URL_LENGTH : 127
STEEMIT_MIN_ACCOUNT_CREATION_FEE : 1
STEEMIT_MIN_POW_REWARD : 1.000 STEEM
STEEMIT_SBD_INTEREST_COMPOUND_INTERVAL_SEC : 2.592e+06
STEEMIT_APR_PERCENT_MULTIPLY_PER_BLOCK : 102035135585887
STEEMIT_BLOCKCHAIN_PRECISION : 1000
STEEMIT_MAX_RESERVE_RATIO : 20000
STEEMIT_MAX_SHARE_SUPPLY : 1000000000000000
STEEMIT_MAX_UNDO_HISTORY : 10000
STEEMIT_MIN_LIQUIDITY_REWARD : 1200.000 STEEM
IS_TEST_NET : false
STEEMIT_BANDWIDTH_PRECISION : 1e+06
STEEMIT_BLOCK_INTERVAL : 3
STEEMIT_MINING_TIME : 2016-03-24T17:00:00
STEEMIT_BLOCKCHAIN_PRECISION_DIGITS : 3
STEEMIT_IRREVERSIBLE_THRESHOLD : 5100
STEEMIT_SYMBOL : STEEM
STEEMIT_TEMP_ACCOUNT : temp
STEEMIT_DEFAULT_SBD_INTEREST_RATE : 1000
STEEMIT_GENESIS_TIME : 2016-03-24T16:00:00
STEEMIT_MAX_ACCOUNT_NAME_LENGTH : 16
STEEMIT_APR_PERCENT_MULTIPLY_PER_HOUR : 119577151364285
STEEMIT_BANDWIDTH_AVERAGE_WINDOW_SECONDS : 604800
STEEMIT_INIT_MINER_NAME : initminer
STEEMIT_INIT_PUBLIC_KEY_STR : STM8GC13uCZbP44HzMLV6zPZGwVQ8Nt4Kji8PapsPiNq1BK153XTX
STEEMIT_LIQUIDITY_TIMEOUT_SEC : 604800000000
STEEMIT_MINER_ACCOUNT : miners
STEEMIT_APR_PERCENT_MULTIPLY_PER_ROUND : 133921203762304
STEEMIT_CURATE_APR_PERCENT : 3875
STEEMIT_VESTING_WITHDRAW_INTERVALS : 104
STEEMIT_ADDRESS_PREFIX : STM
STEEMIT_MAX_BLOCK_SIZE : 7.86432e+08
STEEMIT_MAX_FEED_AGE : 604800000000
STEEMIT_MAX_INSTANCE_ID : 281474976710655
STEEMIT_MIN_TRANSACTION_SIZE_LIMIT : 1024
STEEMIT_FEED_HISTORY_WINDOW : 168
STEEMIT_LIQUIDITY_REWARD_BLOCKS : 1200
STEEMIT_MAX_SIG_CHECK_DEPTH : 2
STEEMIT_MIN_BLOCK_SIZE_LIMIT : 131072
STEEMIT_MINER_PAY_PERCENT : 100
STEEMIT_POW_APR_PERCENT : 750
STEEMIT_PRODUCER_APR_PERCENT : 750
STEEMIT_PROXY_TO_SELF_ACCOUNT :
STEEMIT_START_VESTING_BLOCK : 201600
SBD_SYMBOL : 1.145197315e+09
STEEMIT_FREE_TRANSACTIONS_WITH_NEW_ACCOUNT : 100
STEEMIT_MAX_MEMO_SIZE : 2048
STEEMIT_MIN_FEEDS : 7
STEEMIT_VESTING_WITHDRAW_INTERVAL_SECONDS : 604800
GRAPHENE_CURRENT_DB_VERSION : GPH2.4
STEEMIT_100_PERCENT : 10000
STEEMIT_BLOCKS_PER_HOUR : 1200
STEEMIT_CONVERSION_DELAY : 604800000000
STEEMIT_INIT_SUPPLY : 0
STEEMIT_MAX_TRANSACTION_SIZE : 131072
STEEMIT_MIN_UNDO_HISTORY : 10
STEEMIT_VOTE_REGENERATION_SECONDS : 86400
STEEMIT_APR_PERCENT_SHIFT_PER_BLOCK : 87
STEEMIT_CONTENT_APR_PERCENT : 3875
STEEMIT_LIQUIDITY_APR_PERCENT : 750
STEEMIT_START_MINER_VOTING_BLOCK : 864000
VESTS_SYMBOL : 91621639411206
STEEMIT_APR_PERCENT_SHIFT_PER_HOUR : 77
STEEMIT_CASHOUT_WINDOW_SECONDS : 86400
STEEMIT_CHAIN_ID : 0000000000000000000000000000000000000000000000000000000000000000
STEEMIT_MAX_AUTHORITY_MEMBERSHIP : 10
STEEMIT_MAX_MINERS : 21
STEEMIT_MIN_CONTENT_REWARD : 1.000 STEEM
STEEMIT_MIN_CURATE_REWARD : 1.000 STEEM
STEEMIT_SECONDS_PER_YEAR : 3.1536e+07
STEEMIT_FIRST_CASHOUT_TIME : 2016-07-04T00:00:00
STEEMIT_MAX_RATION_DECAY_RATE : 1e+06
STEEMIT_MIN_LIQUIDITY_REWARD_PERIOD_SEC : 6e+07
STEEMIT_MIN_RATION : 100000
STEEMIT_MIN_TRANSACTION_EXPIRATION_LIMIT : 15
STEEMIT_NULL_ACCOUNT : null
STEEMIT_BLOCKS_PER_YEAR : 1.0512e+07
STEEMIT_MAX_COMMENT_DEPTH : 6
STEEMIT_MINING_REWARD : 1.000 STEEM
STEEMIT_MIN_PAYOUT_SBD : 0.020 SBD
STEEMIT_NUM_INIT_MINERS : 1
STEEM_SYMBOL : 84959911236355
STEEMIT_1_PERCENT : 100
STEEMIT_INIT_TIME : 1970-01-01T00:00:00
STEEMIT_MAX_ASSET_WHITELIST_AUTHORITIES : 10
STEEMIT_MIN_ACCOUNT_NAME_LENGTH : 3
STEEMIT_MIN_PRODUCER_REWARD : 1.000 STEEM
*/

/*
Data:

time : 2016-05-11T07:19:51
total_pow : 6.364549e+06
num_pow_witnesses : 92
confidential_supply : 0.000 STEEM
total_vesting_shares : 421708213897.314264 VESTS
current_reserve_ratio : 20000
id : 2.0.0
current_supply : 42061544.000 STEEM
maximum_block_size : 131072
recent_slots_filled : 340282366920938463463374607431768211455
current_witness : abit
total_reward_shares2 : 67867182200021104017152081459503
average_block_size : 116
current_aslot : 1.371997e+06
last_irreversible_block_num : 1.348204e+06
total_vesting_fund_steem : 38879546.582 STEEM
head_block_id : 00149276dc1c55ae8a430b9e750f5f754b5806f6
virtual_supply : 42061544.000 STEEM
current_sbd_supply : 0.000 SBD
confidential_sbd_supply : 0.000 SBD
total_reward_fund_steem : 2696428.000 STEEM
sbd_interest_rate : 1000
max_virtual_bandwidth : 5824555244896037546
head_block_number : 1.348214e+06
*/

// XXX: Make sure the types are correct.
type DynamicGlobalProperties struct {
	Time                     Time    `json:"time"`
	TotalPow                 int     `json:"total_pow"`
	NumPowWitnesses          int     `json:"num_pow_witnesses"`
	ConfidentialSupply       string  `json:"confidential_supply"`
	TotalVestingShares       string  `json:"total_vesting_shares"`
	CurrentReserveRatio      float64 `json:"current_reserve_ratio"`
	Id                       string  `json:"id"`
	CurrentSupply            string  `json:"current_supply"`
	MaximumBlockSize         int     `json:"maximum_block_size"`
	RecentSlotsFilled        string  `json:"recent_slots_filled"`
	CurrentWitness           string  `json:"current_witness"`
	TotalRewardShares2       string  `json:"total_reward_shares2"`
	AverageBlockSize         int     `json:"average_block_size"`
	CurrentAslot             int     `json:"current_aslot"`
	LastIrreversibleBlockNum int     `json:"last_irreversible_block_num"`
	TotalVersingFundSteem    string  `json:"total_vesting_fund_steem"`
	HeadBlockId              string  `json:"head_block_id"`
	VirtualSupply            string  `json:"virtual_supply"`
	CurrentSBDSupply         string  `json:"current_sbd_supply"`
	ConfidentialSBDSupply    string  `json:"confidential_sbd_supply"`
	TotalRewardFundSteem     string  `json:"total_reward_fund_steem"`
	SBDInterestRate          float64 `json:"sbd_interest_rate"`
	MaxVirtualBandwidth      string  `json:"max_virtual_bandwidth"`
	HeadBlockNumber          int     `json:"head_block_number"`
}

type Block struct {
	Timestamp             Time                     `json:"timestamp"`
	Witness               string                   `json:"witness"`
	WitnessSignature      string                   `json:"witness_signature"`
	TransactionMerkleRoot string                   `json:"transaction_merkle_root"`
	Previous              string                   `json:"previous"`
	Extensions            []map[string]interface{} `json:"extensions"`
	Transactions          []*Transaction           `json:"transactions"`
}

type Transaction struct {
	RefBlockNum    int          `json:"ref_block_num"`
	RefBlockPrefix int          `json:"ref_block_prefix"`
	Expiration     string       `json:"expiration"`
	Operations     []*Operation `json:"operations"`
}

type Operation struct {
	Type string
	Body interface{}
}

func (op *Operation) UnmarshalJSON(data []byte) error {
	raw := make([]json.RawMessage, 2)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) != 2 {
		return errors.New("invalid transaction object")
	}

	var operationType string
	if err := json.Unmarshal(raw[0], &operationType); err != nil {
		return err
	}

	switch operationType {
	case OpTypeVote:
		var body VoteOperation
		if err := json.Unmarshal(raw[1], &body); err != nil {
			return err
		}
		op.Body = &body
	case OpTypeComment:
		var body CommentOperation
		if err := json.Unmarshal(raw[1], &body); err != nil {
			return err
		}
		op.Body = &body
	default:
		var body map[string]interface{}
		if err := json.Unmarshal(raw[1], &body); err != nil {
			return err
		}
		op.Body = body
	}

	op.Type = operationType
	return nil
}

type VoteOperation struct {
	Voter    string `json:"voter"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
	Weight   int    `json:"weight"`
}

// CommentOperation represents either a new post or a comment.
//
// In case Title is filled in and ParentAuthor is empty, it is a new post.
// The post category can be read from ParentPermlink.
//
// In case the author is just updating an existing post,
// Body contains only the diff against the original content.
type CommentOperation struct {
	Author         string `json:"author"`
	Title          string `json:"title"`
	Permlink       string `json:"permlink"`
	ParentAuthor   string `json:"parent_author"`
	ParentPermlink string `json:"parent_permlink"`
	Body           string `json:"body"`
}

func (op *CommentOperation) IsNewStory() bool {
	return op.ParentAuthor == ""
}

func (op *CommentOperation) Link() string {
	if op.IsNewStory() {
		return fmt.Sprintf("https://steemit.com/%v/@%v/%v", op.ParentPermlink, op.Author, op.Permlink)
	}

	return ""
}

type Client struct {
	endpoint *jsonrpc2.Client
}

func Dial(addr string) (*Client, error) {
	// Connect to the given address.
	conn, err := websocket.Dial(addr, "", "http://localhost")
	if err != nil {
		return nil, err
	}

	// Instantiate a JSON-RPC client.
	endpoint := jsonrpc2.NewClient(conn)

	// Return a new Client instance.
	return &Client{endpoint}, nil
}

func (client *Client) GetConfig() (*Config, error) {
	var resp Config
	if err := client.endpoint.Call("get_config", []string{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *Client) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	var resp DynamicGlobalProperties
	err := client.endpoint.Call("get_dynamic_global_properties", []string{}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *Client) GetBlock(blockNumber int) (*Block, error) {
	var resp Block
	if err := client.endpoint.Call("get_block", []int{blockNumber}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * Content
 */

/*
EXAMPLE:

{
    "id": 65,
    "result": {
        "abs_rshares": "64813426102199",
        "active": "2016-05-24T15:14:15",
        "active_votes": [
            {
                "voter": "dantheman",
                "weight": "21481034324870"
            },
            {
                "voter": "hr1",
                "weight": "19909334576600"
            },
            {
                "voter": "markopaasila",
                "weight": 200331
            },
            {
                "voter": "void",
                "weight": 41
            },
            {
                "voter": "nenad-ristic",
                "weight": 0
            },
            {
                "voter": "coltmerg420",
                "weight": 196
            }
        ],
        "author": "void",
        "body": "",
        "cashout_time": "2016-05-25T15:14:15",
        "category": "spirituality",
        "children": 0,
        "children_rshares2": "4200780203105210657992635601",
        "created": "2016-05-24T12:45:18",
        "depth": 0,
        "id": "2.8.9503",
        "json_metadata": "",
        "last_update": "2016-05-24T15:14:15",
        "net_rshares": "64813426102199",
        "parent_author": "",
        "parent_permlink": "spirituality",
        "pending_payout_value": "22.093 SBD",
        "permlink": "we-are-taught-yet-we-do-not-understand",
        "replies": [],
        "root_title": "",
        "stats": "2.9.9503",
        "title": "",
        "total_payout_value": "0.000 SBD",
        "total_pending_payout_value": "44.187 SBD",
        "total_vote_weight": "41390369102038",
        "url": "/spirituality/@void/we-are-taught-yet-we-do-not-understand"
    }
}
*/

type Content struct {
	Id                      *Number                  `json:"id"`
	RootTitle               string                   `json:"root_title"`
	Active                  *Time                    `json:"active"`
	AbsRshares              *Number                  `json:"abs_rshares"`
	PendingPayoutValue      string                   `json:"pending_payout_value"`
	TotalPendingPayoutValue string                   `json:"total_pending_payout_value"`
	Category                string                   `json:"category"`
	Title                   string                   `json:"title"`
	LastUpdate              *Time                    `json:"last_update"`
	Stats                   string                   `json:"stats"`
	Body                    string                   `json:"body"`
	Created                 *Time                    `json:"created"`
	Replies                 []map[string]interface{} `json:"replies"`
	Permlink                string                   `json:"permlink"`
	JsonMetadata            string                   `json:"json_metadata"`
	Children                int                      `json:"children"`
	NetRshares              *Number                  `json:"net_rshares"`
	URL                     string                   `json:"url"`
	ActiveVotes             []*Vote                  `json:"active_votes"`
	ParentPermlink          string                   `json:"parent_permlink"`
	CashoutTime             *Time                    `json:"cashout_time"`
	TotalPayoutValue        string                   `json:"total_payout_value"`
	ParentAuthor            string                   `json:"parent_author"`
	ChildrenRshares2        *Number                  `json:"children_rshares2"`
	Author                  string                   `json:"author"`
	Depth                   *Number                  `json:"depth"`
	TotalVoteWeight         *Number                  `json:"total_vote_weight"`
}

type Vote struct {
	Voter  string  `json:"voter"`
	Weight *Number `json:"weight"`
}

func (client *Client) GetContent(author, permlink string) (*Content, error) {
	var resp Content
	err := client.endpoint.Call("get_content", []string{author, permlink}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *Client) Close() error {
	// Close the underlying WebSocket connection.
	return client.endpoint.Close()
}
