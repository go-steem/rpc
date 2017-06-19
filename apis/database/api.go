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

//get_trending_tags                      | **DONE** | **DONE** |

func (api *API) GetTrendingTagsRaw(afterTag string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_trending_tags", []interface{}{afterTag, limit})
}

func (api *API) GetTrendingTags(afterTag string, limit uint32) ([]*TrendingTags, error) {
	raw, err := api.GetTrendingTagsRaw(afterTag, limit)
	if err != nil {
		return nil, err
	}
	var resp []*TrendingTags
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_trending_tags response", APIID)
	}
	return resp, nil
}

//get_tags_used_by_author                | **DONE** | *NONE* |

func (api *API) GetTagsUsedByAuthorRaw(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_tags_used_by_author", []interface{}{accountName})
}

//get_discussions_by_trending            | **DONE** | **DONE** |

func (api *API) GetDiscussionsByTrendingRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_trending", query)
}

func (api *API) GetDiscussionsByTrending(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByTrendingRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_trending response", APIID)
	}
	return resp, nil
}

//get_discussions_by_trending30          | **DONE** | **DONE** |

func (api *API) GetDiscussionsByTrending30Raw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_trending30", query)
}

func (api *API) GetDiscussionsByTrending30(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByTrending30Raw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_trending30 response", APIID)
	}
	return resp, nil
}

//get_discussions_by_created             | **DONE** | **DONE** |

func (api *API) GetDiscussionsByCreatedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_created", query)
}

func (api *API) GetDiscussionsByCreated(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByCreatedRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_created response", APIID)
	}
	return resp, nil
}

//get_discussions_by_active              | **DONE** | **DONE** |

func (api *API) GetDiscussionsByActiveRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_active", query)
}

func (api *API) GetDiscussionsByActive(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByActiveRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_active response", APIID)
	}
	return resp, nil
}

//get_discussions_by_cashout             | **DONE** | **DONE** |

func (api *API) GetDiscussionsByCashoutRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_cashout", query)
}

func (api *API) GetDiscussionsByCashout(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByCashoutRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_cashout response", APIID)
	}
	return resp, nil
}

//get_discussions_by_payout              | **DONE** | **DONE** |

func (api *API) GetDiscussionsByPayoutRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_payout", query)
}

func (api *API) GetDiscussionsByPayout(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByPayoutRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_payout response", APIID)
	}
	return resp, nil
}

//get_discussions_by_votes               | **DONE** | **DONE** |

func (api *API) GetDiscussionsByVotesRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_votes", query)
}

func (api *API) GetDiscussionsByVotes(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByVotesRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_votes response", APIID)
	}
	return resp, nil
}

//get_discussions_by_children            | **DONE** | **DONE** |

func (api *API) GetDiscussionsByChildrenRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_children", query)
}

func (api *API) GetDiscussionsByChildren(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByChildrenRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_children response", APIID)
	}
	return resp, nil
}

//get_discussions_by_hot                 | **DONE** | **DONE** |

func (api *API) GetDiscussionsByHotRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_hot", query)
}

func (api *API) GetDiscussionsByHot(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByHotRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_hot response", APIID)
	}
	return resp, nil
}

//get_discussions_by_feed                | **DONE** | **DONE** |

func (api *API) GetDiscussionsByFeedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_feed", query)
}

func (api *API) GetDiscussionsByFeed(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByFeedRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_feed response", APIID)
	}
	return resp, nil
}

//get_discussions_by_blog                | **DONE** | **DONE** |

func (api *API) GetDiscussionsByBlogRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_blog", query)
}

func (api *API) GetDiscussionsByBlog(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByBlogRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_blog response", APIID)
	}
	return resp, nil
}

//get_discussions_by_comments            | **DONE** | **DONE** |

func (api *API) GetDiscussionsByCommentsRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_comments", query)
}

func (api *API) GetDiscussionsByComments(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByCommentsRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_comments response", APIID)
	}
	return resp, nil
}

//get_discussions_by_promoted            | **DONE** | **DONE** |

func (api *API) GetDiscussionsByPromotedRaw(query *DiscussionQuery) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_promoted", query)
}

