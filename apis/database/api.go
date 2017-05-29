package database

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/golos-go/interfaces"
	"github.com/asuleymanov/golos-go/internal/call"

	// Vendor
	"github.com/pkg/errors"
)

const (
	APIID         = "database_api"
	NumbericAPIID = 0
)

type API struct {
	caller interfaces.Caller
}

func NewAPI(caller interfaces.Caller) *API {
	return &API{caller}
}

/*
   // Subscriptions
   (set_subscribe_callback)
   (set_pending_transaction_callback)
   (set_block_applied_callback)
   (cancel_all_subscriptions)
*/

/*
   // Tags
   (get_trending_tags)
   (get_discussions_by_trending)
   (get_discussions_by_created)
   (get_discussions_by_active)
   (get_discussions_by_cashout)
   (get_discussions_by_payout)
   (get_discussions_by_votes)
   (get_discussions_by_children)
   (get_discussions_by_hot)
   (get_recommended_for)
*/

func (api *API) GetTrendingTagsRaw(afterTag string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_trending_tags", []interface{}{afterTag, limit})
}

type DiscussionQuery struct {
	Tag   string `json:"tag"`
	Limit uint32 `json:"limit"`
	// XXX: Not sure about the type here.
	FilterTags     []string `json:"filter_tags"`
	StartAuthor    string   `json:"start_author,omitempty"`
	StartPermlink  string   `json:"start_permlink,omitempty"`
	ParentAuthor   string   `json:"parent_author,omitempty"`
	ParentPermlink string   `json:"parent_permlink"`
}

func (api *API) GetDiscussionsByTrendingRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_trending", query)
}

func (api *API) GetDiscussionsByCreatedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_created", query)
}

func (api *API) GetDiscussionsByActiveRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_active", query)
}

func (api *API) GetDiscussionsByCashoutRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_cashout", query)
}

func (api *API) GetDiscussionsByPayoutRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_payout", query)
}

func (api *API) GetDiscussionsByVotesRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_votes", query)
}

func (api *API) GetDiscussionsByChildrenRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_children", query)
}

func (api *API) GetDiscussionsByHotRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_hot", query)
}

func (api *API) GetRecommendedForRaw(user string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_votes", []interface{}{user, limit})
}

/*
   // Blocks and transactions
   (get_block_header)
   (get_block)
   (get_state)
   (get_trending_categories)
   (get_best_categories)
   (get_active_categories)
   (get_recent_categories)
*/

func (api *API) GetBlockHeaderRaw(blockNum uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_block_header", []uint32{blockNum})
}

func (api *API) GetBlockRaw(blockNum uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_block", []uint32{blockNum})
}

func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	var resp Block
	if err := api.caller.Call("get_block", []uint32{blockNum}, &resp); err != nil {
		return nil, err
	}
	resp.Number = blockNum
	return &resp, nil
}

func (api *API) GetStateRaw(path string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_state", []string{path})
}

func (api *API) GetTrendingCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_trending_categories", []interface{}{after, limit})
}

func (api *API) GetBestCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_best_categories", []interface{}{after, limit})
}

func (api *API) GetActiveCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_active_categories", []interface{}{after, limit})
}

func (api *API) GetRecentCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_recent_categories", []interface{}{after, limit})
}

/*
   // Globals
   (get_config)
   (get_dynamic_global_properties)
   (get_chain_properties)
   (get_feed_history)
   (get_current_median_history_price)
   (get_witness_schedule)
   (get_hardfork_version)
   (get_next_scheduled_hardfork)
*/

func (api *API) GetConfigRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_config", call.EmptyParams)
}

func (api *API) GetConfig() (*Config, error) {
	var resp Config
	if err := api.caller.Call("get_config", call.EmptyParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (api *API) GetDynamicGlobalPropertiesRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_dynamic_global_properties", call.EmptyParams)
}

func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	var resp DynamicGlobalProperties
	if err := api.caller.Call("get_dynamic_global_properties", call.EmptyParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (api *API) GetChainPropertiesRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_chain_properties", call.EmptyParams)
}

func (api *API) GetFeedHistoryRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_feed_history", call.EmptyParams)
}

