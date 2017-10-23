package database

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/golos-go/interfaces"
	"github.com/asuleymanov/golos-go/types"

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

var EmptyParams = []string{}

func (api *API) Raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call(method, params, &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to call %v\n", APIID, method)
	}
	return &resp, nil
}

//set_subscribe_callback                 | *NONE* | *NONE* |

//set_pending_transaction_callback       | *NONE* | *NONE* |

//set_block_applied_callback             | *NONE* | *NONE* |

//cancel_all_subscriptions               | *NONE* | *NONE* |

//get_trending_tags
func (api *API) GetTrendingTags(afterTag string, limit uint32) ([]*TrendingTags, error) {
	raw, err := api.Raw("get_trending_tags", []interface{}{afterTag, limit})
	if err != nil {
		return nil, err
	}
	var resp []*TrendingTags
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_trending_tags response", APIID)
	}
	return resp, nil
}

//get_tags_used_by_author
func (api *API) GetTagsUsedByAuthor(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_tags_used_by_author", []interface{}{accountName})
}

//get_discussions_by_trending
func (api *API) GetDiscussionsByTrending(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_trending", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_trending response", APIID)
	}
	return resp, nil
}

//get_discussions_by_trending30
func (api *API) GetDiscussionsByTrending30(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_trending30", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_trending30 response", APIID)
	}
	return resp, nil
}

//get_discussions_by_created
func (api *API) GetDiscussionsByCreated(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_created", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_created response", APIID)
	}
	return resp, nil
}

//get_discussions_by_active
func (api *API) GetDiscussionsByActive(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_active", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_active response", APIID)
	}
	return resp, nil
}

//get_discussions_by_cashout
func (api *API) GetDiscussionsByCashout(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_cashout", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_cashout response", APIID)
	}
	return resp, nil
}

//get_discussions_by_payout
func (api *API) GetDiscussionsByPayout(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_payout", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_payout response", APIID)
	}
	return resp, nil
}

//get_discussions_by_votes
func (api *API) GetDiscussionsByVotes(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_votes", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_votes response", APIID)
	}
	return resp, nil
}

//get_discussions_by_children
func (api *API) GetDiscussionsByChildren(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_children", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_children response", APIID)
	}
	return resp, nil
}

//get_discussions_by_hot
func (api *API) GetDiscussionsByHot(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_hot", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_hot response", APIID)
	}
	return resp, nil
}

//get_discussions_by_feed
func (api *API) GetDiscussionsByFeed(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_feed", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_feed response", APIID)
	}
	return resp, nil
}

//get_discussions_by_blog
func (api *API) GetDiscussionsByBlog(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_blog", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_blog response", APIID)
	}
	return resp, nil
}

//get_discussions_by_comments
func (api *API) GetDiscussionsByComments(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_comments", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_comments response", APIID)
	}
	return resp, nil
}

//get_discussions_by_promoted
func (api *API) GetDiscussionsByPromoted(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_promoted", query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_promoted response", APIID)
	}
	return resp, nil
}

//get_block_header
func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	raw, err := api.Raw("get_block_header", []uint32{blockNum})
	if err != nil {
		return nil, err
	}
	var resp BlockHeader
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_block_header response", APIID)
	}
	resp.Number = blockNum
	return &resp, nil
}

//get_block
func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	raw, err := api.Raw("get_block", []uint32{blockNum})
	if err != nil {
		return nil, err
	}
	var resp Block
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_block response", APIID)
	}
	resp.Number = blockNum
	return &resp, nil
}

//get_ops_in_block
func (api *API) GetOpsInBlock(blockNum uint32, only_virtual bool) ([]*types.OperationObject, error) {
	raw, err := api.Raw("get_ops_in_block", []interface{}{blockNum, only_virtual})
	if err != nil {
		return nil, err
	}
	var resp []*types.OperationObject
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_ops_in_block response", APIID)
	}
	return resp, nil
}

//get_state
func (api *API) GetState(path string) (*json.RawMessage, error) {
	return api.Raw("get_state", []string{path})
}