func (api *API) GetDiscussionsByPromoted(query *DiscussionQuery) ([]*Content, error) {
	raw, err := api.GetDiscussionsByPromotedRaw(query)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_promoted response", APIID)
	}
	return resp, nil
}

//get_block_header                       | **DONE** | **DONE** |

func (api *API) GetBlockHeaderRaw(blockNum uint32) (*json.RawMessage, error) {
	return api.Raw("get_block_header", []uint32{blockNum})
}

func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	raw, err := api.GetBlockHeaderRaw(blockNum)
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

//get_block                              | **DONE** | ***PARTIALLY DONE*** |

func (api *API) GetBlockRaw(blockNum uint32) (*json.RawMessage, error) {
	return api.Raw("get_block", []uint32{blockNum})
}

func (api *API) GetBlock(blockNum uint32) (*Block, error) {
	raw, err := api.GetBlockRaw(blockNum)
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

//get_ops_in_block                       | **DONE** | ***PARTIALLY DONE*** |

func (api *API) GetOpsInBlockRaw(blockNum uint32, only_virtual bool) (*json.RawMessage, error) {
	return api.Raw("get_ops_in_block", []interface{}{blockNum, only_virtual})
}

func (api *API) GetOpsInBlock(blockNum uint32, only_virtual bool) ([]*OpsInBlock, error) {
	raw, err := api.GetOpsInBlockRaw(blockNum, only_virtual)
	if err != nil {
		return nil, err
	}
	var resp []*OpsInBlock
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_ops_in_block response", APIID)
	}
	return resp, nil
}

//get_state                              | **DONE** | *NONE* |

func (api *API) GetStateRaw(path string) (*json.RawMessage, error) {
	return api.Raw("get_state", []string{path})
}

//get_trending_categories                | **DONE** | **DONE** |

func (api *API) GetTrendingCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_trending_categories", []interface{}{after, limit})
}

func (api *API) GetTrendingCategories(after string, limit uint32) ([]*Categories, error) {
	raw, err := api.GetTrendingCategoriesRaw(after, limit)
	if err != nil {
		return nil, err
	}
	var resp []*Categories
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_trending_categories response", APIID)
	}
	return resp, nil
}

//get_best_categories                    | **DONE** | *NONE* |

func (api *API) GetBestCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_best_categories", []interface{}{after, limit})
}

//get_active_categories                  | **DONE** | *NONE* |

func (api *API) GetActiveCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_active_categories", []interface{}{after, limit})
}

//get_recent_categories                  | **DONE** | *NONE* |

func (api *API) GetRecentCategoriesRaw(after string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_recent_categories", []interface{}{after, limit})
}

//get_config                             | **DONE** | **DONE** |

func (api *API) GetConfigRaw() (*json.RawMessage, error) {
	return api.Raw("get_config", EmptyParams)
}

func (api *API) GetConfig() (*Config, error) {
	raw, err := api.GetConfigRaw()
	if err != nil {
		return nil, err
	}
	var resp Config
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_config response", APIID)
	}
	return &resp, nil
}

//get_dynamic_global_properties          | **DONE** | **DONE** |

func (api *API) GetDynamicGlobalPropertiesRaw() (*json.RawMessage, error) {
	return api.Raw("get_dynamic_global_properties", EmptyParams)
}

func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	raw, err := api.GetDynamicGlobalPropertiesRaw()
	if err != nil {
		return nil, err
	}
	var resp DynamicGlobalProperties
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_dynamic_global_properties response", APIID)
	}
	return &resp, nil
}

//get_chain_properties                   | **DONE** | **DONE** |

func (api *API) GetChainPropertiesRaw() (*json.RawMessage, error) {
	return api.Raw("get_chain_properties", EmptyParams)
}

func (api *API) GetChainProperties() (*ChainProperties, error) {
	raw, err := api.GetChainPropertiesRaw()
	if err != nil {
		return nil, err
	}
	var resp ChainProperties
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_chain_properties response", APIID)
	}
	return &resp, nil
}

//get_feed_history                       | **DONE** | **DONE** |

