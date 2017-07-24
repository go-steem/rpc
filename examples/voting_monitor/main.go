package main

import (
	"log"
	"time"

	"github.com/asuleymanov/golos-go/client"
	"github.com/asuleymanov/golos-go/types"
)

var cls = client.NewApi()

func main() {
	defer cls.Rpc.Close()
	if err := run(); err != nil {
		log.Fatalln("Error:", err)
	}
}

func run() (err error) {
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
}