func (api *API) GetCurrentMedianHistoryPriceRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_current_median_history_price", call.EmptyParams)
}

func (api *API) GetWitnessScheduleRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_witness_schedule", call.EmptyParams)
}

func (api *API) GetHardforkVersionRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_hardfork_version", call.EmptyParams)
}

func (api *API) GetHardforkVersion() (string, error) {
	var resp string
	if err := api.caller.Call("get_hardfork_version", call.EmptyParams, &resp); err != nil {
		return "", err
	}
	return resp, nil
}

func (api *API) GetNextScheduledHardforkRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_next_scheduled_hardfork", call.EmptyParams)
}

/*
   // Keys
   (get_key_references)
*/

// XXX: Not sure about params.
//func (api *API) GetKeyReferencesRaw(key []string) (*json.RawMessage, error) {
//	return call.Raw(api.caller, "get_key_references", [][]string{key})
//}

/*
   // Accounts
   (get_accounts)
   (get_account_references)
   (lookup_account_names)
   (lookup_accounts)
   (get_account_count)
   (get_conversion_requests)
   (get_account_history)
*/

func (api *API) GetAccountsRaw(accountNames []string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_accounts", [][]string{accountNames})
}

// XXX: Not sure about params.
//func (api *API) GetAccountReferenceRaw(id string) (*json.RawMessage, error) {
//	return call.Raw(api.caller, "get_account_reference", []string{id})
//}

func (api *API) LookupAccountNamesRaw(accountNames []string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "lookup_account_names", [][]string{accountNames})
}

func (api *API) LookupAccountsRaw(lowerBoundName string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "lookup_accounts", []interface{}{lowerBoundName, limit})
}

func (api *API) GetAccountCountRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_account_count", call.EmptyParams)
}

func (api *API) GetConversionRequestsRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_conversion_requests", []string{accountName})
}

func (api *API) GetAccountHistoryRaw(account string, from uint64, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_account_history", []interface{}{account, from, limit})
}

/*
   // Market
   (get_order_book)
*/

func (api *API) GetOrderBookRaw(limit uint32) (*json.RawMessage, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	return call.Raw(api.caller, "get_order_book", []interface{}{limit})
}

/*
   // Authority / validation
   (get_transaction_hex)
   (get_transaction)
   (get_required_signatures)
   (get_potential_signatures)
   (verify_authority)
   (verify_account_authority)
*/

/*
   // Votes
   (get_active_votes)
   (get_account_votes)
*/

func (api *API) GetActiveVotesRaw(author, permlink string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_active_votes", []string{author, permlink})
}

func (api *API) GetActiveVotes(author, permlink string) ([]*VoteState, error) {
	var resp []*VoteState
	if err := api.caller.Call("get_active_votes", []string{author, permlink}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) GetAccountVotesRaw(voter string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_account_votes", []string{voter})
}

/*
   // Content
   (get_content)
   (get_content_replies)
   (get_discussions_by_author_before_date) - MISSING
   (get_replies_by_last_update)
*/

func (api *API) GetContentRaw(author, permlink string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_content", []string{author, permlink})
}

func (api *API) GetContent(author, permlink string) (*Content, error) {
	var resp Content
	if err := api.caller.Call("get_content", []string{author, permlink}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (api *API) GetContentRepliesRaw(parentAuthor, parentPermlink string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_content_replies", []string{parentAuthor, parentPermlink})
}

func (api *API) GetContentReplies(parentAuthor, parentPermlink string) ([]*Content, error) {
	var resp []*Content
	err := api.caller.Call("get_content_replies", []string{parentAuthor, parentPermlink}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) GetRepliesByLastUpdateRaw(
	startAuthor string,
	startPermlink string,
	limit uint32,
) (*json.RawMessage, error) {

	return call.Raw(
		api.caller, "get_replies_by_last_update", []interface{}{startAuthor, startPermlink, limit})
}

/*
   // Witnesses
   (get_witnesses)
   (get_witness_by_account)
   (get_witnesses_by_vote)
   (lookup_witness_accounts)
   (get_witness_count)
   (get_active_witnesses)
   (get_miner_queue)
*/