func (api *API) GetFeedHistoryRaw() (*json.RawMessage, error) {
	return api.Raw("get_feed_history", EmptyParams)
}

func (api *API) GetFeedHistory() (*FeedHistory, error) {
	raw, err := api.GetFeedHistoryRaw()
	if err != nil {
		return nil, err
	}
	var resp FeedHistory
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_feed_history response", APIID)
	}
	return &resp, nil
}

//get_current_median_history_price       | **DONE** | **DONE** |

func (api *API) GetCurrentMedianHistoryPriceRaw() (*json.RawMessage, error) {
	return api.Raw("get_current_median_history_price", EmptyParams)
}

func (api *API) GetCurrentMedianHistoryPrice() (*CurrentMedianHistoryPrice, error) {
	raw, err := api.GetCurrentMedianHistoryPriceRaw()
	if err != nil {
		return nil, err
	}
	var resp CurrentMedianHistoryPrice
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_current_median_history_price response", APIID)
	}
	return &resp, nil
}

//get_witness_schedule                   | **DONE** | **DONE** |

func (api *API) GetWitnessScheduleRaw() (*json.RawMessage, error) {
	return api.Raw("get_witness_schedule", EmptyParams)
}

func (api *API) GetWitnessSchedule() (*WitnessSchedule, error) {
	raw, err := api.GetWitnessScheduleRaw()
	if err != nil {
		return nil, err
	}
	var resp WitnessSchedule
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_witness_schedule response", APIID)
	}
	return &resp, nil
}

//get_hardfork_version                   | **DONE** | **DONE** |

func (api *API) GetHardforkVersionRaw() (*json.RawMessage, error) {
	return api.Raw("get_hardfork_version", EmptyParams)
}

func (api *API) GetHardforkVersion() (string, error) {
	raw, err := api.GetHardforkVersionRaw()
	if err != nil {
		return "", err
	}
	var resp string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return "", errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_hardfork_version response", APIID)
	}
	return resp, nil
}

//get_next_scheduled_hardfork            | **DONE** | **DONE** |

func (api *API) GetNextScheduledHardforkRaw() (*json.RawMessage, error) {
	return api.Raw("get_next_scheduled_hardfork", EmptyParams)
}

func (api *API) GetNextScheduledHardfork() (*NextScheduledHardfork, error) {
	raw, err := api.GetNextScheduledHardforkRaw()
	if err != nil {
		return nil, err
	}
	var resp NextScheduledHardfork
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_next_scheduled_hardfork response", APIID)
	}
	return &resp, nil
}

//get_key_references                     | *NONE* | *NONE* |
//Unfortunately to say what this command does is not possible. (Any call to it leads to an error).

//get_accounts                           | **DONE** | ***PARTIALLY DONE*** |

func (api *API) GetAccountsRaw(accountNames []string) (*json.RawMessage, error) {
	return api.Raw("get_accounts", [][]string{accountNames})
}

func (api *API) GetAccounts(accountNames []string) ([]*Account, error) {
	raw, err := api.GetAccountsRaw(accountNames)
	if err != nil {
		return nil, err
	}
	var resp []*Account
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_accounts response", APIID)
	}
	return resp, nil
}

//get_account_references                 | *NONE* | *NONE* |
//Unfortunately to say what this command does is not possible. (Any call to it leads to an error).

//lookup_account_names                   | **DONE** | *NONE* |

func (api *API) LookupAccountNamesRaw(accountNames []string) (*json.RawMessage, error) {
	return api.Raw("lookup_account_names", [][]string{accountNames})
}

//lookup_accounts                        | **DONE** | **DONE** |

func (api *API) LookupAccountsRaw(lowerBoundName string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("lookup_accounts", []interface{}{lowerBoundName, limit})
}

func (api *API) LookupAccounts(lowerBoundName string, limit uint32) ([]string, error) {
	raw, err := api.LookupAccountsRaw(lowerBoundName, limit)
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal lookup_accounts response", APIID)
	}
	return resp, nil
}

//get_account_count                      | **DONE** | **DONE** |

