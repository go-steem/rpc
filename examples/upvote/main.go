package main

import (
	// Stdlib
	"flag"
	"fmt"
	"log"

	// RPC
	"github.com/asuleymanov/golos-go"

	// Vendor
	"github.com/pkg/errors"
)

var (
	voter = ""
	key   = ""
)

func main() {
	cls, err = client.NewClient([]string{"wss://rpc.buildteam.io"}, "steem")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer cls.Close()

	cls.SetKeys(&client.Keys{PKey: []string{key}})

	if err := run(cls); err != nil {
		log.Fatalln("Error:", err)
	}
}

func run(cls *client.Client) (err error) {
	flag.Parse()
	// Process args.
	args := flag.Args()

	if len(args) != 2 {
		return errors.New("2 arguments required")
	}
	author, permlink := args[0], args[1]

	fmt.Println(cls.Vote(voter, author, permlink, 10000))

	return nil
}
