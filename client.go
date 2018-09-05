package client

import (
	"errors"
	
	//"github.com/asuleymanov/steem-go/apis/condenser"
	"github.com/asuleymanov/steem-go/apis/database"
	"github.com/asuleymanov/steem-go/apis/follow"
	"github.com/asuleymanov/steem-go/apis/market"
	"github.com/asuleymanov/steem-go/apis/networkbroadcast"
	"github.com/asuleymanov/steem-go/transactions"
	"github.com/asuleymanov/steem-go/transports"
	"github.com/asuleymanov/steem-go/transports/websocket"
)

// Client can be used to access Steem remote APIs.
//
// There is a public field for every Steem API available,
// e.g. Client.Database corresponds to database_api.
type Client struct {
	cc transports.CallCloser

	// Database represents database_api.
	Database *database.API

	// Condenser represents condenser_api.
	//Condenser *condenser.API

	// Follow represents follow_api.
	Follow *follow.API

	// Follow represents market_history_api.
	Market *market.API

	// NetworkBroadcast represents network_broadcast_api.
	NetworkBroadcast *networkbroadcast.API

	//Chain Id
	Chain *transactions.Chain

	// Current keys for operations
	CurrentKeys *Keys
}

// NewClient creates a new RPC client that use the given CallCloser internally.
func NewClient(url []string, chain string) (*Client, error) {
	call, err := initclient(url)
	if err != nil {
		return nil, err
	}
	client := &Client{cc: call}

	client.Database = database.NewAPI(client.cc)

	//client.Condenser = condenser.NewAPI(client.cc)

	client.Follow = follow.NewAPI(client.cc)

	client.Market = market.NewAPI(client.cc)

	client.NetworkBroadcast = networkbroadcast.NewAPI(client.cc)

	client.Chain, err = initChainID(chain)
	if err != nil {
		client.Chain = transactions.SteemChain
	}

	return client, nil
}

// Close should be used to close the client when no longer needed.
// It simply calls Close() on the underlying CallCloser.
func (client *Client) Close() error {
	return client.cc.Close()
}

//SetKeys you can specify keys for signing transactions.
func (client *Client) SetKeys(keys *Keys) {
	client.CurrentKeys = keys
}

func initclient(url []string) (*websocket.Transport, error) {
	// Инициализация Websocket
	t, err := websocket.NewTransport(url)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func initChainID(str string) (*transactions.Chain, error) {
	var ChainID transactions.Chain
	// Определяем ChainId
	switch str {
	case "steem":
		ChainID = *transactions.SteemChain
	case "test":
		ChainID = *transactions.TestChain
	default:
		return nil, errors.New("Chain not found")
	}
	return &ChainID, nil
}
