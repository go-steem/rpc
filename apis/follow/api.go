package follow

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/go-steem/rpc/interfaces"

	// Vendor
	"github.com/pkg/errors"
)

type API struct {
	caller interfaces.Caller
}

func NewAPI(caller interfaces.Caller) *API {
	return &API{caller}
}

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{"follow_api", method, params}, resp)
}

func (api *API) GetFollowersRaw(
	accountName string,
	start string,
	kind string,
	limit uint16,
) (*json.RawMessage, error) {

	var resp json.RawMessage
	params := []interface{}{accountName, start, kind, limit}
	if err := api.call("get_followers", params, &resp); err != nil {
		return nil, errors.Wrap(err, "go-steem/rpc: follow_api: failed to call get_followers")
	}
	return &resp, nil
}

func (api *API) GetFollowers(
	accountName string,
	start string,
	kind string,
	limit uint16,
) ([]*FollowObject, error) {

	raw, err := api.GetFollowersRaw(accountName, start, kind, limit)
	if err != nil {
		return nil, err
	}

	var resp []*FollowObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(
			err, "go-steem/rpc: follow_api: failed to unmarshal get_followers response")
	}
	return resp, nil

}

func (api *API) GetFollowingRaw(
	accountName string,
	start string,
	kind string,
	limit uint16,
) (*json.RawMessage, error) {

	var resp json.RawMessage
	params := []interface{}{accountName, start, kind, limit}
	if err := api.call("get_following", params, &resp); err != nil {
		return nil, errors.Wrap(err, "go-steem/rpc: follow_api: failed to call get_following")
	}
	return &resp, nil
}

func (api *API) GetFollowing(
	accountName string,
	start string,
	kind string,
	limit uint16,
) ([]*FollowObject, error) {

	raw, err := api.GetFollowingRaw(accountName, start, kind, limit)
	if err != nil {
		return nil, err
	}

	var resp []*FollowObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(
			err, "go-steem/rpc: follow_api: failed to unmarshal get_following response")
	}
	return resp, nil
}

func (api *API) GetFeedEntriesRaw(
	accountName string,
	entryID uint32,
	limit uint16,
) (*json.RawMessage, error) {

	if limit == 0 {
		limit = 500
	}

	var resp json.RawMessage
	params := []interface{}{accountName, entryID, limit}
	if err := api.call("get_feed_entries", params, &resp); err != nil {
		return nil, errors.Wrap(err, "go-steem/rpc: follow_api: failed to call get_feed_entries")
	}
	return &resp, nil
}

func (api *API) GetFeedEntries(
	accountName string,
	entryID uint32,
	limit uint16,
) ([]*FeedEntry, error) {

	raw, err := api.GetFeedEntriesRaw(accountName, entryID, limit)
	if err != nil {
		return nil, err
	}

	var resp []*FeedEntry
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(
			err, "go-steem/rpc: follow_api: failed to unmarshal get_feed_entries response")
	}
	return resp, nil
}

func (api *API) GetFeedRaw(
	accountName string,
	entryID uint32,
	limit uint16,
) (*json.RawMessage, error) {

	if limit == 0 {
		limit = 500
	}

	var resp json.RawMessage
	params := []interface{}{accountName, entryID, limit}
	if err := api.call("get_feed", params, &resp); err != nil {
		return nil, errors.Wrap(err, "go-steem/rpc: follow_api: failed to call get_feed")
	}
	return &resp, nil
}

func (api *API) GetAccountReputationsRaw(
	lowerBoundName string,
	limit uint32,
) (*json.RawMessage, error) {

	if limit == 0 {
		limit = 1000
	}

	var resp json.RawMessage
	params := []interface{}{lowerBoundName, limit}
	if err := api.call("get_account_reputations", params, &resp); err != nil {
		return nil, errors.Wrap(
			err, "go-steem/rpc: follow_api: failed to call get_account_reputations")
	}
	return &resp, nil
}