func (api *API) GetAccountCountRaw() (*json.RawMessage, error) {
	return api.Raw("get_account_count", EmptyParams)
}

func (api *API) GetAccountCount() (uint32, error) {
	raw, err := api.GetAccountCountRaw()
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_account_count response", APIID)
	}
	return resp, nil
}

//get_conversion_requests                | **DONE** | **DONE** |

func (api *API) GetConversionRequestsRaw(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_conversion_requests", []string{accountName})
}

func (api *API) GetConversionRequests(accountName string) ([]*ConversionRequests, error) {
	raw, err := api.GetConversionRequestsRaw(accountName)
	if err != nil {
		return nil, err
	}
	var resp []*ConversionRequests
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_conversion_requests response", APIID)
	}
	return resp, nil
}

//get_account_history                    | **DONE** | *NONE* |

func (api *API) GetAccountHistoryRaw(account string, from uint64, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_account_history", []interface{}{account, from, limit})
}

//get_owner_history                      | **DONE** | *NONE* |

func (api *API) GetOwnerHistoryRaw(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_owner_history", []interface{}{accountName})
}

//get_recovery_request                   | **DONE** | *NONE* |

func (api *API) GetRecoveryRequestRaw(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_recovery_request", []interface{}{accountName})
}

//get_escrow                             | **DONE** | *NONE* |

func (api *API) GetEscrowRaw(from string, escrow_id uint32) (*json.RawMessage, error) {
	return api.Raw("get_escrow", []interface{}{from, escrow_id})
}

//get_withdraw_routes                    | **DONE** | *NONE* |

func (api *API) GetWuthdrawRoutesRaw(accountName string, withdraw_route_type string) (*json.RawMessage, error) {
	return api.Raw("get_withdraw_routes", []interface{}{accountName, withdraw_route_type})
}

//get_account_bandwidth                  | **DONE** | *NONE* |

func (api *API) GetAccountBandwidthRaw(accountName string, bandwidth_type uint32) (*json.RawMessage, error) {
	return api.Raw("get_account_bandwidth", []interface{}{accountName, bandwidth_type})
}

//get_savings_withdraw_from              | **DONE** | **DONE** |

func (api *API) GetSavingsWithdrawFromRaw(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_savings_withdraw_from", []interface{}{accountName})
}

func (api *API) GetSavingsWithdrawFrom(accountName string) ([]*SavingsWithdraw, error) {
	raw, err := api.GetSavingsWithdrawFromRaw(accountName)
	if err != nil {
		return nil, err
	}
	var resp []*SavingsWithdraw
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_savings_withdraw_from response", APIID)
	}
	return resp, nil
}

//get_savings_withdraw_to                | **DONE** | **DONE** |

func (api *API) GetSavingsWithdrawToRaw(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_savings_withdraw_to", []interface{}{accountName})
}

func (api *API) GetSavingsWithdrawTo(accountName string) ([]*SavingsWithdraw, error) {
	raw, err := api.GetSavingsWithdrawToRaw(accountName)
	if err != nil {
		return nil, err
	}
	var resp []*SavingsWithdraw
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_savings_withdraw_to response", APIID)
	}
	return resp, nil
}

//get_order_book                         | **DONE** | **DONE** |

func (api *API) GetOrderBookRaw(limit uint32) (*json.RawMessage, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	return api.Raw("get_order_book", []interface{}{limit})
}

func (api *API) GetOrderBook(limit uint32) (*OrderBook, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	raw, err := api.GetOrderBookRaw(limit)
	if err != nil {
		return nil, err
	}
	var resp *OrderBook
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_order_book response", APIID)
	}
	return resp, nil
}

//get_open_orders                        | **DONE** | **DONE** |

func (api *API) GetOpenOrdersRaw(accountName string) (*json.RawMessage, error) {
	return api.Raw("get_open_orders", []string{accountName})
}

func (api *API) GetOpenOrders(accountName string) ([]*OpenOrders, error) {
	raw, err := api.GetOpenOrdersRaw(accountName)
	if err != nil {
		return nil, err
	}
	var resp []*OpenOrders
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_open_orders response", APIID)
	}
	return resp, nil
}

