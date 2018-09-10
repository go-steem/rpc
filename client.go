package client

import (
	"github.com/asuleymanov/steem-go/api/database"
	"github.com/asuleymanov/steem-go/api/follow"
	"github.com/asuleymanov/steem-go/api/market_history"
	"github.com/asuleymanov/steem-go/api/network_broadcast"
	"github.com/asuleymanov/steem-go/transports"
	"github.com/asuleymanov/steem-go/transports/websocket"
	"github.com/asuleymanov/steem-go/types"
)

// Client can be used to access Steem remote APIs.
// There is a public field for every Steem API available,
// e.g. Client.Database corresponds to database_api.
type Client struct {
	cc transports.CallCloser

	chainID string

	AsyncProtocol bool

	// Fixed JSONMetadata added to posting all comments
	DefaultContentMetadata types.ContentMetadata

	// Database represents database_api.
	Database *database.API

	// Follow represents follow_api.
	Follow *follow.API

	// Follow represents market_history_api.
	MarketHistory *market_history.API

	// NetworkBroadcast represents network_broadcast_api.
	NetworkBroadcast *network_broadcast.API

	// Current keys for operations
	CurrentKeys *Keys
}

// NewClient creates a new RPC client that use the given CallCloser internally.
// Initialize only server present API. Absent API initialized as nil value.
func NewClient(url []string, options ...websocket.Option) (*Client, error) {
	call, err := initClient(url, options...)
	if err != nil {
		return nil, err
	}
	client := &Client{cc: call}

	client.AsyncProtocol = false

	client.Database = database.NewAPI(client.cc)

	client.Follow = follow.NewAPI(client.cc)

	client.MarketHistory = market_history.NewAPI(client.cc)

	client.NetworkBroadcast = network_broadcast.NewAPI(client.cc)

	chainID, err := client.Database.GetConfig()
	if err != nil {
		return nil, err
	}
	client.chainID = chainID.ChainID

	return client, nil
}

// Close should be used to close the client when no longer needed.
// It simply calls Close() on the underlying CallCloser.
func (client *Client) Close() error {
	return client.cc.Close()
}

func initClient(url []string, options ...websocket.Option) (*websocket.Transport, error) {
	// Initializing Websocket
	t, err := websocket.NewTransport(url, options...)
	if err != nil {
		return nil, err
	}

	return t, nil
}

//GenCommentMetadata generate default CommentMetadata
func (client *Client) GenCommentMetadata(meta *types.ContentMetadata) *types.ContentMetadata {
	if client.DefaultContentMetadata != nil {
		for k := range client.DefaultContentMetadata {
			_, ok := (*meta)[k]
			if !ok {
				// Set fixed value only if value not exists
				(*meta)[k] = client.DefaultContentMetadata[k]
			}
		}
	}
	return meta
}
