package database

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/go-steem/rpc/apis/types"
)

type Config struct {
	SteemitBlockchainHardforkVersion string `json:"STEEMIT_BLOCKCHAIN_HARDFORK_VERSION"`
	SteemitBlockchainVersion         string `json:"STEEMIT_BLOCKCHAIN_VERSION"`
	SteemitBlockInterval             uint   `json:"STEEMIT_BLOCK_INTERVAL"`
}

type DynamicGlobalProperties struct {
	Time                     *types.Time `json:"time"`
	TotalPow                 *types.Int  `json:"total_pow"`
	NumPowWitnesses          *types.Int  `json:"num_pow_witnesses"`
	ConfidentialSupply       string      `json:"confidential_supply"`
	TotalVestingShares       string      `json:"total_vesting_shares"`
	CurrentReserveRatio      *types.Int  `json:"current_reserve_ratio"`
	Id                       *types.ID   `json:"id"`
	CurrentSupply            string      `json:"current_supply"`
	MaximumBlockSize         *types.Int  `json:"maximum_block_size"`
	RecentSlotsFilled        string      `json:"recent_slots_filled"`
	CurrentWitness           string      `json:"current_witness"`
	TotalRewardShares2       string      `json:"total_reward_shares2"`
	AverageBlockSize         *types.Int  `json:"average_block_size"`
	CurrentAslot             *types.Int  `json:"current_aslot"`
	LastIrreversibleBlockNum uint32      `json:"last_irreversible_block_num"`
	TotalVersingFundSteem    string      `json:"total_vesting_fund_steem"`
	HeadBlockId              string      `json:"head_block_id"`
	VirtualSupply            string      `json:"virtual_supply"`
	CurrentSBDSupply         string      `json:"current_sbd_supply"`
	ConfidentialSBDSupply    string      `json:"confidential_sbd_supply"`
	TotalRewardFundSteem     string      `json:"total_reward_fund_steem"`
	SBDInterestRate          *types.Int  `json:"sbd_interest_rate"`
	MaxVirtualBandwidth      string      `json:"max_virtual_bandwidth"`
	HeadBlockNumber          *types.Int  `json:"head_block_number"`
}

type Block struct {
	Number                uint32          `json:"-"`
	Timestamp             *types.Time     `json:"timestamp"`
	Witness               string          `json:"witness"`
	WitnessSignature      string          `json:"witness_signature"`
	TransactionMerkleRoot string          `json:"transaction_merkle_root"`
	Previous              string          `json:"previous"`
	Extensions            [][]interface{} `json:"extensions"`
	Transactions          []*Transaction  `json:"transactions"`
}

type Transaction struct {
	RefBlockNum    *types.Int   `json:"ref_block_num"`
	RefBlockPrefix *types.Int   `json:"ref_block_prefix"`
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
	Voter    string     `json:"voter"`
	Author   string     `json:"author"`
	Permlink string     `json:"permlink"`
	Weight   *types.Int `json:"weight"`
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

func (op *CommentOperation) IsStoryOperation() bool {
	return op.ParentAuthor == ""
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
	ActiveVotes             []*Vote          `json:"active_votes"`
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
	Users []string
	Tags  []string
	Image []string
}

type ContentMetadataRaw struct {
	Users []string `json:"users"`
	Tags  []string `json:"tags"`
	Image []string `json:"image"`
}

func (metadata *ContentMetadata) UnmarshalJSON(data []byte) error {
	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
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

type Vote struct {
	Voter  string     `json:"voter"`
	Weight *types.Int `json:"weight"`
}
