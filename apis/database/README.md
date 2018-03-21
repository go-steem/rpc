# Database API

This package adds support for `database_api`.

## State 

The following subsections document the API completion. The method names
are taken from `database_api.hpp` in `steemit/steem`.


| **ID** | **Command Name** | **Raw version** | **Full version** |
|:-- |:--------------------------------------:|-------:|-------:|
| 0  | set_subscribe_callback                 | *NONE* | *NONE* |
| 1  | set_pending_transaction_callback       | *NONE* | *NONE* |
| 2  | set_block_applied_callback             | *NONE* | *NONE* |
| 3  | cancel_all_subscriptions               | *NONE* | *NONE* |
| 4  | get_trending_tags                      | **DONE** | **DONE** |
| 5  | get_tags_used_by_author                | **DONE** | *NONE* |
| 6  | get_discussions_by_trending            | **DONE** | **DONE** |
| 7  | get_discussions_by_trending30          | **DONE** | **DONE** |
| 8  | get_discussions_by_created             | **DONE** | **DONE** |
| 9  | get_discussions_by_active              | **DONE** | **DONE** |
| 10 | get_discussions_by_cashout             | **DONE** | **DONE** |
| 11 | get_discussions_by_payout              | **DONE** | **DONE** |
| 12 | get_discussions_by_votes               | **DONE** | **DONE** |
| 13 | get_discussions_by_children            | **DONE** | **DONE** |
| 14 | get_discussions_by_hot                 | **DONE** | **DONE** |
| 15 | get_discussions_by_feed                | **DONE** | **DONE** |
| 16 | get_discussions_by_blog                | **DONE** | **DONE** |
| 17 | get_discussions_by_comments            | **DONE** | **DONE** |
| 18 | get_discussions_by_promoted            | **DONE** | **DONE** |
| 19 | get_block_header                       | **DONE** | **DONE** |
| 20 | get_block                              | **DONE** | **DONE** |
| 21 | get_ops_in_block                       | **DONE** | ***PARTIALLY DONE*** |
| 22 | get_state                              | **DONE** | *NONE* |
| 23 | get_trending_categories                | **DONE** | **DONE** |
| 24 | get_best_categories                    | **DONE** | *NONE* |
| 25 | get_active_categories                  | **DONE** | *NONE* |
| 26 | get_recent_categories                  | **DONE** | *NONE* |
| 27 | get_config                             | **DONE** | **DONE** |
| 28 | get_dynamic_global_properties          | **DONE** | **DONE** |
| 29 | get_chain_properties                   | **DONE** | **DONE** |
| 30 | get_feed_history                       | **DONE** | **DONE** |
| 31 | get_current_median_history_price       | **DONE** | **DONE** |
| 32 | get_witness_schedule                   | **DONE** | **DONE** |
| 33 | get_hardfork_version                   | **DONE** | **DONE** |
| 34 | get_next_scheduled_hardfork            | **DONE** | **DONE** |
| 35 | get_key_references                     | *NONE* | *NONE* |
| 36 | get_accounts                           | **DONE** | ***PARTIALLY DONE*** |
| 37 | get_account_references                 | *NONE* | *NONE* |
| 38 | lookup_account_names                   | **DONE** | *NONE* |
| 39 | lookup_accounts                        | **DONE** | **DONE** |
| 40 | get_account_count                      | **DONE** | **DONE** |
| 41 | get_conversion_requests                | **DONE** | **DONE** |
| 42 | get_account_history                    | **DONE** | *NONE* |
| 43 | get_owner_history                      | **DONE** | *NONE* |
| 44 | get_recovery_request                   | **DONE** | *NONE* |
| 45 | get_escrow                             | **DONE** | *NONE* |
| 46 | get_withdraw_routes                    | **DONE** | *NONE* |
| 47 | get_account_bandwidth                  | **DONE** | *NONE* |
| 48 | get_savings_withdraw_from              | **DONE** | **DONE** |
| 49 | get_savings_withdraw_to                | **DONE** | **DONE** |
| 50 | get_order_book                         | **DONE** | **DONE** |
| 51 | get_open_orders                        | **DONE** | **DONE** |
| 52 | get_liquidity_queue                    | **DONE** | *NONE* |
| 53 | get_transaction_hex                    | *NONE* | *NONE* |
| 54 | get_transaction                        | **DONE** | **DONE** |
| 55 | get_required_signatures                | *NONE* | *NONE* |
| 56 | get_potential_signatures               | *NONE* | *NONE* |
| 57 | verify_authority                       | *NONE* | *NONE* |
| 58 | verify_account_authority               | *NONE* | *NONE* |
| 59 | get_active_votes                       | **DONE** | **DONE** |
| 60 | get_account_votes                      | **DONE** | **DONE** |
| 61 | get_content                            | **DONE** | **DONE** |
| 62 | get_content_replies                    | **DONE** | **DONE** |
| 63 | get_discussions_by_author_before_date  | **DONE** | **DONE** |
| 64 | get_replies_by_last_update             | **DONE** | **DONE** |
| 65 | get_witnesses                          | **DONE** | **DONE** |
| 66 | get_witness_by_account                 | **DONE** | **DONE** |
| 67 | get_witnesses_by_vote                  | **DONE** | **DONE** |
| 68 | lookup_witness_accounts                | **DONE** | **DONE** |
| 69 | get_witness_count                      | **DONE** | **DONE** |
| 70 | get_active_witnesses                   | **DONE** | **DONE** |
| 71 | get_miner_queue                        | **DONE** | **DONE** |
=======
### Subscriptions

