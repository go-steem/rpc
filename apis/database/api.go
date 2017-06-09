package database

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/golos-go/interfaces"
	"github.com/asuleymanov/golos-go/internal/call"
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

//set_subscribe_callback                 | *NONE* | *NONE* |

//set_pending_transaction_callback       | *NONE* | *NONE* |

//set_block_applied_callback             | *NONE* | *NONE* |

//cancel_all_subscriptions               | *NONE* | *NONE* |

//get_trending_tags                      | **DONE** | **DONE** |

func (api *API) GetTrendingTagsRaw(afterTag string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_trending_tags", []interface{}{afterTag, limit})
}

func (api *API) GetTrendingTags(afterTag string, limit uint32) ([]*TrendingTags, error) {
	var resp []*TrendingTags
	if err := api.caller.Call("get_trending_tags", []interface{}{afterTag, limit}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_tags_used_by_author                | **DONE** | *NONE* |

func (api *API) GetTagsUsedByAuthorRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_tags_used_by_author", []interface{}{accountName})
}

//get_discussions_by_trending            | **DONE** | **DONE** |

func (api *API) GetDiscussionsByTrendingRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_trending", query)
}

func (api *API) GetDiscussionsByTrending(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_trending", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_trending30          | **DONE** | **DONE** |

func (api *API) GetDiscussionsByTrending30Raw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_trending30", query)
}

func (api *API) GetDiscussionsByTrending30(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_trending30", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_created             | **DONE** | **DONE** |

func (api *API) GetDiscussionsByCreatedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_created", query)
}

func (api *API) GetDiscussionsByCreated(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_created", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_active              | **DONE** | **DONE** |

func (api *API) GetDiscussionsByActiveRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_active", query)
}

func (api *API) GetDiscussionsByActive(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_active", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_cashout             | **DONE** | **DONE** |

func (api *API) GetDiscussionsByCashoutRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_cashout", query)
}

func (api *API) GetDiscussionsByCashout(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_cashout", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_payout              | **DONE** | **DONE** |

func (api *API) GetDiscussionsByPayoutRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_payout", query)
}

func (api *API) GetDiscussionsByPayout(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_payout", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_votes               | **DONE** | **DONE** |

func (api *API) GetDiscussionsByVotesRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_votes", query)
}

func (api *API) GetDiscussionsByVotes(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_votes", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_children            | **DONE** | **DONE** |

func (api *API) GetDiscussionsByChildrenRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_children", query)
}

func (api *API) GetDiscussionsByChildren(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_children", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_hot                 | **DONE** | **DONE** |

func (api *API) GetDiscussionsByHotRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_hot", query)
}

func (api *API) GetDiscussionsByHot(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_hot", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_feed                | **DONE** | **DONE** |

func (api *API) GetDiscussionsByFeedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_feed", query)
}

func (api *API) GetDiscussionsByFeed(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_feed", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_blog                | **DONE** | **DONE** |

func (api *API) GetDiscussionsByBlogRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_blog", query)
}

func (api *API) GetDiscussionsByBlog(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_blog", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_comments            | **DONE** | **DONE** |

func (api *API) GetDiscussionsByCommentsRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_comments", query)
}

func (api *API) GetDiscussionsByComments(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_comments", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_discussions_by_promoted            | **DONE** | **DONE** |

func (api *API) GetDiscussionsByPromotedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_promoted", query)
}

func (api *API) GetDiscussionsByPromoted(query *DiscussionQuery) ([]*Content, error) {
	var resp []*Content
	if err := api.caller.Call("get_discussions_by_promoted", query, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_block_header                       | **DONE** | **DONE** |

func (api *API) GetBlockHeaderRaw(blockNum uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_block_header", []uint32{blockNum})
}

func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	var resp BlockHeader
	if err := api.caller.Call("get_block_header", []uint32{blockNum}, &resp); err != nil {
		return nil, err
	}
	resp.Number = blockNum
	return &resp, nil
}

//get_block                              | **DONE** | ***PARTIALLY DONE*** |

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

//get_ops_in_block                       | **DONE** | ***PARTIALLY DONE*** |

func (api *API) GetOpsInBlockRaw(blockNum uint32, only_virtual bool) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_ops_in_block", []interface{}{blockNum, only_virtual})
}

