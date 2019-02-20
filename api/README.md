# API

This package adds support for `steem api`.

## State

| **ID** | **Command Name** | **Version** |
| :-: | :-: | :-: |
| **account_by_key** |
| 1 | get_key_references | **RAW** |
| **account_history** |
| 1 | get_account_history | **DONE** |
| **committee_api** |
| 1 | get_committee_request | **DONE** |
| 2 | get_committee_request_votes | **DONE** |
| 3 | get_committee_requests_list | **DONE** |
| **database_api** |
| 1 | get_account_count | **DONE** |
| 2 | get_accounts | **DONE** |
| 3 | get_block | **DONE** |
| 4 | get_block_header | **DONE** |
| 5 | get_chain_properties | **DONE** |
| 6 | get_config | **DONE** |
| 7 | get_database_info | **DONE** |
| 8 | get_dynamic_global_properties | **DONE** |
| 9 | get_escrow | **DONE** |
| 10 | get_expiring_vesting_delegations | **DONE** |
| 11| get_hardfork_version | **DONE** |
| 12| get_next_scheduled_hardfork | **DONE** |
| 13 | get_owner_history | **DONE** |
| 14 | get_potential_signatures | **DONE** |
| 15 | get_proposed_transaction | **DONE** |
| 16 | get_recovery_request | **DONE** |
| 17 | get_required_signatures | **DONE** |
| 18 | get_transaction_hex | **DONE** |
| 19 | get_vesting_delegations | **DONE** |
| 20 | get_withdraw_routes | **DONE** |
| 21 | lookup_account_names | **DONE** |
| 22 | lookup_accounts | **DONE** |
| 23 | verify_account_authority | **DONE** |
| 24 | verify_authority | **DONE** |
| **follow** |
| 1 | get_blog | **DONE** |
| 2 | get_blog_authors | **DONE** |
| 3 | get_blog_entries | **DONE** |
| 4 | get_feed | **DONE** |
| 5 | get_feed_entries | **DONE** |
| 6 | get_follow_count | **DONE** |
| 7 | get_followers | **DONE** |
| 8 | get_following | **DONE** |
| 9 | get_reblogged_by | **DONE** |
| **invite_api** |
| 1 | get_invites_list | **DONE** |
| 2 | get_invite_by_id | **DONE** |
| 3 | get_invite_by_key | **DONE** |
| **network_broadcast_api** |
| 1 | broadcast_block | **NONE** |
| 2 | broadcast_transaction | **DONE** |
| 3 | broadcast_transaction_synchronous | **DONE** |
| 4 | broadcast_transaction_with_callback | **NONE** |
| **operation_history** |
| 1 | get_ops_in_block | **DONE** |
| 2 | get_transaction | **DONE** |
| **social_network** |
| 1 | get_account_votes | **DONE** |
| 2 | get_active_votes | **DONE** |
| 3 | get_all_content_replies | **DONE** |
| 4 | get_content | **DONE** |
| 5 | get_content_replies | **DONE** |
| 6 | get_replies_by_last_update | **DONE** |
| **tags** |
| 1 | get_discussions_by_active | **DONE** |
| 2 | get_discussions_by_author_before_date | **DONE** |
| 3 | get_discussions_by_blog | **DONE** |
| 4 | get_discussions_by_cashout | **DONE** |
| 5 | get_discussions_by_children | **DONE** |
| 6 | get_discussions_by_contents | **DONE** |
| 7 | get_discussions_by_created | **DONE** |
| 8 | get_discussions_by_feed | **DONE** |
| 9 | get_discussions_by_hot | **DONE** |
| 10 | get_discussions_by_payout | **DONE** |
| 11 | get_discussions_by_trending | **DONE** |
| 12 | get_discussions_by_votes | **DONE** |
| 13 | get_languages | **DONE** |
| 14 | get_tags_used_by_author | **RAW** |
| 15 | get_trending_tags | **DONE** |
| **witness_api** |
| 1 | get_active_witnesses | **DONE** |
| 2 | get_miner_queue | **DONE** |
| 3 | get_witness_by_account | **DONE** |
| 4 | get_witness_count | **DONE** |
| 5 | get_witness_schedule | **DONE** |
| 6 | get_witnesses | **DONE** |
| 7 | get_witnesses_by_vote | **DONE** |
| 8 | lookup_witness_accounts | **DONE** |

## License

MIT, see the `LICENSE` file.