//get_trending_categories
func (api *API) GetTrendingCategories(after string, limit uint32) ([]*Categories, error) {
	raw, err := api.Raw("get_trending_categories", []interface{}{after, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Categories
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_trending_categories response", APIID)
	}
	return resp, nil
}

//get_best_categories
func (api *API) GetBestCategories(after string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_best_categories", []interface{}{after, limit})
}

//get_active_categories
func (api *API) GetActiveCategories(after string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_active_categories", []interface{}{after, limit})
}

//get_recent_categories
func (api *API) GetRecentCategories(after string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_recent_categories", []interface{}{after, limit})
}

//get_config
func (api *API) GetConfig() (*Config, error) {
	raw, err := api.Raw("get_config", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp Config
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_config response", APIID)
	}
	return &resp, nil
}

//get_dynamic_global_properties
func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	raw, err := api.Raw("get_dynamic_global_properties", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp DynamicGlobalProperties
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_dynamic_global_properties response", APIID)
	}
	return &resp, nil
}

//get_chain_properties
func (api *API) GetChainProperties() (*ChainProperties, error) {
	raw, err := api.Raw("get_chain_properties", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp ChainProperties
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_chain_properties response", APIID)
	}
	return &resp, nil
}

//get_feed_history
func (api *API) GetFeedHistory() (*FeedHistory, error) {
	raw, err := api.Raw("get_feed_history", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp FeedHistory
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_feed_history response", APIID)
	}
	return &resp, nil
}

//get_current_median_history_price
func (api *API) GetCurrentMedianHistoryPrice() (*CurrentMedianHistoryPrice, error) {
	raw, err := api.Raw("get_current_median_history_price", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp CurrentMedianHistoryPrice
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_current_median_history_price response", APIID)
	}
	return &resp, nil
}

//get_witness_schedule
func (api *API) GetWitnessSchedule() (*WitnessSchedule, error) {
	raw, err := api.Raw("get_witness_schedule", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp WitnessSchedule
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_witness_schedule response", APIID)
	}
	return &resp, nil
}

//get_hardfork_version
func (api *API) GetHardforkVersion() (string, error) {
	raw, err := api.Raw("get_hardfork_version", EmptyParams)
	if err != nil {
		return "", err
	}
	var resp string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return "", errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_hardfork_version response", APIID)
	}
	return resp, nil
}

//get_next_scheduled_hardfork
func (api *API) GetNextScheduledHardfork() (*NextScheduledHardfork, error) {
	raw, err := api.Raw("get_next_scheduled_hardfork", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp NextScheduledHardfork
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_next_scheduled_hardfork response", APIID)
	}
	return &resp, nil
}

//get_key_references
//Unfortunately to say what this command does is not possible. (Any call to it leads to an error).

//get_accounts
func (api *API) GetAccounts(accountNames []string) ([]*Account, error) {
	raw, err := api.Raw("get_accounts", [][]string{accountNames})
	if err != nil {
		return nil, err
	}
	var resp []*Account
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_accounts response", APIID)
	}
	return resp, nil
}

//get_account_references
//Unfortunately to say what this command does is not possible. (Any call to it leads to an error).

//lookup_account_names
func (api *API) LookupAccountNames(accountNames []string) (*json.RawMessage, error) {
	return api.Raw("lookup_account_names", [][]string{accountNames})
}

//lookup_accounts
func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]string, error) {
	raw, err := api.Raw("lookup_accounts", []interface{}{lowerBoundName, limit})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal lookup_accounts response", APIID)
	}
	return resp, nil
}

//get_account_count
func (api *API) GetAccountCount() (uint32, error) {
	raw, err := api.Raw("get_account_count", EmptyParams)
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_account_count response", APIID)
	}
	return resp, nil
}

