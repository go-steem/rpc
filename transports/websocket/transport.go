package websocket

import (
	"github.com/go-steem/rpc-codec/jsonrpc2"
	"golang.org/x/net/websocket"
)

type Transport struct {
	client *jsonrpc2.Client
}

func Dial(address string) (*Transport, error) {
	// Connect to the given WebSocket URL.
	conn, err := websocket.Dial(address, "", "http://localhost")
	if err != nil {
		return nil, err
	}

	// Instantiate a JSON-RPC client.
	client := jsonrpc2.NewClient(conn)

	// Return the transport.
	return &Transport{client}, nil
}

func (t *Transport) Call(method string, params, response interface{}) error {
	return t.client.Call(method, params, response)
}

func (t *Transport) Close() error {
	return t.client.Close()
}