//get_liquidity_queue                    | **DONE** | *NONE* |

func (api *API) GetLiquidityQueueRaw(startAccount string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_liquidity_queue", []interface{}{startAccount, limit})
}

//get_transaction_hex                    | **DONE** | *NONE* |
func (api *API) GetTransactionHexRaw(trx *types.Transaction) (*json.RawMessage, error) {
	return api.Raw("get_transaction_hex", []interface{}{&trx})
}

//get_transaction                        | **DONE** | **DONE** |

func (api *API) GetTransactionRaw(id string) (*json.RawMessage, error) {
	return api.Raw("get_transaction", []string{id})
}

func (api *API) GetTransaction(id string) (*types.Transaction, error) {
	raw, err := api.GetTransactionRaw(id)
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

//get_potential_signatures               | **DONE** | **DONE** |

func (api *API) GetPotentialSignaturesRaw(trx *types.Transaction) (*json.RawMessage, error) {
	return api.Raw("get_potential_signatures", []interface{}{&trx})
}

func (api *API) GetPotentialSignatures(trx *types.Transaction) ([]string, error) {
	raw, err := api.GetPotentialSignaturesRaw(trx)
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_potential_signatures response", APIID)
	}
	return resp, nil
}

//verify_authority                       | **DONE** | **DONE** |

func (api *API) GetVerifyAuthorutyRaw(trx *types.Transaction) (*json.RawMessage, error) {
	return api.Raw("verify_authority", []interface{}{&trx})
}

func (api *API) GetVerifyAuthoruty(trx *types.Transaction) (bool, error) {
	raw, err := api.GetVerifyAuthorutyRaw(trx)
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

//get_active_votes                       | **DONE** | **DONE** |

func (api *API) GetActiveVotesRaw(author, permlink string) (*json.RawMessage, error) {
	return api.Raw("get_active_votes", []string{author, permlink})
}

func (api *API) GetActiveVotes(author, permlink string) ([]*VoteState, error) {
	raw, err := api.GetActiveVotesRaw(author, permlink)
	if err != nil {
		return nil, err
	}
	var resp []*VoteState
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_active_votes response", APIID)
	}
	return resp, nil
}

//get_account_votes                      | **DONE** | **DONE** |

func (api *API) GetAccountVotesRaw(author string) (*json.RawMessage, error) {
	return api.Raw("get_account_votes", []string{author})
}

func (api *API) GetAccountVotes(author string) ([]*Votes, error) {
	raw, err := api.GetAccountVotesRaw(author)
	if err != nil {
		return nil, err
	}
	var resp []*Votes
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_account_votes response", APIID)
	}
	return resp, nil
}

//get_content                            | **DONE** | **DONE** |

func (api *API) GetContentRaw(author, permlink string) (*json.RawMessage, error) {
	return api.Raw("get_content", []string{author, permlink})
}

func (api *API) GetContent(author, permlink string) (*Content, error) {
	raw, err := api.GetContentRaw(author, permlink)
	if err != nil {
		return nil, err
	}
	var resp Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_content response", APIID)
	}
	return &resp, nil
}

//get_content_replies                    | **DONE** | **DONE** |

func (api *API) GetContentRepliesRaw(parentAuthor, parentPermlink string) (*json.RawMessage, error) {
	return api.Raw("get_content_replies", []string{parentAuthor, parentPermlink})
}

func (api *API) GetContentReplies(parentAuthor, parentPermlink string) ([]*Content, error) {
	raw, err := api.GetContentRepliesRaw(parentAuthor, parentPermlink)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_content_replies response", APIID)
	}
	return resp, nil
}

//get_discussions_by_author_before_date  | **DONE** | **DONE** |

func (api *API) GetDiscussionsByAuthorBeforeDateRaw(Author, Permlink, Date string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_discussions_by_author_before_date", []interface{}{Author, Permlink, Date, limit})
}

func (api *API) GetDiscussionsByAuthorBeforeDate(Author, Permlink, Date string, limit uint32) ([]*Content, error) {
	raw, err := api.GetDiscussionsByAuthorBeforeDateRaw(Author, Permlink, Date, limit)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_discussions_by_author_before_date response", APIID)
	}
	return resp, nil
}

