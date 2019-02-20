package client

import (
	"net/url"

	"github.com/asuleymanov/steem-go/api"
	"github.com/asuleymanov/steem-go/transports"
	"github.com/asuleymanov/steem-go/transports/http"
	"github.com/asuleymanov/steem-go/transports/websocket"
	"github.com/pkg/errors"
)

var (
	ErrInitializeTransport = errors.New("Failed to initialize transport.")
)

// Client can be used to access STEEM remote APIs.
// There is a public field for every STEEM API available,
// e.g. Client.Database corresponds to database_api.
type Client struct {
	cc transports.CallCloser

	chainID string

	AsyncProtocol bool

	// Database represents database_api.
	API *api.API

	// Current keys for operations
	CurrentKeys *Keys
}

// NewClient creates a new RPC client that use the given CallCloser internally.
// Initialize only server present API. Absent API initialized as nil value.
func NewClient(s string) (*Client, error) {
	// Parse URL
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	// Initializing Transport
	var call transports.CallCloser
	switch u.Scheme {
	case "wss", "ws":
		call, err = websocket.NewTransport(s)
		if err != nil {
			return nil, err
		}
	case "https", "http":
		call, err = http.NewTransport(s)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrInitializeTransport
	}
	client := &Client{cc: call}

	client.AsyncProtocol = false

	client.API = api.NewAPI(client.cc)

	chainID, err := client.API.GetConfig()
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
