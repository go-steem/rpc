package follow

import (
	"encoding/json"

	"github.com/asuleymanov/steem-go/transports"
	"github.com/pkg/errors"
)

const apiID = "follow_api"

type API struct {
	caller transports.Caller
}

func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var emptyParams = struct{}{}

func (api *API) raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{apiID, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "steem-go: %v: failed to call %v\n", apiID, method)
	}
	return &resp, nil
}

//get_followers
func (api *API) GetFollowers(accountName, start, kind string, limit uint16) ([]*FollowObject, error) {
	raw, err := api.raw("get_followers", []interface{}{accountName, start, kind, limit})
	if err != nil {
		return nil, err
	}
	var resp []*FollowObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: follow_api: failed to unmarshal get_followers response")
	}
	return resp, nil
}

//get_following
func (api *API) GetFollowing(accountName, start, kind string, limit uint16) ([]*FollowObject, error) {
	raw, err := api.raw("get_following", []interface{}{accountName, start, kind, limit})
	if err != nil {
		return nil, err
	}
	var resp []*FollowObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: follow_api: failed to unmarshal get_following response")
	}
	return resp, nil
}

//get_follow_count
func (api *API) GetFollowCount(accountName string) (*FollowCount, error) {
	raw, err := api.raw("get_follow_count", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp *FollowCount
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: follow_api: failed to unmarshal get_follow_count response")
	}
	return resp, nil
}

//get_feed_entries
func (api *API) GetFeedEntries(accountName string, entryID uint32, limit uint16) ([]*FeedEntry, error) {
	if limit > 500 {
		return nil, errors.New("steem-go: follow_api: get_feed_entries -> limit must not exceed 500")
	}
	raw, err := api.raw("get_feed_entries", []interface{}{accountName, entryID, limit})
	if err != nil {
		return nil, err
	}
	var resp []*FeedEntry
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: follow_api: failed to unmarshal get_feed_entries response")
	}
	return resp, nil
}

//get_feed
func (api *API) GetFeed(accountName string, entryID uint32, limit uint16) ([]*Feeds, error) {
	if limit > 500 {
		return nil, errors.New("steem-go: follow_api: get_feed -> limit must not exceed 500")
	}
	raw, err := api.raw("get_feed", []interface{}{accountName, entryID, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Feeds
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: follow_api: failed to unmarshal get_feed response")
	}
	return resp, nil
}

//get_blog_entries
func (api *API) GetBlogEntries(accountName string, entryID uint32, limit uint16) ([]*BlogEntries, error) {
	if limit > 500 {
		return nil, errors.New("steem-go: follow_api: get_blog_entries -> limit must not exceed 500")
	}
	raw, err := api.raw("get_blog_entries", []interface{}{accountName, entryID, limit})
	if err != nil {
		return nil, err
	}
	var resp []*BlogEntries
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: follow_api: failed to unmarshal get_feed_entries response")
	}
	return resp, nil
}

//get_blog
func (api *API) GetBlog(accountName string, entryID uint32, limit uint16) ([]*Blogs, error) {
	if limit > 500 {
		return nil, errors.New("steem-go: follow_api: get_blog -> limit must not exceed 500")
	}
	raw, err := api.raw("get_blog", []interface{}{accountName, entryID, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Blogs
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: follow_api: failed to unmarshal get_feed response")
	}
	return resp, nil
}

//get_account_reputations
func (api *API) GetAccountReputations(lowerBoundName string, limit uint32) ([]*AccountReputation, error) {
	if limit > 1000 {
		return nil, errors.New("steem-go: follow_api: get_account_reputations -> limit must not exceed 1000")
	}
	raw, err := api.raw("get_account_reputations", []interface{}{lowerBoundName, limit})
	if err != nil {
		return nil, err
	}
	var resp []*AccountReputation
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: follow_api: failed to unmarshal get_account_reputations response")
	}
	return resp, nil
}

//get_reblogged_by
func (api *API) GetRebloggedBy(author, permlink string) ([]*string, error) {
	raw, err := api.raw("get_reblogged_by", []interface{}{author, permlink})
	if err != nil {
		return nil, err
	}
	var resp []*string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: market_history_api: failed to unmarshal get_reblogged_by response")
	}
	return resp, nil
}

//get_blog_authors
func (api *API) GetBlogAuthors(author string) (*BlogAuthors, error) {
	raw, err := api.raw("get_blog_authors", []interface{}{author})
	if err != nil {
		return nil, err
	}
	var resp BlogAuthors
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "steem-go: market_history_api: failed to unmarshal get_blog_authors response")
	}
	return &resp, nil
}
