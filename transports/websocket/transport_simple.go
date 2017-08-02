package websocket

import (
	// Vendor
	"github.com/asuleymanov/rpc-codec/jsonrpc2"
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
)

// simpleTransport is not trying to be particularly clever about network errors.
// In case an error occurs, it is immediately returned and the transport is closed.
type simpleTransport struct {
	parent *Transport

	ws     *websocket.Conn
	client *jsonrpc2.Client
}

// newSimpleTransport establishes a new WebSocket connection.
// The function blocks until the process is finished.
func newSimpleTransport(parent *Transport) (*simpleTransport, error) {
	// Establish the WebSocket connection.
	ws, err := parent.dial(nil)
	if err != nil {
		return nil, err
	}

	// Instantiate a JSON-RPC client.
	client := jsonrpc2.NewClient(ws)

	// Return a new simple transport.
	return &simpleTransport{parent, ws, client}, nil
}

// Call implements interfaces.CallCloser.
func (t *simpleTransport) Call(method string, params, response interface{}) error {
	if err := t.parent.updateDeadline(t.ws); err != nil {
		return err
	}
	if err := t.client.Call(method, params, response); err != nil {
		return errors.Wrapf(err, "failed to call %v(%v)", method, params)
	}
	return nil
}

// Close implements interfaces.CallCloser.
func (t *simpleTransport) Close() error {
	if err := t.client.Close(); err != nil {
		return errors.Wrap(err, "failed to close JSON-RPC client")
	}
	return nil
}
