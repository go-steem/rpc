package follow

import (
	"encoding/json"
	"github.com/asuleymanov/golos-go/types"
	"strconv"
	"strings"
)

type FollowObject struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}

type FeedEntry struct {
	Author   string      `json:"author"`
	Permlink string      `json:"permlink"`
	ReblogBy []string    `json:"reblog_by"`
	ReblogOn *types.Time `json:"reblog_on"`
	EntryID  *types.Int  `json:"entry_id"`
}

type Feeds struct {
	Comment  *CommentData `json:"comment"`
	ReblogBy []string     `json:"reblog_by"`
	ReblogOn *types.Time  `json:"reblog_on"`
	EntryID  *types.Int   `json:"entry_id"`
}

type FollowCount struct {
	Account        string     `json:"account"`
	FollowerCount  *types.Int `json:"follower_count"`
	FollowingCount *types.Int `json:"following_count"`
}

type BlogEntries struct {
	Author   string      `json:"author"`
	Permlink string      `json:"permlink"`
	Blog     string      `json:"blog"`
	ReblogOn *types.Time `json:"reblog_on"`
	EntryID  *types.Int  `json:"entry_id"`
}

type Blogs struct {
	Comment  *CommentData `json:"comment"`
	Blog     string       `json:"blog"`
	ReblogOn *types.Time  `json:"reblog_on"`
	EntryID  *types.Int   `json:"entry_id"`
}

type CommentData struct {
	ID                   *types.Int  `json:"id"`
	Author               string      `json:"author"`
	Permlink             string      `json:"permlink"`
	Category             string      `json:"category"`
	ParentAuthor         string      `json:"parent_author"`
	ParentPermlink       string      `json:"parent_permlink"`
	Title                string      `json:"title"`
	Body                 string      `json:"body"`
	JSONMetadata         *JSONcd     `json:"json_metadata"`
	LastUpdate           *types.Time `json:"last_update"`
	Created              *types.Time `json:"created"`
	Active               *types.Time `json:"active"`
	LastPayout           *types.Time `json:"last_payout"`
	Depth                *types.Int  `json:"depth"`
	Children             *types.Int  `json:"children"`
	ChildrenRshares2     string      `json:"children_rshares2"`
	NetRshares           *types.Int  `json:"net_rshares"`
	AbsRshares           *types.Int  `json:"abs_rshares"`
	VoteRshares          *types.Int  `json:"vote_rshares"`
	ChildrenAbsRshares   *types.Int  `json:"children_abs_rshares"`
	CashoutTime          *types.Time `json:"cashout_time"`
	MaxCashoutTime       *types.Time `json:"max_cashout_time"`
	TotalVoteWeight      *types.Int  `json:"total_vote_weight"`
	RewardWeight         *types.Int  `json:"reward_weight"`
	TotalPayoutValue     string      `json:"total_payout_value"`
	CuratorPayoutValue   string      `json:"curator_payout_value"`
	AuthorRewards        *types.Int  `json:"author_rewards"`
	NetVotes             *types.Int  `json:"net_votes"`
	RootComment          *types.Int  `json:"root_comment"`
	Mode                 string      `json:"mode"`
	MaxAcceptedPayout    string      `json:"max_accepted_payout"`
	PercentSteemDollars  *types.Int  `json:"percent_steem_dollars"`
	AllowReplies         bool        `json:"allow_replies"`
	AllowVotes           bool        `json:"allow_votes"`
	AllowCurationRewards bool        `json:"allow_curation_rewards"`
}

type JSONcd struct {
	Flag   bool
	Tags   []string `json:"tags"`
	Links  []string `json:"links"`
	Image  []string `json:"image"`
	App    string   `json:"app"`
	Format string   `json:"format"`
}

type JSONcdRaw struct {
	Tags   types.StringSlice `json:"tags"`
	Links  types.StringSlice `json:"links"`
	Image  types.StringSlice `json:"image"`
	App    string            `json:"app"`
	Format string            `json:"format"`
}

func (metadata *JSONcd) UnmarshalJSON(data []byte) error {
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
		var value JSONcd
		metadata = &value
		return nil
	}

	var raw JSONcdRaw
	if err := json.NewDecoder(strings.NewReader(unquoted)).Decode(&raw); err != nil {
		return err
	}

	metadata.Links = raw.Links
	metadata.Tags = raw.Tags
	metadata.Image = raw.Image
	metadata.App = raw.App
	metadata.Format = raw.Format

	return nil
}

type AccountReputation struct {
	Account    string      `json:"account"`
	Reputation interface{} `json:"reputation"`
}

type BlogAuthors struct {
	BlogAuthor []*BlogAuthor
}

type BlogAuthor struct {
	Name  string
	Value float64
}

func (b *BlogAuthors) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	for _, r := range v {
		rawRow := r.([]interface{})
		row := &BlogAuthor{rawRow[0].(string), rawRow[1].(float64)}
		b.BlogAuthor = append(b.BlogAuthor, row)
	}

	return nil
}
