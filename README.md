# asuleymanov/golos-go

[![GoDoc](https://godoc.org/github.com/asuleymanov/golos-go?status.svg)](https://godoc.org/github.com/asuleymanov/golos-go)

Golang RPC client library for [Golos](https://golos.io).

## Compatibility

`golosd 0.16.0`
`steemd 0.18.0`

## Usage

```go
import "github.com/asuleymanov/golos-go/client"
```


## Example

This is just a code snippet. Please check the `examples` directory
for more complete and ready to use examples.

```go
	cls := client.NewApi([]string{"wss://ws.golos.io"},"golos")
	defer cls.Rpc.Close()
	
	// Get config.
	log.Println("---> GetConfig()")
	config, err := cls.Rpc.Database.GetConfig()
	if err != nil {
		return err
	}

	// Use the last irreversible block number as the initial last block number.
	props, err := cls.Rpc.Database.GetDynamicGlobalProperties()
	if err != nil {
		return err
	}
	lastBlock := props.LastIrreversibleBlockNum

	// Keep processing incoming blocks forever.
	log.Printf("---> Entering the block processing loop (last block = %v)\n", lastBlock)
	for {
		// Get current properties.
		props, err := cls.Rpc.Database.GetDynamicGlobalProperties()
		if err != nil {
			return err
		}

		// Process new blocks.
		for props.LastIrreversibleBlockNum-lastBlock > 0 {
			block, err := cls.Rpc.Database.GetBlock(lastBlock)
			if err != nil {
				return err
			}

			// Process the transactions.
			for _, tx := range block.Transactions {
				for _, operation := range tx.Operations {
					switch op := operation.Data().(type) {
					case *types.VoteOperation:
						log.Printf("@%v voted for @%v/%v\n", op.Voter, op.Author, op.Permlink)

						// You can add more cases here, it depends on
						// what operations you actually need to process.
					}
				}
			}

			lastBlock++
		}

		// Sleep for STEEMIT_BLOCK_INTERVAL seconds before the next iteration.
		time.Sleep(time.Duration(config.SteemitBlockInterval) * time.Second)
	}
```

## Package Organisation


You need to create a `Client` object to be able to do anything.
Then you just need to call `NewApi()`.

Once you create a `Client` object, you can start calling the methods exported
via `golosd(steemd)`'s RPC endpoint by invoking associated methods on the client object.
There are multiple APIs that can be exported, e.g. `database_api` and `login_api`,
so the methods on the Client object are also namespaced accoding to these APIs.
For example, to call `get_block` from `database_api`, you need to use
`Client.Rpc.Database.GetBlock` method.

When looking for a method to call, all you need is to turn the method name into
CamelCase, e.g. `get_config` becomes `Client.Rpc.Database.GetConfig`.

## Status

This package is still under rapid development and it is by no means complete.
For now there is no promise considering API stability. Some response objects
maybe be typed incorrectly. The package is already usable, though. See the
`examples` directory.

To check the API coverate, please check the README files in the relevat API
package in the `apis` subdirectory.

## License

MIT, see the `LICENSE` file.
