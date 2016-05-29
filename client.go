package rpc

import (
	"encoding/json"
	"errors"
	"net/url"
)

var emptyParams = []string{}

type Client struct {
	t Transport
}

func Dial(address string) (*Client, error) {
	// Parse the address URL.
	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	// Look for the constructor associated with the given URL scheme.
	constructor, ok := registeredTransportConstructors[u.Scheme]
	if !ok {
		return nil, errors.New("no transport registered for URL scheme: " + u.Scheme)
	}

	// Use the constructor to get a Transport.
	transport, err := constructor(address)
	if err != nil {
		return nil, err
	}

	// Return the new Client, at last.
	return &Client{transport}, nil
}

func (client *Client) Close() error {
	return client.t.Close()
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

func (client *Client) GetTrendingTagsRaw(afterTag string, limit uint32) (*json.RawMessage, error) {
	return client.callRaw("get_trending_tags", []interface{}{afterTag, limit})
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

func (client *Client) GetDiscussionsByTrendingRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return client.callRaw("get_discussions_by_trending", query)
}

func (client *Client) GetDiscussionsByCreatedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return client.callRaw("get_discussions_by_created", query)
}

func (client *Client) GetDiscussionsByActiveRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return client.callRaw("get_discussions_by_active", query)
}

func (client *Client) GetDiscussionsByCashoutRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return client.callRaw("get_discussions_by_cashout", query)
}

func (client *Client) GetDiscussionsByPayoutRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return client.callRaw("get_discussions_by_payout", query)
}

func (client *Client) GetDiscussionsByVotesRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return client.callRaw("get_discussions_by_votes", query)
}

func (client *Client) GetDiscussionsByChildrenRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return client.callRaw("get_discussions_by_children", query)
}

func (client *Client) GetDiscussionsByHotRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return client.callRaw("get_discussions_by_hot", query)
}

func (client *Client) GetRecommendedForRaw(user string, limit uint32) (*json.RawMessage, error) {
	return client.callRaw("get_discussions_by_votes", []interface{}{user, limit})
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

func (client *Client) GetBlockHeaderRaw(blockNum uint32) (*json.RawMessage, error) {
	return client.callRaw("get_block_header", []uint32{blockNum})
}

func (client *Client) GetBlockRaw(blockNum uint32) (*json.RawMessage, error) {
	return client.callRaw("get_block", []uint32{blockNum})
}

func (client *Client) GetBlock(blockNum uint32) (*Block, error) {
	var resp Block
	if err := client.t.Call("get_block", []uint32{blockNum}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *Client) GetStateRaw(path string) (*json.RawMessage, error) {
	return client.callRaw("get_state", []string{path})
}

func (client *Client) GetTrendingCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return client.callRaw("get_trending_categories", []interface{}{after, limit})
}

func (client *Client) GetBestCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return client.callRaw("get_best_categories", []interface{}{after, limit})
}

func (client *Client) GetActiveCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return client.callRaw("get_active_categories", []interface{}{after, limit})
}

func (client *Client) GetRecentCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return client.callRaw("get_recent_categories", []interface{}{after, limit})
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

func (client *Client) GetConfigRaw() (*json.RawMessage, error) {
	return client.callRaw("get_config", emptyParams)
}

func (client *Client) GetConfig() (*Config, error) {
	var resp Config
	if err := client.t.Call("get_config", emptyParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *Client) GetDynamicGlobalPropertiesRaw() (*json.RawMessage, error) {
	return client.callRaw("get_dynamic_global_properties", emptyParams)
}

func (client *Client) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	var resp DynamicGlobalProperties
	if err := client.t.Call("get_dynamic_global_properties", emptyParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *Client) GetChainPropertiesRaw() (*json.RawMessage, error) {
	return client.callRaw("get_chain_properties", emptyParams)
}

func (client *Client) GetFeedHistoryRaw() (*json.RawMessage, error) {
	return client.callRaw("get_feed_history", emptyParams)
}

func (client *Client) GetCurrentMedianHistoryPriceRaw() (*json.RawMessage, error) {
	return client.callRaw("get_current_median_history_price", emptyParams)
}

func (client *Client) GetWitnessScheduleRaw() (*json.RawMessage, error) {
	return client.callRaw("get_witness_schedule", emptyParams)
}

func (client *Client) GetHardforkVersionRaw() (*json.RawMessage, error) {
	return client.callRaw("get_hardfork_version", emptyParams)
}

func (client *Client) GetNextScheduledHardforkRaw() (*json.RawMessage, error) {
	return client.callRaw("get_next_scheduled_hardfork", emptyParams)
}

/*
   // Keys
   (get_key_references)
*/

// XXX: Not sure about params.
func (client *Client) GetKeyReferencesRaw(key []string) (*json.RawMessage, error) {
	return client.callRaw("get_key_references", [][]string{key})
}

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

func (client *Client) GetAccountsRaw(accountNames []string) (*json.RawMessage, error) {
	return client.callRaw("get_accounts", [][]string{accountNames})
}

// XXX: Not sure about params.
func (client *Client) GetAccountReferenceRaw(id string) (*json.RawMessage, error) {
	return client.callRaw("get_account_reference", []string{id})
}

func (client *Client) LookupAccountNamesRaw(accountNames []string) (*json.RawMessage, error) {
	return client.callRaw("lookup_account_names", [][]string{accountNames})
}

func (client *Client) LookupAccountsRaw(lowerBoundName string, limit uint32) (*json.RawMessage, error) {
	return client.callRaw("lookup_accounts", []interface{}{lowerBoundName, limit})
}

func (client *Client) GetAccountCountRaw() (*json.RawMessage, error) {
	return client.callRaw("get_account_count", emptyParams)
}

func (client *Client) GetConversionRequestsRaw(accountName string) (*json.RawMessage, error) {
	return client.callRaw("get_conversion_requests", []string{accountName})
}

func (client *Client) GetAccountHistoryRaw(account string, from uint64, limit uint32) (*json.RawMessage, error) {
	return client.callRaw("get_account_history", []interface{}{account, from, limit})
}

/*
   // Market
   (get_order_book)
*/

func (client *Client) GetOrderBookRaw(limit uint32) (*json.RawMessage, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	return client.callRaw("get_order_book", []interface{}{limit})
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

func (client *Client) GetActiveVotesRaw(author, permlink string) (*json.RawMessage, error) {
	return client.callRaw("get_active_votes", []string{author, permlink})
}

func (client *Client) GetAccountVotesRaw(voter string) (*json.RawMessage, error) {
	return client.callRaw("get_account_votes", []string{voter})
}

/*
   // Content
   (get_content)
   (get_content_replies)
   (get_discussions_by_author_before_date) - MISSING
   (get_replies_by_last_update)
*/

func (client *Client) GetContentRaw(author, permlink string) (*json.RawMessage, error) {
	return client.callRaw("get_content", []string{author, permlink})
}

func (client *Client) GetContent(author, permlink string) (*Content, error) {
	var resp Content
	if err := client.t.Call("get_content", []string{author, permlink}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *Client) GetContentRepliesRaw(parent, parentPermlink string) (*json.RawMessage, error) {
	return client.callRaw("get_content_replies", []string{parent, parentPermlink})
}

func (client *Client) GetRepliesByLastUpdateRaw(
	startAuthor string,
	startPermlink string,
	limit uint32,
) (*json.RawMessage, error) {

	return client.callRaw("get_replies_by_last_update", []interface{}{startAuthor, startPermlink, limit})
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

/*
 * Helpers
 */

func (client *Client) callRaw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := client.t.Call(method, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