func (api *API) GetOpsInBlock(blockNum uint32, only_virtual bool) ([]*OpsInBlock, error) {
	var resp []*OpsInBlock
	if err := api.caller.Call("get_ops_in_block", []interface{}{blockNum, only_virtual}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_state                              | **DONE** | *NONE* |

func (api *API) GetStateRaw(path string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_state", []string{path})
}

//get_trending_categories                | **DONE** | **DONE** |

func (api *API) GetTrendingCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_trending_categories", []interface{}{after, limit})
}

func (api *API) GetTrendingCategories(after string, limit uint32) ([]*Categories, error) {
	var resp []*Categories
	if err := api.caller.Call("get_trending_categories", []interface{}{after, limit}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_best_categories                    | **DONE** | *NONE* |

func (api *API) GetBestCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_best_categories", []interface{}{after, limit})
}

//get_active_categories                  | **DONE** | *NONE* |

func (api *API) GetActiveCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_active_categories", []interface{}{after, limit})
}

//get_recent_categories                  | **DONE** | *NONE* |

func (api *API) GetRecentCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_recent_categories", []interface{}{after, limit})
}

//get_config                             | **DONE** | **DONE** |

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

//get_dynamic_global_properties          | **DONE** | **DONE** |

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

//get_chain_properties                   | **DONE** | **DONE** |

func (api *API) GetChainPropertiesRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_chain_properties", call.EmptyParams)
}