//get_conversion_requests
func (api *API) GetConversionRequests(accountName string) ([]*ConversionRequests, error) {
	raw, err := api.Raw("get_conversion_requests", []string{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*ConversionRequests
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_conversion_requests response", APIID)
	}
	return resp, nil
}

//get_account_history
/*func (api *API) GetAccountHistory(account string, from uint64, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_account_history", []interface{}{account, from, limit})
}*/

func (api *API) GetAccountHistory(account string, from uint64, limit uint32) ([]*types.OperationObject, error) {
	raw, err := api.Raw("get_account_history", []interface{}{account, from, limit})
	if err != nil {
		return nil, err
	}
	var tmp1 [][]interface{}
	if err := json.Unmarshal([]byte(*raw), &tmp1); err != nil {
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

//get_owner_history
func (api *API) GetOwnerHistory(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_owner_history", []interface{}{accountName})
}

//get_recovery_request
func (api *API) GetRecoveryRequest(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_recovery_request", []interface{}{accountName})
}

//get_escrow
func (api *API) GetEscrow(from string, escrow_id uint32) (*json.RawMessage, error) {
	return api.Raw("get_escrow", []interface{}{from, escrow_id})
}

//get_withdraw_routes
func (api *API) GetWithdrawRoutes(accountName string, withdraw_route_type string) (*json.RawMessage, error) {
	return api.Raw("get_withdraw_routes", []interface{}{accountName, withdraw_route_type})
}

//get_account_bandwidth
func (api *API) GetAccountBandwidth(accountName string, bandwidth_type uint32) (*json.RawMessage, error) {
	return api.Raw("get_account_bandwidth", []interface{}{accountName, bandwidth_type})
}

//get_savings_withdraw_from
func (api *API) GetSavingsWithdrawFrom(accountName string) ([]*SavingsWithdraw, error) {
	raw, err := api.Raw("get_savings_withdraw_from", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*SavingsWithdraw
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_savings_withdraw_from response", APIID)
	}
	return resp, nil
}

//get_savings_withdraw_to
func (api *API) GetSavingsWithdrawTo(accountName string) ([]*SavingsWithdraw, error) {
	raw, err := api.Raw("get_savings_withdraw_to", []interface{}{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*SavingsWithdraw
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_savings_withdraw_to response", APIID)
	}
	return resp, nil
}

//get_order_book
func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	raw, err := api.Raw("get_order_book", []interface{}{limit})
	if err != nil {
		return nil, err
	}
	var resp *OrderBook
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_order_book response", APIID)
	}
	return resp, nil
}

//get_open_orders
func (api *API) GetOpenOrders(accountName string) ([]*OpenOrders, error) {
	raw, err := api.Raw("get_open_orders", []string{accountName})
	if err != nil {
		return nil, err
	}
	var resp []*OpenOrders
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_open_orders response", APIID)
	}
	return resp, nil
}

//get_liquidity_queue
func (api *API) GetLiquidityQueue(startAccount string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_liquidity_queue", []interface{}{startAccount, limit})
}

//get_transaction_hex
func (api *API) GetTransactionHex(trx *types.Transaction) (*json.RawMessage, error) {
	return api.Raw("get_transaction_hex", []interface{}{&trx})
}

//get_transaction
func (api *API) GetTransaction(id string) (*types.Transaction, error) {
	raw, err := api.Raw("get_transaction", []string{id})
	if err != nil {
		return nil, err
	}
	var resp types.Transaction
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_transaction response", APIID)
	}
	return &resp, nil
}

//get_required_signatures                | *NONE* | *NONE* |

//get_potential_signatures
func (api *API) GetPotentialSignatures(trx *types.Transaction) ([]string, error) {
	raw, err := api.Raw("get_potential_signatures", []interface{}{&trx})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_potential_signatures response", APIID)
	}
	return resp, nil
}

//verify_authority
func (api *API) GetVerifyAuthoruty(trx *types.Transaction) (bool, error) {
	raw, err := api.Raw("verify_authority", []interface{}{&trx})
	if err != nil {
		return false, err
	}
	var resp bool
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return false, errors.Wrapf(err, "golos-go: %v: failed to unmarshal verify_authority response", APIID)
	}
	return resp, nil
}

