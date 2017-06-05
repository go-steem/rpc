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

//get_trending_tags                      | **DONE** | *NONE* |

func (api *API) GetTrendingTagsRaw(afterTag string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_trending_tags", []interface{}{afterTag, limit})
}

//get_tags_used_by_author                | **DONE** | *NONE* |

func (api *API) GetTagsUsedByAuthorRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_tags_used_by_author", []interface{}{accountName})
}

//get_discussions_by_trending            | **DONE** | *NONE* |

func (api *API) GetDiscussionsByTrendingRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_trending", query)
}

//get_discussions_by_trending30          | **DONE** | *NONE* |

func (api *API) GetDiscussionsByTrending30Raw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_trending30", query)
}

//get_discussions_by_created             | **DONE** | *NONE* |

func (api *API) GetDiscussionsByCreatedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_created", query)
}

//get_discussions_by_active              | **DONE** | *NONE* |

func (api *API) GetDiscussionsByActiveRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_active", query)
}

//get_discussions_by_cashout             | **DONE** | *NONE* |

func (api *API) GetDiscussionsByCashoutRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_cashout", query)
}

//get_discussions_by_payout              | **DONE** | *NONE* |

func (api *API) GetDiscussionsByPayoutRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_payout", query)
}

//get_discussions_by_votes               | **DONE** | *NONE* |

func (api *API) GetDiscussionsByVotesRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_votes", query)
}

//get_discussions_by_children            | **DONE** | *NONE* |

func (api *API) GetDiscussionsByChildrenRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_children", query)
}

//get_discussions_by_hot                 | **DONE** | *NONE* |

func (api *API) GetDiscussionsByHotRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_hot", query)
}

//get_discussions_by_feed                | **DONE** | *NONE* |

func (api *API) GetDiscussionsByFeedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_feed", query)
}

//get_discussions_by_blog                | **DONE** | *NONE* |

func (api *API) GetDiscussionsByBlogRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_blog", query)
}

//get_discussions_by_comments            | **DONE** | *NONE* |

func (api *API) GetDiscussionsByCommentsRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_comments", query)
}

//get_discussions_by_promoted            | **DONE** | *NONE* |

func (api *API) GetDiscussionsByPromotedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_discussions_by_promoted", query)
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

//get_ops_in_block                       | **DONE** | *NONE* |

func (api *API) GetOpsInBlockRaw(blockNum uint32, only_virtual bool) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_ops_in_block", []interface{}{blockNum, only_virtual})
}

//get_state                              | **DONE** | *NONE* |

func (api *API) GetStateRaw(path string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_state", []string{path})
}

//get_trending_categories                | **DONE** | *NONE* |

func (api *API) GetTrendingCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_trending_categories", []interface{}{after, limit})
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

//lookup_accounts                        | **DONE** | *NONE* |

func (api *API) LookupAccountsRaw(lowerBoundName string, limit uint32) (*json.RawMessage, error) {
	return call.Raw(api.caller, "lookup_accounts", []interface{}{lowerBoundName, limit})
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

//get_savings_withdraw_from              | **DONE** | *NONE* |

func (api *API) GetSavingsWithdrawFromRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_savings_withdraw_from", []interface{}{accountName})
}

//get_savings_withdraw_to                | **DONE** | *NONE* |

func (api *API) GetSavingsWithdrawToRaw(accountName string) (*json.RawMessage, error) {
	return call.Raw(api.caller, "get_savings_withdraw_to", []interface{}{accountName})
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

//get_transaction_hex                    | *NONE* | *NONE* |

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

//get_potential_signatures               | *NONE* | *NONE* |

//verify_authority                       | *NONE* | *NONE* |

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
