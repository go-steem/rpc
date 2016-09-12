package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"encoding/json"
	"os/signal"
	"syscall"
	"time"
        r "gopkg.in/dancannon/gorethink.v2"
	"github.com/go-steem/rpc"
	"github.com/go-steem/rpc/transports/websocket"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("Error:", err)
	}
}

// Set the settings for the DB
func run() (err error) {
    Rsession, err := r.Connect(r.ConnectOpts{
    Addresses: []string{"192.168.194.98:28015","192.168.194.91:28015","192.168.194.145:28015","192.168.194.209:28015"},
    })
    if err != nil {
        log.Fatalln(err.Error())
    }

// Create a table in the DB
var rethinkdbname string = "steemit"
_, err = r.DBCreate(rethinkdbname).RunWrite(Rsession)
Rsession.Use(rethinkdbname)
	if err != nil {
		fmt.Println("rethindb DB already made")
}
	_, err = r.DB(rethinkdbname).TableCreate("blocks").RunWrite(Rsession)

	if err != nil {
		fmt.Println("Probably already made a table for blocks")

	}


	// Process flags.
	flagAddress := flag.String("rpc_endpoint", "ws://138.201.198.167:8090", "steemd RPC endpoint address")
	flagReconnect := flag.Bool("reconnect", false, "enable auto-reconnect mode")
	flag.Parse()

	var (
		url       = *flagAddress
		reconnect = *flagReconnect
	)

	// Start catching signals.
	var interrupted bool
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Drop the error in case it is a request being interrupted.
	defer func() {
		if err == websocket.ErrClosing && interrupted {
			err = nil
		}
	}()
	// This allows you to tell the app which block to start on.
	// TODO: Make all of the vars into a config file and package the binaries
	Startblock := 3100000
	U := uint32(Startblock)
	// Start the connection monitor.
	monitorChan := make(chan interface{}, 1)
	if reconnect {
		go func() {
			for {
				event, ok := <-monitorChan
				if ok {
					log.Println(event)
				}
			}
		}()
	}

	// Instantiate the WebSocket transport.
	log.Printf("---> Dial(\"%v\")\n", url)
	t, err := websocket.NewTransport(url,
		websocket.SetAutoReconnectEnabled(reconnect),
		websocket.SetAutoReconnectMaxDelay(30*time.Second),
		websocket.SetMonitor(monitorChan))
	if err != nil {
		return err
	}

	// Use the transport to get an RPC client.
	client, err := rpc.NewClient(t)
	if err != nil {
		return err
	}
	defer func() {
		if !interrupted {
			client.Close()
		}
	}()

	// Start processing signals.
	go func() {
		<-signalCh
		fmt.Println()
		log.Println("Signal received, exiting...")
		signal.Stop(signalCh)
		interrupted = true
		client.Close()
	}()

	// Get config.
	log.Println("---> GetConfig()")
	config, err := client.Database.GetConfig()
	if err != nil {
		return err
	}


	// Keep processing incoming blocks forever.
	fmt.Println("---> Entering the block processing loop")
	for {
		// Get current properties.
		props, err := client.Database.GetDynamicGlobalProperties()
		if err != nil {
			return err
		}

		// Process new blocks.
		for props.LastIrreversibleBlockNum-U > 0 {
			block, err := client.Database.GetBlockRaw(U)
			lastblock := props.LastIrreversibleBlockNum
			var f interface{}
			json.Unmarshal(*block, &f)
			fmt.Println(U)
			fmt.Println(f)
			r.Table("blocks").
			Insert(f).
			Exec(Rsession)
			if err != nil {
				return err
			}

			// Process the transactions.
			if U != lastblock {
				U++
			}


}
time.Sleep(time.Duration(config.SteemitBlockInterval) * time.Second)
		}

	}