//verify_account_authority               | *NONE* | *NONE* |

//get_active_votes
func (api *API) GetActiveVotes(author, permlink string) ([]*VoteState, error) {
	raw, err := api.Raw("get_active_votes", []string{author, permlink})
	if err != nil {
		return nil, err
	}
	var resp []*VoteState
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_active_votes response", APIID)
	}
	return resp, nil
}

//get_account_votes
func (api *API) GetAccountVotes(author string) ([]*Votes, error) {
	raw, err := api.Raw("get_account_votes", []string{author})
	if err != nil {
		return nil, err
	}
	var resp []*Votes
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_account_votes response", APIID)
	}
	return resp, nil
}

//get_content
func (api *API) GetContent(author, permlink string) (*Content, error) {
	raw, err := api.Raw("get_content", []string{author, permlink})
	if err != nil {
		return nil, err
	}
	var resp Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_content response", APIID)
	}
	return &resp, nil
}

//get_content_replies
func (api *API) GetContentReplies(parentAuthor, parentPermlink string) ([]*Content, error) {
	raw, err := api.Raw("get_content_replies", []string{parentAuthor, parentPermlink})
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_content_replies response", APIID)
	}
	return resp, nil
}

//get_discussions_by_author_before_date
func (api *API) GetDiscussionsByAuthorBeforeDate(Author, Permlink, Date string, limit uint32) ([]*Content, error) {
	raw, err := api.Raw("get_discussions_by_author_before_date", []interface{}{Author, Permlink, Date, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_author_before_date response", APIID)
	}
	return resp, nil
}

//get_replies_by_last_update
func (api *API) GetRepliesByLastUpdate(startAuthor, startPermlink string, limit uint32) ([]*Content, error) {
	raw, err := api.Raw("get_replies_by_last_update", []interface{}{startAuthor, startPermlink, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_replies_by_last_update response", APIID)
	}
	return resp, nil
}

//get_witnesses
func (api *API) GetWitnesses(id []uint32) ([]*Witness, error) {
	raw, err := api.Raw("get_witnesses", [][]uint32{id})
	if err != nil {
		return nil, err
	}
	var resp []*Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_witnesses response", APIID)
	}
	return resp, nil
}

//get_witness_by_account
func (api *API) GetWitnessByAccount(author string) (*Witness, error) {
	raw, err := api.Raw("get_witness_by_account", []string{author})
	if err != nil {
		return nil, err
	}
	var resp Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_witness_by_account response", APIID)
	}
	return &resp, nil
}

//get_witnesses_by_vote
func (api *API) GetWitnessByVote(author string, limit uint) ([]*Witness, error) {
	if limit > 1000 {
		return nil, errors.New("GetWitnessByVote: limit must not exceed 1000")
	}
	raw, err := api.Raw("get_witnesses_by_vote", []interface{}{author, limit})
	if err != nil {
		return nil, err
	}
	var resp []*Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_witnesses_by_vote response", APIID)
	}
	return resp, nil
}

//lookup_witness_accounts
func (api *API) LookupWitnessAccounts(author string, limit uint) ([]string, error) {
	if limit > 1000 {
		return nil, errors.New("LookupWitnessAccounts: limit must not exceed 1000")
	}
	raw, err := api.Raw("lookup_witness_accounts", []interface{}{author, limit})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal lookup_witness_accounts response", APIID)
	}
	return resp, nil
}

//get_witness_count
func (api *API) GetWitnessCount() (uint32, error) {
	raw, err := api.Raw("get_witness_count", EmptyParams)
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_witness_count response", APIID)
	}
	return resp, nil
}

//get_active_witnesses
func (api *API) GetActiveWitnesses() ([]string, error) {
	raw, err := api.Raw("get_active_witnesses", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_active_witnesses response", APIID)
	}
	return resp, nil
}

//get_miner_queue
func (api *API) GetMinerQueue() ([]string, error) {
	raw, err := api.Raw("get_miner_queue", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_miner_queue response", APIID)
	}
	return resp, nil
}
