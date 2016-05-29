package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/go-steem/rpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	// Process flags.
	flagAddress := flag.String("rpc_endpoint", "ws://localhost:8090", "steemd RPC endpoint address")
	flag.Parse()

	// Connect to the RPC endpoint.
	addr := *flagAddress
	log.Printf("---> Dial(\"%v\")\n", addr)
	client, err := rpc.Dial(addr)
	if err != nil {
		return err
	}
	defer client.Close()

	// Get config.
	log.Println("---> GetConfig()")
	config, err := client.GetConfig()
	if err != nil {
		return err
	}

	// Process blocks as they come.
	lastBlock := props.LastIrreversibleBlockNum
	log.Printf("---> Entering the block processing loop (last block = %v)\n", lastBlock)
	for {
		// Get current properties.
		log.Println("---> GetDynamicGlobalProperties()")
		props, err := client.GetDynamicGlobalProperties()
		if err != nil {
			return err
		}

		// Process new blocks.
		for props.LastIrreversibleBlockNum-lastBlock > 0 {
			log.Printf("---> GetBlock(%v)\n", lastBlock)
			block, err := client.GetBlock(lastBlock)
			if err != nil {
				return err
			}

			// Process the transactions.
			for _, tx := range block.Transactions {
				for _, op := range tx.Operations {
					switch body := op.Body.(type) {
					case *rpc.VoteOperation:
						fmt.Printf("@%v voted for @%v/%v\n", body.Voter, body.Author, body.Permlink)

						// You can add more cases here, it depends on what
						// operations you actually need to process.
					}
				}
			}

			lastBlock++
		}

		// Sleep for STEEMIT_BLOCK_INTERVAL seconds before the next iteration.
		time.Sleep(time.Duration(config.SteemitBlockInterval) * time.Second)
	}
}
