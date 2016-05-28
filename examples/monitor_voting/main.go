package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/go-steem/rpc"
)

const DefaultNextBlock = 1800000

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	// Process flags.
	flagAddress := flag.String("rpc_endpoint", "ws://localhost:8090", "steemd RPC endpoint address")
	flagNextBlock := flag.Uint("next_block", DefaultNextBlock, "the block where to start processing")
	flag.Parse()

	// Connect to the RPC endpoint.
	addr := *flagAddress
	log.Printf("---> Dial(%v)\n", addr)
	client, err := rpc.Dial(addr)
	if err != nil {
		return err
	}
	defer client.Close()

	// Get config.
	log.Println("---> GetConfig")
	config, err := client.GetConfig()
	if err != nil {
		return err
	}

	// Loop.
	nextBlock := uint32(*flagNextBlock)
	log.Printf("---> Entering the block processing loop (starting with block %v)\n", nextBlock)
	for {
		log.Println("---> GetDynamicGlobalProperties")
		props, err := client.GetDynamicGlobalProperties()
		if err != nil {
			return err
		}

		for props.LastIrreversibleBlockNum-nextBlock >= 0 {
			log.Printf("---> GetBlock(%v)\n", nextBlock)
			block, err := client.GetBlock(nextBlock)
			if err != nil {
				return err
			}

			// Process the transactions.
			for _, tx := range block.Transactions {
				for _, op := range tx.Operations {
					switch body := op.Body.(type) {
					// Vote operation.
					case *rpc.VoteOperation:
						fmt.Printf("@%v voted for @%v/%v\n", body.Voter, body.Author, body.Permlink)

						// You can add more cases, it depends on what
						// operations you actually need to process.
					}
				}
			}

			nextBlock++
		}

		time.Sleep(time.Duration(config.SteemitBlockInterval) * time.Second)
	}
}
