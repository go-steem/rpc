package follow

import (
	"fmt"

	"github.com/asuleymanov/steem-go/transports"
)

const apiID = "follow_api"

//API plug-in structure
type API struct {
	caller transports.Caller
}

//NewAPI plug-in initialization
func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{apiID, method, params}, resp)
}

//GetAccountReputations api request get_account_reputations
func (api *API) GetAccountReputations(accounts []string) ([]*uint32, error) {
	var resp []*uint32
	err := api.call("get_account_reputations", []interface{}{accounts}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetBlog api request get_blog
func (api *API) GetBlog(accountName string, entryID uint32, limit uint16) ([]*Blogs, error) {
	if limit > 500 {
		return nil, fmt.Errorf("%v: get_blog -> limit must not exceed 500", apiID)
	}
	var resp []*Blogs
	err := api.call("get_blog", []interface{}{accountName, entryID, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetBlogAuthors api request get_blog_authors
func (api *API) GetBlogAuthors(author string) (*BlogAuthors, error) {
	var resp BlogAuthors
	err := api.call("get_blog_authors", []interface{}{author}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetBlogEntries api request get_blog_entries
func (api *API) GetBlogEntries(accountName string, entryID uint32, limit uint16) ([]*BlogEntries, error) {
	if limit > 500 {
		return nil, fmt.Errorf("%v: get_blog_entries -> limit must not exceed 500", apiID)
	}
	var resp []*BlogEntries
	err := api.call("get_blog_entries", []interface{}{accountName, entryID, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetFeed api request get_feed
func (api *API) GetFeed(accountName string, entryID uint32, limit uint16) ([]*Feeds, error) {
	if limit > 500 {
		return nil, fmt.Errorf("%v: get_feed -> limit must not exceed 500", apiID)
	}
	var resp []*Feeds
	err := api.call("get_feed", []interface{}{accountName, entryID, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetFeedEntries api request get_feed_entries
func (api *API) GetFeedEntries(accountName string, entryID uint32, limit uint16) ([]*FeedEntry, error) {
	if limit > 500 {
		return nil, fmt.Errorf("%v: get_feed_entries -> limit must not exceed 500", apiID)
	}
	var resp []*FeedEntry
	err := api.call("get_feed_entries", []interface{}{accountName, entryID, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetFollowCount api request get_follow_count
func (api *API) GetFollowCount(accountName string) (*FollowCount, error) {
	var resp FollowCount
	err := api.call("get_follow_count", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetFollowers api request get_followers
/*
kind:
undefined
blog
ignore
*/
func (api *API) GetFollowers(accountName, start, kind string, limit uint16) ([]*FollowObject, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("%v: get_followers -> limit must not exceed 1000", apiID)
	}
	if kind != "undefined" || kind != "blog" || kind != "ignore" {
		return nil, fmt.Errorf("%v: get_followers -> kind can take values only \"undefined\", \"blog\" and \"ignore\"", apiID)
	}
	var resp []*FollowObject
	err := api.call("get_followers", []interface{}{accountName, start, kind, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetFollowing api request get_following
/*
kind:
undefined
blog
ignore
*/
func (api *API) GetFollowing(accountName, start, kind string, limit uint16) ([]*FollowObject, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("%v: get_following -> limit must not exceed 1000", apiID)
	}
	if kind != "undefined" || kind != "blog" || kind != "ignore" {
		return nil, fmt.Errorf("%v: get_following -> kind can take values only \"undefined\", \"blog\" and \"ignore\"", apiID)
	}
	var resp []*FollowObject
	err := api.call("get_following", []interface{}{accountName, start, kind, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetRebloggedBy api request get_reblogged_by
func (api *API) GetRebloggedBy(author, permlink string) ([]*string, error) {
	var resp []*string
	err := api.call("get_reblogged_by", []interface{}{author, permlink}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
