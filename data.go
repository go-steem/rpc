package rpc

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Config struct {
	SteemitBlockInterval uint `json:"STEEMIT_BLOCK_INTERVAL"`
}

type DynamicGlobalProperties struct {
	Time                     *Time  `json:"time"`
	TotalPow                 *Int   `json:"total_pow"`
	NumPowWitnesses          *Int   `json:"num_pow_witnesses"`
	ConfidentialSupply       string `json:"confidential_supply"`
	TotalVestingShares       string `json:"total_vesting_shares"`
	CurrentReserveRatio      *Int   `json:"current_reserve_ratio"`
	Id                       *Int   `json:"id"`
	CurrentSupply            string `json:"current_supply"`
	MaximumBlockSize         *Int   `json:"maximum_block_size"`
	RecentSlotsFilled        string `json:"recent_slots_filled"`
	CurrentWitness           string `json:"current_witness"`
	TotalRewardShares2       string `json:"total_reward_shares2"`
	AverageBlockSize         *Int   `json:"average_block_size"`
	CurrentAslot             *Int   `json:"current_aslot"`
	LastIrreversibleBlockNum *Int   `json:"last_irreversible_block_num"`
	TotalVersingFundSteem    string `json:"total_vesting_fund_steem"`
	HeadBlockId              string `json:"head_block_id"`
	VirtualSupply            string `json:"virtual_supply"`
	CurrentSBDSupply         string `json:"current_sbd_supply"`
	ConfidentialSBDSupply    string `json:"confidential_sbd_supply"`
	TotalRewardFundSteem     string `json:"total_reward_fund_steem"`
	SBDInterestRate          *Int   `json:"sbd_interest_rate"`
	MaxVirtualBandwidth      string `json:"max_virtual_bandwidth"`
	HeadBlockNumber          *Int   `json:"head_block_number"`
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
	RefBlockNum    *Int         `json:"ref_block_num"`
	RefBlockPrefix *Int         `json:"ref_block_prefix"`
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
	Weight   *Int   `json:"weight"`
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

type Content struct {
	Id                      *Int                     `json:"id"`
	RootTitle               string                   `json:"root_title"`
	Active                  *Time                    `json:"active"`
	AbsRshares              *Int                     `json:"abs_rshares"`
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
	Children                *Int                     `json:"children"`
	NetRshares              *Int                     `json:"net_rshares"`
	URL                     string                   `json:"url"`
	ActiveVotes             []*Vote                  `json:"active_votes"`
	ParentPermlink          string                   `json:"parent_permlink"`
	CashoutTime             *Time                    `json:"cashout_time"`
	TotalPayoutValue        string                   `json:"total_payout_value"`
	ParentAuthor            string                   `json:"parent_author"`
	ChildrenRshares2        *Int                     `json:"children_rshares2"`
	Author                  string                   `json:"author"`
	Depth                   *Int                     `json:"depth"`
	TotalVoteWeight         *Int                     `json:"total_vote_weight"`
}

type Vote struct {
	Voter  string `json:"voter"`
	Weight *Int   `json:"weight"`
}