**TODO:** Is this actually callable over the RPC endpoint?
It is a bit confusing to see `set_` prefix. Needs research.

```
   (set_subscribe_callback)
   (set_pending_transaction_callback)
   (set_block_applied_callback)
   (cancel_all_subscriptions)
```

### Tags

| Method Name                 | Raw Version | Full Version |
| --------------------------- |:-----------:|:------------:|
| get_trending_tags           | DONE        |              |
| get_discussions_by_trending | DONE        |              |
| get_discussions_by_created  | DONE        |              |
| get_discussions_by_active   | DONE        |              |
| get_discussions_by_cashout  | DONE        |              |
| get_discussions_by_payout   | DONE        |              |
| get_discussions_by_votes    | DONE        |              |
| get_discussions_by_children | DONE        |              |
| get_discussions_by_hot      | DONE        |              |
| get_recommended_for         | DONE        |              |

### Blocks and Transactions

| Method Name             | Raw Version | Full Version   |
| ----------------------- |:-----------:|:--------------:|
| get_block_header        | DONE        |                |
| get_block               | DONE        | PARTIALLY DONE |
| get_ops_in_block        | DONE        | DONE           |
| get_state               | DONE        |                |
| get_trending_categories | DONE        |                |
| get_best_categories     | DONE        |                |
| get_active_categories   | DONE        |                |
| get_recent_categories   | DONE        |                |
| get_ops_in_block        | DONE        |                |

### Globals

| Method Name                      | Raw Version | Full Version   |
| -------------------------------- |:-----------:|:--------------:|
| get_config                       | DONE        | PARTIALLY DONE |
| get_dynamic_global_properties    | DONE        | DONE           |
| get_chain_properties             | DONE        |                |
| get_feed_history                 | DONE        |                |
| get_current_median_history_price | DONE        |                |
| get_witness_schedule             | DONE        |                |
| get_hardfork_version             | DONE        | DONE           |
| get_next_scheduled_hardfork      | DONE        |                |

### Keys

| Method Name       | Raw Version | Full Version |
| ----------------- |:-----------:|:------------:|
| get_key_reference |             |              |

### Accounts

| Method Name               | Raw Version | Full Version |
| ------------------------- |:-----------:|:------------:|
| get_accounts              | DONE        |              |
| get_account_references    |             |              |
| lookup_account_names      | DONE        |              |
| lookup_accounts           | DONE        |              |
| get_account_count         | DONE        |              |
| get_conversation_requests | DONE        |              |
| get_account_history       | DONE        |              |

### Market

| Method Name     | Raw Version | Full Version |
| --------------- |:-----------:|:------------:|
| get_order_book  |             |              |
| get_open_orders |             |              |

### Authority / Validation

| Method Name              | Raw Version | Full Version |
| ------------------------ |:-----------:|:------------:|
| get_transaction_hex      |             |              |
| get_transaction          |             |              |
| get_required_signatures  |             |              |
| get_potential_signatures |             |              |
| verify_authority         |             |              |
| verity_account_authority |             |              |

### Votes

| Method Name       | Raw Version | Full Version |
| ----------------- |:-----------:|:------------:|
| get_active_votes  | DONE        | DONE         |
| get_account_votes | DONE        |              |

### Cotent

| Method Name                           | Raw Version | Full Version   |
| ------------------------------------- |:-----------:|:--------------:|
| get_content                           | DONE        | PARTIALLY DONE |
| get_content_replies                   | DONE        | PARTIALLY DONE |
| get_discussions_by_author_before_date |             |                |
| get_replies_by_last_update            | DONE        |                |

### Witnesses

| Method Name             | Raw Version | Full Version |
| ----------------------- |:-----------:|:------------:|
| get_witnesses           |             |              |
| get_witness_by_account  |             |              |
| get_witnesses_by_vote   |             |              |
| lookup_witness_accounts |             |              |
| get_witness_count       |             |              |
| get_active_witnesses    |             |              |
| get_miner_queue         |             |              |


## License

MIT, see the `LICENSE` file.
