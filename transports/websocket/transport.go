package websocket

import (
	"io"
	// Stdlib
	"context"
	"net"
	"time"

	// Vendor
	"github.com/asuleymanov/jsonrpc2"
	"github.com/asuleymanov/websocket"
	"github.com/pkg/errors"
	"gopkg.in/tomb.v2"
)

const (
	DefaultHandshakeTimeout      = 30 * time.Second
	DefaultWriteTimeout          = 10 * time.Second
	DefaultReadTimeout           = 20 * time.Second
	DefaultAutoReconnectMaxDelay = 1 * time.Minute

	InitialAutoReconnectDelay       = 1 * time.Second
	AutoReconnectBackoffCoefficient = 1.5
)

var netDialer net.Dialer

// Transport implements a CallCloser accessing the Golos RPC endpoint over WebSocket.
type Transport struct {
	// URLs as passed into the constructor.
	urls         []string
	nextURLIndex int
	currentURL   string

	// Options.
	handshakeTimeout time.Duration
	readTimeout      time.Duration
	writeTimeout     time.Duration

	autoReconnectEnabled  bool
	autoReconnectMaxDelay time.Duration

	monitorChan chan<- interface{}

	// The underlying JSON-RPC connection.
	connCh chan chan *jsonrpc2.Conn
	errCh  chan error

	t *tomb.Tomb
}

// Option represents an option that can be passed into the transport constructor.
type Option func(*Transport)

// SetDialTimeout can be used to set the timeout when establishing a new connection.
//
// This function is deprecated, please use SetHandshakeTimeout.
func SetDialTimeout(timeout time.Duration) Option {
	return SetHandshakeTimeout(timeout)
}

// SetHandshakeTimeout can be used to set the timeout for WebSocket handshake.
func SetHandshakeTimeout(timeout time.Duration) Option {
	return func(t *Transport) {
		t.handshakeTimeout = timeout
	}
}

// SetReadTimeout sets the connection read timeout.
// The timeout is implemented using net.Conn.SetReadDeadline.
func SetReadTimeout(timeout time.Duration) Option {
	return func(t *Transport) {
		t.readTimeout = timeout
	}
}

// SetWriteTimeout sets the connection read timeout.
// The timeout is implemented using net.Conn.SetWriteDeadline.
func SetWriteTimeout(timeout time.Duration) Option {
	return func(t *Transport) {
		t.writeTimeout = timeout
	}
}

// SetReadWriteTimeout sets the connection read and write timeout.
// The timeout is implemented using net.Conn.SetDeadline.
func SetReadWriteTimeout(timeout time.Duration) Option {
	return func(t *Transport) {
		t.readTimeout = timeout
		t.writeTimeout = timeout
	}
}

// SetAutoReconnectEnabled can be used to enable automatic reconnection to the RPC endpoint.
// Exponential backoff is used when the connection cannot be established repetitively.
//
// See SetAutoReconnectMaxDelay to set the maximum delay between the reconnection attempts.
func SetAutoReconnectEnabled(enabled bool) Option {
	return func(t *Transport) {
		t.autoReconnectEnabled = enabled
	}
}

// SetAutoReconnectMaxDelay can be used to set the maximum delay between the reconnection attempts.
//
// This option only takes effect when the auto-reconnect mode is enabled.
//
// The default value is 5 minutes.
func SetAutoReconnectMaxDelay(delay time.Duration) Option {
	return func(t *Transport) {
		t.autoReconnectMaxDelay = delay
	}
}

// SetMonitor can be used to set the monitoring channel that can be used to watch
// connection-related state changes.
//
// All channel send operations are happening synchronously, so not receiving messages
// from the channel will lead to the whole thing getting stuck completely.
//
// This option only takes effect when the auto-reconnect mode is enabled.
//
// The channel is closed when the transport is closed.
func SetMonitor(monitorChan chan<- interface{}) Option {
	return func(t *Transport) {
		t.monitorChan = monitorChan
	}
}

