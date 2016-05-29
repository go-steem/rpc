# go-steem/rpc

Golang RPC client library for [Steem](https://steem.io).

## Compatibility

`steemd 0.5.0`

## Usage

1. `import "github.com/go-steem/rpc"`
2. Use `rpc.Dial` to get the RPC client.
3. PROFIT!

## Package Organisation

Once you create a `Client` object, you can start calling the methods exported
via `steemd`'s RPC endpoint by invoking associated methods on the client object.

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

As we can see, the difference is that the raw version returns`*json.RawMessage`,
 so it is not trying to unmarshall the response into a properly typed response.

There are two reasons for this:

1. To be able to see raw data.
2. To be able to call most of the remote methods even though the response
   object is not yet known or specified.

## Status

This package is still under rapid development and it is by no means complete.
For now there is no promise considering API stability.

The following table documents the API completion:


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
					case *rpc.CommentOperations:
						content, _ := client.GetContent(body.Author, body.Permlink)
						fmt.Printf("COMMENT @%v %v\n", content.Author, content.URL)

					// Vote operation.
					case *rpc.VoteOperation:
						fmt.Printf("@%v voted for @%v/%v\n", body.Voter, body.Author, body.Permlink)

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

## License

MIT, see the `LICENSE` file.
