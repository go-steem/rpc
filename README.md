# go-steem/rpc

[![GoDoc](https://godoc.org/github.com/go-steem/rpc?status.svg)](https://godoc.org/github.com/go-steem/rpc)

Golang RPC client library for [Steem](https://steem.io).

## Compatibility

`steemd 0.5.0`

## Usage

1. `import "github.com/go-steem/rpc"`
2. Use `rpc.Dial` to get the RPC client.
3. PROFIT!

## Example

This is just a code snippet. Please check the `examples` directory
for more complete and ready to use examples.

```go
// Instantiate a new client.
client, _ := rpc.Dial("ws://localhost:8090")
defer client.Close()

// Call "get_config".
config, _ := client.GetConfig()

// Start processing blocks.
lastBlock := 1800000
for {
	// Call "get_dynamic_global_properties".
	props, _ := client.GetDynamicGlobalProperties()

	for props.LastIrreversibleBlockNum-lastBlock > 0 {
		// Call "get_block".
		block, _ := client.GetBlock(lastBlock)

		// Process the transactions.
		for _, tx := range block.Transactions {
			for _, op := range tx.Operations {
				switch body := op.Body.(type) {
					// Comment operation.
					case *rpc.CommentOperation:
						content, _ := client.GetContent(body.Author, body.Permlink)
						fmt.Printf("COMMENT @%v %v\n", content.Author, content.URL)

					// Vote operation.
					case *rpc.VoteOperation:
						fmt.Printf("VOTE @%v @%v/%v\n", body.Voter, body.Author, body.Permlink)

					// You can add more cases, it depends on what
					// operations you actually need to process.
				}
			}
		}

		lastBlock++
	}

	time.Sleep(time.Duration(config.SteemitBlockInterval) * time.Second)
}
```

## Package Organisation

Once you create a `Client` object, you can start calling the methods exported
via `steemd`'s RPC endpoint by invoking associated methods on the client object.
It is enough to just turn the method names into CamelCase, e.g. `get_config`
becomes `Client.GetConfig`.

There are two methods implemented for the `Client` object for every
method exported via the RPC endpoint. The regular version and the raw version.
Let's see an example for `get_config`:

```go
func (client *Client) GetConfig() (*Config, error) {
	...
}

func (client *Client) GetConfigRaw() (*json.RawMessage, error) {
	...
}
```

As we can see, the difference is that the raw version returns `*json.RawMessage`,
 so it is not trying to unmarshall the response into a properly typed response.

There are two reasons for this:

1. To be able to see raw data.
2. To be able to call most of the remote methods even though the response
   object is not yet known or specified.

It is already benefitial to just have the raw version because at least
the method parameters are statically typed.

## Status

This package is still under rapid development and it is by no means complete.
For now there is no promise considering API stability. Some response objects
maybe be typed incorrectly. The package is already usable, though. See the
`examples` directory.

The following subsections document the API completion. The method names
are taken from `database_api.hpp` in `steemit/steem`.

**TODO:** Verify that these are actually the methods available over RPC :D

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
| get_account_refenreces    |             |              |
| lookup_account_names      | DONE        |              |
| lookup_accounts           | DONE        |              |
| get_account_count         | DONE        |              |
| get_conversation_requests | DONE        |              |
| get_account_history       | DONE        |              |

### Market

| Method Name    | Raw Version | Full Version |
| -------------- |:-----------:|:------------:|
| get_order_book |             |              |

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
| get_active_votes  | DONE        |              |
| get_account_votes | DONE        |              |

### Cotent

| Method Name                           | Raw Version | Full Version   |
| ------------------------------------- |:-----------:|:--------------:|
| get_content                           | DONE        | PARTIALLY DONE |
| get_content_replies                   | DONE        |                |
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
