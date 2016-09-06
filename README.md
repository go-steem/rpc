# go-steem/rpc

[![GoDoc](https://godoc.org/github.com/go-steem/rpc?status.svg)](https://godoc.org/github.com/go-steem/rpc)

Golang RPC client library for [Steem](https://steem.io).

## Compatibility

`steemd 0.13.0`

## Usage

```go
import "github.com/go-steem/rpc"
```

This package is still very much in development, so `gopkg.in` is not yet available.

## Installation

This package calls [bitcoin-core/secp256k1](https://github.com/bitcoin-core/secp256k1)
using CGO to implement signed transactions, so you need to install `secp256k1` first.
Then it will be possible to build `go-steem/rpc`.

In case you don't need signed transactions, i.e. you don't need to use
`network_broadcast_api`, it is possible to build the package with `nosigning`
tag to exclude the functionality:

```bash
$ go build -tags nosigning
```

## Example

This is just a code snippet. Please check the `examples` directory
for more complete and ready to use examples.

```go
// Instantiate the WebSocket transport.
t, _ := websocket.NewTransport("ws://localhost:8090")

// Use the transport to create an RPC client.
client, _ := rpc.NewClient(t)
defer client.Close()

// Call "get_config".
config, _ := client.Database.GetConfig()

// Start processing blocks.
lastBlock := 1800000
for {
	// Call "get_dynamic_global_properties".
	props, _ := client.Database.GetDynamicGlobalProperties()

	for props.LastIrreversibleBlockNum-lastBlock > 0 {
		// Call "get_block".
		block, _ := client.Database.GetBlock(lastBlock)

		// Process the transactions.
		for _, tx := range block.Transactions {
			for _, op := range tx.Operations {
				switch body := op.Data().(type) {
					// Comment operation.
					case *types.CommentOperation:
						content, _ := client.Database.GetContent(body.Author, body.Permlink)
						fmt.Printf("COMMENT @%v %v\n", content.Author, content.URL)

					// Vote operation.
					case *types.VoteOperation:
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

You need to create a `Client` object to be able to do anything. To be able to
instantiate a `Client`, you first need to create a transport to be used to
execute RPC calls. The WebSocket transport is available in `transports/websocket`.
Then you just need to call `NewClient(transport)`.

Once you create a `Client` object, you can start calling the methods exported
via `steemd`'s RPC endpoint by invoking associated methods on the client object.
There are multiple APIs that can be exported, e.g. `database_api` and `login_api`,
so the methods on the Client object are also namespaced accoding to these APIs.
For example, to call `get_block` from `database_api`, you need to use
`Client.Database.GetBlock` method.

When looking for a method to call, all you need is to turn the method name into
CamelCase, e.g. `get_config` becomes `Client.Database.GetConfig`.

### Raw and Full Methods

There are two methods implemented for every method exported via the RPC endpoint.
The regular version and the raw version. Let's see an example for `get_config`:

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

To check the API coverate, please check the README files in the relevat API
package in the `apis` subdirectory.

## License

MIT, see the `LICENSE` file.
