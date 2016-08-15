package rpc

import (
	// RPC
	"github.com/go-steem/rpc/apis/database"
	"github.com/go-steem/rpc/apis/login"
	"github.com/go-steem/rpc/interfaces"
)

// Client can be used to access Steem remote APIs.
//
// There is a public field for every Steem API available,
// e.g. Client.Database corresponds to database_api.
type Client struct {
	cc interfaces.CallCloser

	// Login represents login_api.
	Login *login.API

	// Database represents database_api.
	Database *database.API
}

// NewClient creates a new RPC client that use the given CallCloser internally.
func NewClient(cc interfaces.CallCloser) *Client {
	client := &Client{cc: cc}
	client.Login = login.NewAPI(client.cc)
	client.Database = database.NewAPI(client.cc)
	return client
}

// Close should be used to close the client when no longer needed.
// It simply calls Close() on the underlying CallCloser.
func (client *Client) Close() error {
	return client.cc.Close()
}
