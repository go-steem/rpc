package main

import (
	// Stdlib
	"flag"
	"fmt"
	"log"

	// RPC
	"github.com/asuleymanov/golos-go/client"

	// Vendor
	"github.com/pkg/errors"
)

var cls = client.NewApi()

func main() {
	defer cls.Rpc.Close()
	if err := run(); err != nil {
		log.Fatalln("Error:", err)
	}
}

func run() (err error) {
	flag.Parse()
	// Process args.
	args := flag.Args()
	if len(args) != 2 {
		return errors.New("2 arguments required")
	}
	author, permlink := args[0], args[1]

	fmt.Println(cls.Vote(author, permlink, 10000))

	return nil
}
