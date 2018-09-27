package database

import (
	"encoding/json"
	"fmt"

	"github.com/asuleymanov/steem-go/transports"
	"github.com/asuleymanov/steem-go/types"
)

const apiID = "database_api"

//API plug-in structure
type API struct {
	caller transports.Caller
}

//NewAPI plug-in initialization
func NewAPI(caller transports.Caller) *API {
	return &API{caller}
}

var emptyParams = []struct{}{}

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{apiID, method, params}, resp)
}

//GetTrendingTags api request get_trending_tags
func (api *API) GetTrendingTags(afterTag string, limit uint32) ([]*TrendingTags, error) {
	var resp []*TrendingTags
	err := api.call("get_trending_tags", []interface{}{afterTag, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetTagsUsedByAuthor api request get_tags_used_by_author
func (api *API) GetTagsUsedByAuthor(accountName string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_tags_used_by_author", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetPostDiscussionsByPayout api request get_post_discussions_by_payout
func (api *API) GetPostDiscussionsByPayout(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_post_discussions_by_payout", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetCommentDiscussionsByPayout api request get_comment_discussions_by_payout
func (api *API) GetCommentDiscussionsByPayout(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_comment_discussions_by_payout", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByTrending api request get_discussions_by_trending
func (api *API) GetDiscussionsByTrending(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_trending", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByTrending30 api request get_discussions_by_trending30
func (api *API) GetDiscussionsByTrending30(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_trending30", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByCreated api request get_discussions_by_created
func (api *API) GetDiscussionsByCreated(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_created", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByActive api request get_discussions_by_active
func (api *API) GetDiscussionsByActive(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_active", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByCashout api request get_discussions_by_cashout
func (api *API) GetDiscussionsByCashout(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_cashout", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByPayout api request get_discussions_by_payout
func (api *API) GetDiscussionsByPayout(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_payout", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByVotes api request get_discussions_by_votes
func (api *API) GetDiscussionsByVotes(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_votes", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByChildren api request get_discussions_by_children
func (api *API) GetDiscussionsByChildren(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_children", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByHot api request get_discussions_by_hot
func (api *API) GetDiscussionsByHot(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_hot", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByFeed api request get_discussions_by_feed
func (api *API) GetDiscussionsByFeed(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_feed", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByBlog api request get_discussions_by_blog
func (api *API) GetDiscussionsByBlog(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_blog", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByComments api request get_discussions_by_comments
func (api *API) GetDiscussionsByComments(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_comments", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByPromoted api request get_discussions_by_promoted
func (api *API) GetDiscussionsByPromoted(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_promoted", query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetBlockHeader api request get_block_header
func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	var resp BlockHeader
	err := api.call("get_block_header", []uint32{blockNum}, &resp)
	if err != nil {
		return nil, err
	}
	resp.Number = blockNum
	return &resp, nil
}

//GetBlock api request get_block
func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	var resp Block
	err := api.call("get_block", []uint32{blockNum}, &resp)
	if err != nil {
		return nil, err
	}
	resp.Number = blockNum
	return &resp, nil
}

//GetOpsInBlock api request get_ops_in_block
func (api *API) GetOpsInBlock(blockNum uint32, onlyVirtual bool) ([]*types.OperationObject, error) {
	var resp []*types.OperationObject
	err := api.call("get_ops_in_block", []interface{}{blockNum, onlyVirtual}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetState api request get_state
func (api *API) GetState(path string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_state", []string{path}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetTrendingCategories api request get_trending_categories
func (api *API) GetTrendingCategories(after string, limit uint32) ([]*Categories, error) {
	var resp []*Categories
	err := api.call("get_trending_categories", []interface{}{after, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetBestCategories api request get_best_categories
func (api *API) GetBestCategories(after string, limit uint32) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_best_categories", []interface{}{after, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetActiveCategories api request get_active_categories
func (api *API) GetActiveCategories(after string, limit uint32) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_active_categories", []interface{}{after, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetRecentCategories api request get_recent_categories
func (api *API) GetRecentCategories(after string, limit uint32) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_recent_categories", []interface{}{after, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetConfig api request get_config
func (api *API) GetConfig() (*Config, error) {
	var resp Config
	err := api.call("get_config", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetDynamicGlobalProperties api request get_dynamic_global_properties
func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	var resp DynamicGlobalProperties
	err := api.call("get_dynamic_global_properties", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetChainProperties api request get_chain_properties
func (api *API) GetChainProperties() (*types.ChainProperties, error) {
	var resp types.ChainProperties
	err := api.call("get_chain_properties", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetFeedHistory api request get_feed_history
func (api *API) GetFeedHistory() (*FeedHistory, error) {
	var resp FeedHistory
	err := api.call("get_feed_history", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetCurrentMedianHistoryPrice api request get_current_median_history_price
func (api *API) GetCurrentMedianHistoryPrice() (*CurrentMedianHistoryPrice, error) {
	var resp CurrentMedianHistoryPrice
	err := api.call("get_current_median_history_price", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetWitnessSchedule api request get_witness_schedule
func (api *API) GetWitnessSchedule() (*WitnessSchedule, error) {
	var resp WitnessSchedule
	err := api.call("get_witness_schedule", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetHardforkVersion api request get_hardfork_version
func (api *API) GetHardforkVersion() (*string, error) {
	var resp string
	err := api.call("get_hardfork_version", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetNextScheduledHardfork api request get_next_scheduled_hardfork
func (api *API) GetNextScheduledHardfork() (*NextScheduledHardfork, error) {
	var resp NextScheduledHardfork
	err := api.call("get_next_scheduled_hardfork", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetKeyReferences api request get_key_references
func (api *API) GetKeyReferences(pubkey string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_key_references", []interface{}{pubkey}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetAccounts api request get_accounts
func (api *API) GetAccounts(accountNames []string) ([]*Account, error) {
	var resp []*Account
	err := api.call("get_accounts", [][]string{accountNames}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetAccountReferences api request get_account_references
func (api *API) GetAccountReferences(accountID uint32) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_account_references", []interface{}{accountID}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//LookupAccountNames api request lookup_account_names
func (api *API) LookupAccountNames(accountNames []string) ([]*Account, error) {
	var resp []*Account
	err := api.call("lookup_account_names", [][]string{accountNames}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//LookupAccounts api request lookup_accounts
func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]*string, error) {
	var resp []*string
	err := api.call("lookup_accounts", []interface{}{lowerBoundName, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetAccountCount api request get_account_count
func (api *API) GetAccountCount() (*uint32, error) {
	var resp uint32
	err := api.call("get_account_count", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetConversionRequests api request get_conversion_requests
func (api *API) GetConversionRequests(accountName string) ([]*ConversionRequests, error) {
	var resp []*ConversionRequests
	err := api.call("get_conversion_requests", []string{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetAccountHistory api request get_account_history
func (api *API) GetAccountHistory(account string, from int64, limit uint32) ([]*types.OperationObject, error) {
	if limit > 10000 {
		return nil, fmt.Errorf("%v: get_account_history -> limit must not exceed 10000", apiID)
	}
	if from == 0 {
		return nil, fmt.Errorf("%v: get_account_history -> from can not have the value 0", apiID)
	}
	if from < int64(limit) && !(from < 0) {
		return nil, fmt.Errorf("%v: get_account_history -> from must be greater than or equal to the limit", apiID)
	}
	var raw json.RawMessage
	err := api.call("get_account_history", []interface{}{account, from, limit}, &raw)
	if err != nil {
		return nil, err
	}
	var tmp1 [][]interface{}
	if err := json.Unmarshal([]byte(raw), &tmp1); err != nil {
		return nil, err
	}
	var resp []*types.OperationObject
	for _, v := range tmp1 {
		byteData, errm := json.Marshal(&v[1])
		if errm != nil {
			return nil, errm
		}
		var tmp *types.OperationObject
		if err := json.Unmarshal(byteData, &tmp); err != nil {
			return nil, err
		}
		resp = append(resp, tmp)
	}
	return resp, nil
}

//GetOwnerHistory api request get_owner_history
func (api *API) GetOwnerHistory(accountName string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_owner_history", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetRecoveryRequest api request get_recovery_request
func (api *API) GetRecoveryRequest(accountName string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_recovery_request", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetEscrow api request get_escrow
func (api *API) GetEscrow(from string, escrow_id uint32) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_escrow", []interface{}{from, escrow_id}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetWithdrawRoutes api request get_withdraw_routes
func (api *API) GetWithdrawRoutes(accountName string, withdraw_route_type string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_withdraw_routes", []interface{}{accountName, withdraw_route_type}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetAccountBandwidth api request get_account_bandwidth
func (api *API) GetAccountBandwidth(accountName string, bandwidth_type uint32) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_account_bandwidth", []interface{}{accountName, bandwidth_type}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetSavingsWithdrawFrom api request get_savings_withdraw_from
func (api *API) GetSavingsWithdrawFrom(accountName string) ([]*SavingsWithdraw, error) {
	var resp []*SavingsWithdraw
	err := api.call("get_savings_withdraw_from", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetSavingsWithdrawTo api request get_savings_withdraw_to
func (api *API) GetSavingsWithdrawTo(accountName string) ([]*SavingsWithdraw, error) {
	var resp []*SavingsWithdraw
	err := api.call("get_savings_withdraw_to", []interface{}{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetOrderBook api request get_order_book
func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("%v: get_order_book -> limit must not exceed 1000", apiID)
	}
	var resp OrderBook
	err := api.call("get_order_book", []interface{}{limit}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetOpenOrders api request get_open_orders
func (api *API) GetOpenOrders(accountName string) ([]*OpenOrders, error) {
	var resp []*OpenOrders
	err := api.call("get_open_orders", []string{accountName}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetLiquidityQueue api request get_liquidity_queue
func (api *API) GetLiquidityQueue(startAccount string, limit uint32) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_liquidity_queue", []interface{}{startAccount, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetTransactionHex api request get_transaction_hex
func (api *API) GetTransactionHex(trx *types.Transaction) (*string, error) {
	var resp string
	err := api.call("get_transaction_hex", []interface{}{&trx}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetTransaction api request get_transaction
func (api *API) GetTransaction(id string) (*types.Transaction, error) {
	var resp types.Transaction
	err := api.call("get_transaction", []string{id}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetRequiredSignatures api request get_required_signatures
func (api *API) GetRequiredSignatures(trx *types.Transaction, keys []string) ([]*string, error) {
	var resp []*string
	err := api.call("get_required_signatures", []interface{}{trx, keys}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetPotentialSignatures api request get_potential_signatures
func (api *API) GetPotentialSignatures(trx *types.Transaction) ([]*string, error) {
	var resp []*string
	err := api.call("get_potential_signatures", []interface{}{&trx}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetVerifyAuthority api request verify_authority
func (api *API) GetVerifyAuthority(trx *types.Transaction) (*bool, error) {
	var resp bool
	err := api.call("verify_authority", []interface{}{&trx}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetVerifyAccountAuthority api request verify_account_authority
func (api *API) GetVerifyAccountAuthority(accountName string, keys []string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("verify_account_authority", []interface{}{accountName, keys}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetActiveVotes api request get_active_votes
func (api *API) GetActiveVotes(author, permlink string) ([]*VoteState, error) {
	var resp []*VoteState
	err := api.call("get_active_votes", []interface{}{author, permlink}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetAccountVotes api request get_account_votes
func (api *API) GetAccountVotes(author string) ([]*Votes, error) {
	var resp []*Votes
	err := api.call("get_account_votes", []interface{}{author}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetContent api request get_content
func (api *API) GetContent(author, permlink string) (*Content, error) {
	var resp Content
	err := api.call("get_content", []interface{}{author, permlink}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetContentReplies api request get_content_replies
func (api *API) GetContentReplies(parentAuthor, parentPermlink string, opts ...interface{}) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_content_replies", []interface{}{parentAuthor, parentPermlink}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetDiscussionsByAuthorBeforeDate api request get_discussions_by_author_before_date
func (api *API) GetDiscussionsByAuthorBeforeDate(author, permlink, date string, limit uint32) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_discussions_by_author_before_date", []interface{}{author, permlink, date, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetRepliesByLastUpdate api request get_replies_by_last_update
func (api *API) GetRepliesByLastUpdate(startAuthor, startPermlink string, limit uint32) ([]*Content, error) {
	var resp []*Content
	err := api.call("get_replies_by_last_update", []interface{}{startAuthor, startPermlink, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetWitnesses api request get_witnesses
func (api *API) GetWitnesses(id []uint32) ([]*Witness, error) {
	var resp []*Witness
	err := api.call("get_witnesses", [][]uint32{id}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetWitnessByAccount api request get_witness_by_account
func (api *API) GetWitnessByAccount(author string) (*Witness, error) {
	var resp Witness
	err := api.call("get_witness_by_account", []string{author}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetWitnessByVote api request get_witnesses_by_vote
func (api *API) GetWitnessByVote(author string, limit uint) ([]*Witness, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("%v: get_witnesses_by_vote -> limit must not exceed 1000", apiID)
	}
	var resp []*Witness
	err := api.call("get_witnesses_by_vote", []interface{}{author, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//LookupWitnessAccounts api request lookup_witness_accounts
func (api *API) LookupWitnessAccounts(author string, limit uint) ([]*string, error) {
	if limit > 1000 {
		return nil, fmt.Errorf("%v: lookup_witness_accounts -> limit must not exceed 1000", apiID)
	}
	var resp []*string
	err := api.call("lookup_witness_accounts", []interface{}{author, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetWitnessCount api request get_witness_count
func (api *API) GetWitnessCount() (*uint32, error) {
	var resp uint32
	err := api.call("get_witness_count", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetActiveWitnesses api request get_active_witnesses
func (api *API) GetActiveWitnesses() ([]*string, error) {
	var resp []*string
	err := api.call("get_active_witnesses", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetMinerQueue api request get_miner_queue
func (api *API) GetMinerQueue() ([]*string, error) {
	var resp []*string
	err := api.call("get_miner_queue", emptyParams, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//GetRewardFund api request get_reward_fund
func (api *API) GetRewardFund(name string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_reward_fund", []interface{}{name}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GetVestingDelegations api request get_vesting_delegations
func (api *API) GetVestingDelegations(account, from string, limit uint32) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("get_vesting_delegations", []interface{}{account, from, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
