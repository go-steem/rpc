# Database API

This package adds support for `database_api`.

## State 

| **ID** | **Command Name** | **Version** |
| :-: | :-: | :-: |
| 0 | set_subscribe_callback | **NONE** |
| 1 | set_pending_transaction_callback | **NONE** |
| 2 | set_block_applied_callback | **NONE** |
| 3 | cancel_all_subscriptions | **NONE** |
| 4 | get_trending_tags | **DONE** |
| 5 | get_tags_used_by_author | **RAW** |
| 6 | get_post_discussions_by_payout | **DONE** |
| 7 | get_comment_discussions_by_payout | **DONE** |
| 8 | get_discussions_by_trending | **DONE** |
| 9 | get_discussions_by_trending30 | **DONE** |
| 10 | get_discussions_by_created | **DONE** |
| 11 | get_discussions_by_active | **DONE** |
| 12 | get_discussions_by_cashout | **DONE** |
| 13 | get_discussions_by_payout | **DONE** |
| 14 | get_discussions_by_votes | **DONE** |
| 15 | get_discussions_by_children | **DONE** |
| 16 | get_discussions_by_hot | **DONE** |
| 17 | get_discussions_by_feed | **DONE** |
| 18 | get_discussions_by_blog | **DONE** |
| 19 | get_discussions_by_comments | **DONE** |
| 20 | get_discussions_by_promoted | **DONE** |
| 21 | get_block_header | **DONE** |
| 22 | get_block | **DONE** |
| 23 | get_ops_in_block | **DONE** |
| 24 | get_state | **RAW** |
| 25 | get_trending_categories | **DONE** |
| 26 | get_best_categories | **RAW** |
| 27 | get_active_categories | **RAW** |
| 28 | get_recent_categories | **RAW** |
| 29 | get_config | **DONE** |
| 30 | get_dynamic_global_properties | **DONE** |
| 31 | get_chain_properties | **DONE** |
| 32 | get_feed_history | **DONE** |
| 33 | get_current_median_history_price | **DONE** |
| 34 | get_witness_schedule | **DONE** |
| 35 | get_hardfork_version | **DONE** |
| 36 | get_next_scheduled_hardfork | **DONE** |
| 37 | get_key_references | **RAW** |
| 38 | get_accounts | **DONE** |
| 39 | get_account_references | **RAW** |
| 40 | lookup_account_names | **DONE** |
| 41 | lookup_accounts | **DONE** |
| 42 | get_account_count | **DONE** |
| 43 | get_conversion_requests | **DONE** |
| 44 | get_account_history | **DONE** |
| 45 | get_owner_history | **RAW** |
| 46 | get_recovery_request | **RAW** |
| 47 | get_escrow | **RAW** |
| 48 | get_withdraw_routes | **RAW** |
| 49 | get_account_bandwidth | **RAW** |
| 50 | get_savings_withdraw_from | **DONE** |
| 51 | get_savings_withdraw_to | **DONE** |
| 52 | get_order_book | **DONE** |
| 53 | get_open_orders | **DONE** |
| 54 | get_liquidity_queue | **RAW** |
| 55 | get_transaction_hex | **DONE** |
| 56 | get_transaction | **DONE** |
| 57 | get_required_signatures | **DONE** |
| 58 | get_potential_signatures | **DONE** |
| 59 | verify_authority | **DONE** |
| 60 | verify_account_authority  | **RAW** |
| 61 | get_active_votes | **DONE** |
| 62 | get_account_votes | **DONE** |
| 63 | get_content | **DONE** |
| 64 | get_content_replies | **DONE** |
| 65 | get_discussions_by_author_before_date | **DONE** |
| 66 | get_replies_by_last_update | **DONE** |
| 67 | get_witnesses | **DONE** |
| 68 | get_witness_by_account | **DONE** |
| 69 | get_witnesses_by_vote | **DONE** |
| 70 | lookup_witness_accounts | **DONE** |
| 71 | get_witness_count | **DONE** |
| 72 | get_active_witnesses | **DONE** |
| 73 | get_miner_queue | **DONE** |
| 74 | get_reward_fund | **RAW** |
| 75 | get_vesting_delegations | **RAW** |

## License

MIT, see the `LICENSE` file.