// NewTransport creates a new transport that connects to the given WebSocket URLs.
//
// It is possible to specify multiple WebSocket endpoint URLs.
// In case the transport is configured to reconnect automatically,
// the URL to connect to is rotated on every connect attempt using round-robin.
func NewTransport(urls []string, options ...Option) (*Transport, error) {
	// Prepare a transport instance.
	t := &Transport{
		urls:                  urls,
		handshakeTimeout:      DefaultHandshakeTimeout,
		readTimeout:           DefaultReadTimeout,
		writeTimeout:          DefaultWriteTimeout,
		autoReconnectMaxDelay: DefaultAutoReconnectMaxDelay,
		connCh:                make(chan chan *jsonrpc2.Conn),
		errCh:                 make(chan error),
		t:                     &tomb.Tomb{},
	}

	// Apply the options.
	for _, opt := range options {
		opt(t)
	}

	t.t.Go(t.dialer)

	// Return the new transport.
	return t, nil
}

// Call implements transports.CallCloser.
func (t *Transport) Call(method string, params, result interface{}) error {
	// Limit the request context with the tomb context.
	ctx := t.t.Context(nil)

Loop:
	for {
		// Request a connection.
		connCh := make(chan *jsonrpc2.Conn, 1)
		select {
		case t.connCh <- connCh:
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context closed")
		}

		// Receive the connection.
		conn := <-connCh

		// Perform the call.
		err := conn.Call(ctx, method, params, result)
		if err == nil {
			return nil
		}

		// In case this is a context error, return immediately.
		if err := ctx.Err(); err != nil {
			return errors.Wrap(err, "context closed")
		}

		// In case auto-reconnect is disabled, fail immediately.
		if !t.autoReconnectEnabled {
			return errors.Wrap(err, "call failed")
		}

		// In case this is a connection error, request a new connection.
		err = errors.Cause(err)
		if _, ok := err.(*websocket.CloseError); ok || err == io.ErrUnexpectedEOF {
			select {
			case t.errCh <- errors.Wrap(err, "WebSocket closed"):
				continue Loop
			case <-ctx.Done():
				return errors.Wrap(ctx.Err(), "context closed")
			}
		}

		// Some other error occurred, return it immediately.
		return errors.Wrap(err, "call failed")
	}
}

func (t *Transport) dialer() error {
	ctx := t.t.Context(nil)

	var conn *jsonrpc2.Conn
	defer func() {
		if conn != nil {
			conn.Close()
			err := errors.Wrap(ctx.Err(), "context closed")
			t.emit(&DisconnectedEvent{t.currentURL, err})
		}
	}()

	connect := func() {
		delay := InitialAutoReconnectDelay

		for {
			var err error
			conn, err = t.dial(ctx)
			if err == nil {
				break
			}

			select {
			case <-time.After(delay):
				delay = time.Duration(float64(delay) * AutoReconnectBackoffCoefficient)
				if delay > t.autoReconnectMaxDelay {
					delay = t.autoReconnectMaxDelay
				}

			case <-ctx.Done():
				return
			}
		}
	}

	// Establish the initial connection.
	connect()

	for {
		select {
		case connCh := <-t.connCh:
			connCh <- conn

		case err := <-t.errCh:
			conn.Close()
			t.emit(&DisconnectedEvent{t.currentURL, err})
			connect()

		case <-ctx.Done():
			return nil
		}
	}
}

func (t *Transport) dial(ctx context.Context) (*jsonrpc2.Conn, error) {
	// Set up a dialer.
	dialer := websocket.Dialer{
		NetDial: func(network, addr string) (net.Conn, error) {
			return netDialer.DialContext(ctx, network, addr)
		},
		HandshakeTimeout: t.handshakeTimeout,
	}

	// Get the next URL to try.
	u := t.urls[t.nextURLIndex]
	t.nextURLIndex = (t.nextURLIndex + 1) % len(t.urls)
	t.currentURL = u

	// Connect the WebSocket.
	t.emit(&ConnectingEvent{u})
	ws, _, err := dialer.Dial(u, nil)
	if err != nil {
		err = errors.Wrapf(err, "failed to dial %v", u)
		t.emit(&DisconnectedEvent{u, err})
		return nil, err
	}
	t.emit(&ConnectedEvent{u})

	// Wrap the WebSocket with JSON-RPC2.
	stream := NewObjectStream(ws, t.writeTimeout, t.readTimeout)
	return jsonrpc2.NewConn(ctx, stream, nil), nil
}

func (t *Transport) emit(v interface{}) {
	if t.monitorChan != nil {
		select {
		case t.monitorChan <- v:
		default:
		}
	}
}

// Close implements transports.CallCloser.
func (t *Transport) Close() error {
	t.t.Kill(nil)
	return t.t.Wait()
}