func (api *API) GetChainProperties() (*ChainProperties, error) {
	var resp ChainProperties
	if err := api.caller.Call("get_chain_properties", call.EmptyParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

//get_feed_history                       | **DONE** | **DONE** |

func (api *API) GetFeedHistoryRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_feed_history", call.EmptyParams)
}

func (api *API) GetFeedHistory() (*FeedHistory, error) {
	var resp FeedHistory
	if err := api.caller.Call("get_feed_history", call.EmptyParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

//get_current_median_history_price       | **DONE** | **DONE** |

func (api *API) GetCurrentMedianHistoryPriceRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_current_median_history_price", call.EmptyParams)
}

func (api *API) GetCurrentMedianHistoryPrice() (*CurrentMedianHistoryPrice, error) {
	var resp CurrentMedianHistoryPrice
	if err := api.caller.Call("get_current_median_history_price", call.EmptyParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

//get_witness_schedule                   | **DONE** | **DONE** |

func (api *API) GetWitnessScheduleRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_witness_schedule", call.EmptyParams)
}

func (api *API) GetWitnessSchedule() (*WitnessSchedule, error) {
	var resp WitnessSchedule
	if err := api.caller.Call("get_witness_schedule", call.EmptyParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

//get_hardfork_version                   | **DONE** | **DONE** |

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

//get_next_scheduled_hardfork            | **DONE** | **DONE** |

func (api *API) GetNextScheduledHardforkRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_next_scheduled_hardfork", call.EmptyParams)
}

func (api *API) GetNextScheduledHardfork() (*NextScheduledHardfork, error) {
	var resp NextScheduledHardfork
	if err := api.caller.Call("get_next_scheduled_hardfork", call.EmptyParams, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

//get_key_references                     | *NONE* | *NONE* |
//Unfortunately to say what this command does is not possible. (Any call to it leads to an error).

//get_accounts                           | **DONE** | ***PARTIALLY DONE*** |

func (api *API) GetAccountsRaw(accountNames []string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_accounts", [][]string{accountNames})
}

func (api *API) GetAccounts(accountNames []string) ([]*Account, error) {
	var resp []*Account
	if err := api.caller.Call("get_accounts", [][]string{accountNames}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_account_references                 | *NONE* | *NONE* |
//Unfortunately to say what this command does is not possible. (Any call to it leads to an error).

//lookup_account_names                   | **DONE** | *NONE* |

func (api *API) LookupAccountNamesRaw(accountNames []string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "lookup_account_names", [][]string{accountNames})
}

//lookup_accounts                        | **DONE** | **DONE** |

func (api *API) LookupAccountsRaw(lowerBoundName string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "lookup_accounts", []interface{}{lowerBoundName, limit})
}

func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]string, error) {
	var resp []string
	if err := api.caller.Call("lookup_accounts", []interface{}{lowerBoundName, limit}, &resp); err != nil {
		return []string{""}, err
	}
	return resp, nil
}

//get_account_count                      | **DONE** | **DONE** |

func (api *API) GetAccountCountRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_account_count", call.EmptyParams)
}

func (api *API) GetAccountCount() (uint32, error) {
	var resp uint32
	if err := api.caller.Call("get_account_count", call.EmptyParams, &resp); err != nil {
		return 0, err
	}
	return resp, nil
}

//get_conversion_requests                | **DONE** | **DONE** |

func (api *API) GetConversionRequestsRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_conversion_requests", []string{accountName})
}

func (api *API) GetConversionRequests(accountName string) ([]*ConversionRequests, error) {
	var resp []*ConversionRequests
	if err := api.caller.Call("get_conversion_requests", []string{accountName}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_account_history                    | **DONE** | *NONE* |

func (api *API) GetAccountHistoryRaw(account string, from uint64, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_account_history", []interface{}{account, from, limit})
}

//get_owner_history                      | **DONE** | *NONE* |

func (api *API) GetOwnerHistoryRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_owner_history", []interface{}{accountName})
}

//get_recovery_request                   | **DONE** | *NONE* |

func (api *API) GetRecoveryRequestRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_recovery_request", []interface{}{accountName})
}

//get_escrow                             | **DONE** | *NONE* |

func (api *API) GetEscrowRaw(from string, escrow_id uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_escrow", []interface{}{from, escrow_id})
}

//get_withdraw_routes                    | **DONE** | *NONE* |

func (api *API) GetWuthdrawRoutesRaw(accountName string, withdraw_route_type string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_withdraw_routes", []interface{}{accountName, withdraw_route_type})
}

//get_account_bandwidth                  | **DONE** | *NONE* |

func (api *API) GetAccountBandwidthRaw(accountName string, bandwidth_type uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_account_bandwidth", []interface{}{accountName, bandwidth_type})
}

//get_savings_withdraw_from              | **DONE** | **DONE** |

func (api *API) GetSavingsWithdrawFromRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_savings_withdraw_from", []interface{}{accountName})
}

func (api *API) GetSavingsWithdrawFrom(accountName string) ([]*SavingsWithdraw, error) {
	var resp []*SavingsWithdraw
	if err := api.caller.Call("get_savings_withdraw_from", []string{accountName}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_savings_withdraw_to                | **DONE** | **DONE** |

func (api *API) GetSavingsWithdrawToRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_savings_withdraw_to", []interface{}{accountName})
}

func (api *API) GetSavingsWithdrawTo(accountName string) ([]*SavingsWithdraw, error) {
	var resp []*SavingsWithdraw
	if err := api.caller.Call("get_savings_withdraw_to", []string{accountName}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_order_book                         | **DONE** | **DONE** |

func (api *API) GetOrderBookRaw(limit uint32) (*json.RawMessage, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	return call.Raw(api.caller, "get_order_book", []interface{}{limit})
}

func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	var resp *OrderBook
	if err := api.caller.Call("get_order_book", []interface{}{limit}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_open_orders                        | **DONE** | **DONE** |

func (api *API) GetOpenOrdersRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_open_orders", []string{accountName})
}

func (api *API) GetOpenOrders(accountName string) ([]*OpenOrders, error) {
	var resp []*OpenOrders
	if err := api.caller.Call("get_open_orders", []string{accountName}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_liquidity_queue                    | **DONE** | *NONE* |

func (api *API) GetLiquidityQueueRaw(startAccount string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_liquidity_queue", []interface{}{startAccount, limit})
}

//get_transaction_hex                    | **DONE** | *NONE* |
func (api *API) GetTransactionHexRaw(trx *types.Transaction) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_transaction_hex", []interface{}{&trx})
}

//get_transaction                        | **DONE** | **DONE** |

func (api *API) GetTransactionRaw(id string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_transaction", []string{id})
}

func (api *API) GetTransaction(id string) (*types.Transaction, error) {
	var resp types.Transaction
	if err := api.caller.Call("get_transaction", []string{id}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

//get_required_signatures                | *NONE* | *NONE* |

//get_potential_signatures               | **DONE** | **DONE** |

func (api *API) GetPotentialSignaturesRaw(trx *types.Transaction) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_potential_signatures", []interface{}{&trx})
}

func (api *API) GetPotentialSignatures(trx *types.Transaction) ([]string, error) {
	var resp []string
	if err := api.caller.Call("get_potential_signatures", []interface{}{&trx}, &resp); err != nil {
		return []string{""}, err
	}
	return resp, nil
}

//verify_authority                       | **DONE** | **DONE** |

func (api *API) GetVerifyAuthorutyRaw(trx *types.Transaction) (*json.RawMessage, error) {
	return call.Raw(api.caller, "verify_authority", []interface{}{&trx})
}

func (api *API) GetVerifyAuthoruty(trx *types.Transaction) (bool, error) {
	var resp bool
	if err := api.caller.Call("verify_authority", []interface{}{&trx}, &resp); err != nil {
		return false, err
	}
	return resp, nil
}

//verify_account_authority               | *NONE* | *NONE* |

//get_active_votes                       | **DONE** | **DONE** |

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

//get_account_votes                      | **DONE** | **DONE** |

func (api *API) GetAccountVotesRaw(author string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_account_votes", []string{author})
}

func (api *API) GetAccountVotes(author string) ([]*Votes, error) {
	var resp []*Votes
	if err := api.caller.Call("get_account_votes", []string{author}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_content                            | **DONE** | **DONE** |

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

//get_content_replies                    | **DONE** | **DONE** |

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

//get_discussions_by_author_before_date  | **DONE** | **DONE** |

func (api *API) GetDiscussionsByAuthorBeforeDateRaw(Author, Permlink, Date string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_author_before_date", []interface{}{Author, Permlink, Date, limit})
}

func (api *API) GetDiscussionsByAuthorBeforeDate(Author, Permlink, Date string, limit uint32) ([]*Content, error) {
	var resp []*Content
	err := api.caller.Call("get_discussions_by_author_before_date", []interface{}{Author, Permlink, Date, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//get_replies_by_last_update             | **DONE** | **DONE** |

func (api *API) GetRepliesByLastUpdateRaw(startAuthor, startPermlink string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_replies_by_last_update", []interface{}{startAuthor, startPermlink, limit})
}

func (api *API) GetRepliesByLastUpdate(startAuthor, startPermlink string, limit uint32) ([]*Content, error) {
	var resp []*Content
	err := api.caller.Call("get_replies_by_last_update", []interface{}{startAuthor, startPermlink, limit}, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//get_witnesses                          | **DONE** | **DONE** |

func (api *API) GetWitnessesRaw(id []uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_witnesses", [][]uint32{id})
}

func (api *API) GetWitnesses(id []uint32) ([]*Witness, error) {
	var resp []*Witness
	if err := api.caller.Call("get_witnesses", [][]uint32{id}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//get_witness_by_account                 | **DONE** | **DONE** |

func (api *API) GetWitnessByAccountRaw(author string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_witness_by_account", []string{author})
}

func (api *API) GetWitnessByAccount(author string) (*Witness, error) {
	var resp Witness
	if err := api.caller.Call("get_witness_by_account", []string{author}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

//get_witnesses_by_vote                  | **DONE** | **DONE** |

func (api *API) GetWitnessByVoteRaw(author string, limit uint) (*json.RawMessage, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	return call.Raw(api.caller, "get_witnesses_by_vote", []interface{}{author, limit})
}

func (api *API) GetWitnessByVote(author string, limit uint) ([]*Witness, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	var resp []*Witness
	if err := api.caller.Call("get_witnesses_by_vote", []interface{}{author, limit}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//lookup_witness_accounts                | **DONE** | **DONE** |

func (api *API) LookupWitnessAccountsRaw(author string, limit uint) (*json.RawMessage, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	return call.Raw(api.caller, "lookup_witness_accounts", []interface{}{author, limit})
}

func (api *API) LookupWitnessAccounts(author string, limit uint) ([]string, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	var resp []string
	if err := api.caller.Call("lookup_witness_accounts", []interface{}{author, limit}, &resp); err != nil {
		return []string{""}, err
	}
	return resp, nil
}

//get_witness_count                      | **DONE** | **DONE** |

func (api *API) GetWitnessCountRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_witness_count", call.EmptyParams)
}

func (api *API) GetWitnessCount() (uint32, error) {
	var resp uint32
	if err := api.caller.Call("get_witness_count", call.EmptyParams, &resp); err != nil {
		return 0, err
	}
	return resp, nil
}

//get_active_witnesses                   | **DONE** | **DONE** |

func (api *API) GetActiveWitnessesRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_active_witnesses", call.EmptyParams)
}

func (api *API) GetActiveWitnesses() ([]string, error) {
	var resp []string
	if err := api.caller.Call("get_active_witnesses", call.EmptyParams, &resp); err != nil {
		return []string{""}, err
	}
	return resp, nil
}

//get_miner_queue                        | **DONE** | **DONE** |

func (api *API) GetMinerQueueRaw() (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_miner_queue", call.EmptyParams)
}

func (api *API) GetMinerQueue() ([]string, error) {
	var resp []string
	if err := api.caller.Call("get_miner_queue", call.EmptyParams, &resp); err != nil {
		return []string{""}, err
	}
	return resp, nil
}
