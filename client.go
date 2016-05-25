package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/tchap/rpc-codec/jsonrpc2"
	"golang.org/x/net/websocket"
)

var emptyParams = []string{}

type Client struct {
	rpc *jsonrpc2.Client
}

func Dial(addr string) (*Client, error) {
	// Connect to the given WebSocket URL.
	conn, err := websocket.Dial(addr, "", "http://localhost")
	if err != nil {
		return nil, err
	}

	// Instantiate a JSON-RPC client.
	client := jsonrpc2.NewClient(conn)

	// Return a new Client instance.
	return &Client{client}, nil
}

func (client *Client) Close() error {
	return client.rpc.Close()
}

/*
   // Subscriptions
   (set_subscribe_callback)
   (set_pending_transaction_callback)
   (set_block_applied_callback)
   (cancel_all_subscriptions)
*/

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

func (client *Client) GetRawBlockRaw(blockNum uint32) (*json.RawMessage, error) {
	return client.callRaw("get_block", []uint32{blockNum})
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
*/

func (client *Client) GetConfigRaw() (*json.RawMessage, error) {
	return client.callRaw("get_config", emptyParams)
}

func (client *Client) GetDynamicGlobalPropertiesRaw() (*json.RawMessage, error) {
	return client.callRaw("get_dynamic_global_properties", emptyParams)
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

/*
   // Keys
   (get_key_references)
*/

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

/*
   // Market
   (get_order_book)
*/

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

/*
   // Content
   (get_content)
   (get_content_replies)
   (get_discussions_by_total_pending_payout)
   (get_discussions_in_category_by_total_pending_payout)
   (get_discussions_by_last_update)
   (get_discussions_by_last_active)
   (get_discussions_by_votes)
   (get_discussions_by_created)
   (get_discussions_in_category_by_last_update)
   (get_discussions_in_category_by_last_active)
   (get_discussions_in_category_by_votes)
   (get_discussions_in_category_by_created)
   (get_discussions_by_author_before_date)
   (get_discussions_by_cashout_time)
   (get_discussions_in_category_by_cashout_time)
*/

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

func (client *Client) callRaw(
	method string,
	params []interface{},
) (*json.RawMessage, error) {

	var resp json.RawMessage
	if err := client.rpc.Call(method, params, &resp); err != nil {
		return nil, err
	}
	return &resp.nil
}