//get_replies_by_last_update             | **DONE** | **DONE** |

func (api *API) GetRepliesByLastUpdateRaw(startAuthor, startPermlink string, limit uint32) (*json.RawMessage, error) {
	return api.Raw("get_replies_by_last_update", []interface{}{startAuthor, startPermlink, limit})
}

func (api *API) GetRepliesByLastUpdate(startAuthor, startPermlink string, limit uint32) ([]*Content, error) {
	raw, err := api.GetRepliesByLastUpdateRaw(startAuthor, startPermlink, limit)
	if err != nil {
		return nil, err
	}
	var resp []*Content
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_replies_by_last_update response", APIID)
	}
	return resp, nil
}

//get_witnesses                          | **DONE** | **DONE** |

func (api *API) GetWitnessesRaw(id []uint32) (*json.RawMessage, error) {
	return api.Raw("get_witnesses", [][]uint32{id})
}

func (api *API) GetWitnesses(id []uint32) ([]*Witness, error) {
	raw, err := api.GetWitnessesRaw(id)
	if err != nil {
		return nil, err
	}
	var resp []*Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_witnesses response", APIID)
	}
	return resp, nil
}

//get_witness_by_account                 | **DONE** | **DONE** |

func (api *API) GetWitnessByAccountRaw(author string) (*json.RawMessage, error) {
	return api.Raw("get_witness_by_account", []string{author})
}

func (api *API) GetWitnessByAccount(author string) (*Witness, error) {
	raw, err := api.GetWitnessByAccountRaw(author)
	if err != nil {
		return nil, err
	}
	var resp Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_witness_by_account response", APIID)
	}
	return &resp, nil
}

//get_witnesses_by_vote                  | **DONE** | **DONE** |

func (api *API) GetWitnessByVoteRaw(author string, limit uint) (*json.RawMessage, error) {
	return api.Raw("get_witnesses_by_vote", []interface{}{author, limit})
}

func (api *API) GetWitnessByVote(author string, limit uint) ([]*Witness, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	raw, err := api.GetWitnessByVoteRaw(author, limit)
	if err != nil {
		return nil, err
	}
	var resp []*Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_witnesses_by_vote response", APIID)
	}
	return resp, nil
}

//lookup_witness_accounts                | **DONE** | **DONE** |

func (api *API) LookupWitnessAccountsRaw(author string, limit uint) (*json.RawMessage, error) {
	return api.Raw("lookup_witness_accounts", []interface{}{author, limit})
}

func (api *API) LookupWitnessAccounts(author string, limit uint) ([]string, error) {
	if limit > 1000 {
		return nil, errors.New("GetOrderBook: limit must not exceed 1000")
	}
	raw, err := api.LookupWitnessAccountsRaw(author, limit)
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal lookup_witness_accounts response", APIID)
	}
	return resp, nil
}

//get_witness_count                      | **DONE** | **DONE** |

func (api *API) GetWitnessCountRaw() (*json.RawMessage, error) {
	return api.Raw("get_witness_count", EmptyParams)
}

func (api *API) GetWitnessCount() (uint32, error) {
	raw, err := api.GetWitnessCountRaw()
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_witness_count response", APIID)
	}
	return resp, nil
}

//get_active_witnesses                   | **DONE** | **DONE** |

func (api *API) GetActiveWitnessesRaw() (*json.RawMessage, error) {
	return api.Raw("get_active_witnesses", EmptyParams)
}

func (api *API) GetActiveWitnesses() ([]string, error) {
	raw, err := api.GetActiveWitnessesRaw()
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_active_witnesses response", APIID)
	}
	return resp, nil
}

//get_miner_queue                        | **DONE** | **DONE** |

func (api *API) GetMinerQueueRaw() (*json.RawMessage, error) {
	return api.Raw("get_miner_queue", EmptyParams)
}

func (api *API) GetMinerQueue() ([]string, error) {
	raw, err := api.GetMinerQueueRaw()
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to unmarshal get_miner_queue response", APIID)
	}
	return resp, nil
}
