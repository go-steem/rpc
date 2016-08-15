package follow

const (
	FollowKindFollow = "blog"
	FollowKindIgnore = "ignore"
)

type FollowObject struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}

type FeedEntry struct {
	Author   string `json:"string"`
	Permlink string `json:"permlink"`
	EntryID  uint32 `json:"entry_id"`
}

/*
type CommentFeedEntry struct {
	Comment *CommentObject `json:"comment"`
	EntryID uint32         `json:"entry_id"`
}
*/

/*
type AccountReputation struct {
	Account string `json:"account"`
	Reputation ??? `json:"reputation"`
}
*/
