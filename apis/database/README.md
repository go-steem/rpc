# Database API

This package adds support for `database_api`.

## State

The following subsections document the API completion. The method names
are taken from `database_api.hpp` in `steemit/steem`.

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
| get_state               | DONE        |                |
| get_trending_categories | DONE        |                |
| get_best_categories     | DONE        |                |
| get_active_categories   | DONE        |                |
| get_recent_categories   | DONE        |                |

### Globals

| Method Name                      | Raw Version | Full Version   |
| -------------------------------- |:-----------:|:--------------:|
| get_config                       | DONE        | DONE           |
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
