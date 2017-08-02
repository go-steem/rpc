package websocket

import (
	// Stdlib
	"io"
	"net"
	"time"

	// Vendor
	"github.com/asuleymanov/rpc-codec/jsonrpc2"
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
	"gopkg.in/tomb.v2"
)

type callRequest struct {
	method   string
	params   interface{}
	response interface{}
	errCh    chan<- error
}

type reconnectingTransport struct {
	parent *Transport

	ws        *websocket.Conn
	client    *jsonrpc2.Client
	requestCh chan *callRequest

	t *tomb.Tomb
}

func newReconnectingTransport(parent *Transport) *reconnectingTransport {
	cc := &reconnectingTransport{
		parent:    parent,
		requestCh: make(chan *callRequest),
		t:         &tomb.Tomb{},
	}
	cc.t.Go(cc.worker)
	return cc
}

// Call implements interfaces.CallCloser.
func (t *reconnectingTransport) Call(method string, params, response interface{}) error {
	errCh := make(chan error, 1)
	select {
	case t.requestCh <- &callRequest{method, params, response, errCh}:
		return <-errCh
	case <-t.t.Dying():
		return ErrClosing
	}
}

// Close implements interfaces.CallCloser.
func (t *reconnectingTransport) Close() error {
	t.t.Kill(nil)
	return t.t.Wait()
}

func (t *reconnectingTransport) worker() error {
	// Close the monitoring channel when returning.
	if ch := t.parent.monitorChan; ch != nil {
		defer func() {
			close(ch)
		}()
	}

	// Keep processing incoming call requests until interrupted.
	for {
		select {
		case req := <-t.requestCh:
			req.errCh <- t.handleCall(req.method, req.params, req.response)

		case <-t.t.Dying():
			return nil
		}
	}
}

func (t *reconnectingTransport) handleCall(method string, params, response interface{}) error {
	for {
		// Get an RPC client. This blocks until the client is available or Close() is called.
		client, err := t.getClient()
		if err != nil {
			return err
		}

		// Update the connection timeout if necessary.
		if err := t.parent.updateDeadline(t.ws); err != nil {
			return err
		}

		// Perform the call.
		if err := client.Call(method, params, response); err != nil {
			// In case there is a network error, we retry immediately.
			if err, ok := asNetworkError(err); ok {
				t.dropClient(err)
				continue
			}
			// The connection can be also closed unexpectedly.
			// That counts as a network error for us as well.
			if err == io.ErrUnexpectedEOF {
				t.dropClient(err)
				continue
			}
			// Otherwise we just return the error.
			return err
		}

		// Done.
		return nil
	}
}

func (t *reconnectingTransport) getClient() (*jsonrpc2.Client, error) {
	// In case the client is not set, establish a new connection.
	if t.client == nil {
		ws, err := t.connect()
		if err != nil {
			return nil, err
		}
		t.ws = ws
		t.client = jsonrpc2.NewClient(ws)
	}

	// Return the cached client.
	return t.client, nil
}

func (t *reconnectingTransport) dropClient(err error) {
	if t.client != nil {
		// Close and drop the client.
		t.client.Close()
		t.client = nil
		t.ws = nil

		// Emit DISCONNECTED.
		t.emitEvent(&DisconnectedEvent{
			URL: t.parent.url.String(),
			Err: err,
		})
	}
}

func (t *reconnectingTransport) connect() (*websocket.Conn, error) {
	// Get a new client. Keep trying to establish a new connection using exponential backoff.
	timeout := 1 * time.Second
	wait := func() error {
		// Wait for the given period.
		select {
		case <-time.After(timeout):
		case <-t.t.Dying():
			return ErrClosing
		}

		// Update the timeout value.
		timeout = 2 * timeout
		if timeout > t.parent.autoReconnectMaxDelay {
			timeout = t.parent.autoReconnectMaxDelay
		}
		return nil
	}

	urlString := t.parent.url.String()

	for {
		// Emit CONNECTING.
		t.emitEvent(&ConnectingEvent{urlString})

		// Try to establish a new WebSocket connection.
		ws, err := t.parent.dial(t.t.Dying())
		if err != nil {
			// Handle network errors.
			if err, ok := asNetworkError(err); ok {
				if err.Timeout() {
					// Emit DIAL_TIMEOUT.
					t.emitEvent(&DialTimeoutEvent{
						URL:     urlString,
						Err:     err,
						Timeout: timeout,
					})
				} else {
					// Emit DISCONNECTED.
					t.emitEvent(&DisconnectedEvent{
						URL: urlString,
						Err: err,
					})
				}

				// Wait for the given period.
				if err := wait(); err != nil {
					return nil, err
				}
				// Try again.
				continue
			}

			// Otherwise just return the error.
			return nil, err
		}

		// Connection established.
		// Emit CONNECTED and return a new client.
		t.emitEvent(&ConnectedEvent{urlString})
		return ws, nil
	}
}

func (t *reconnectingTransport) emitEvent(event interface{}) {
	if ch := t.parent.monitorChan; ch != nil {
		ch <- event
	}
}

func asNetworkError(err error) (opError *net.OpError, ok bool) {
	opError, ok = errors.Cause(err).(*net.OpError)
	return
}
