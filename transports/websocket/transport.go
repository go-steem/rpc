package websocket

import (
	// Stdlib
	"crypto/tls"
	"net"
	"net/url"
	"time"

	// RPC
	"github.com/asuleymanov/golos-go/interfaces"

	// Vendor
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
)

const (
	DefaultDialTimeout           = 30 * time.Second
	DefaultAutoReconnectMaxDelay = 5 * time.Minute
)

// Transport implements a CallCloser accessing the Steem RPC endpoint over WebSocket.
type Transport struct {
	// URL as passed into the constructor.
	url *url.URL

	// Options.
	dialTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration

	autoReconnectEnabled  bool
	autoReconnectMaxDelay time.Duration

	monitorChan chan<- interface{}

	// Underlying CallCloser.
	cc interfaces.CallCloser
}

// Option represents an option that can be passed into the transport constructor.
type Option func(*Transport)

// SetDialTimeout can be used to set the timeout when establishing a new connection.
func SetDialTimeout(timeout time.Duration) Option {
	return func(t *Transport) {
		t.dialTimeout = timeout
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

// NewTransport creates a new transport that connects to the given WebSocket URL.
func NewTransport(endpointURL string, options ...Option) (*Transport, error) {
	// Parse the URL.
	epURL, err := url.Parse(endpointURL)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint URL")
	}

	// Prepare a transport instance.
	t := &Transport{
		url:                   epURL,
		dialTimeout:           DefaultDialTimeout,
		autoReconnectMaxDelay: DefaultAutoReconnectMaxDelay,
	}

	// Apply the options.
	for _, opt := range options {
		opt(t)
	}

	// Instantiate the underlying CallCloser based on the options.
	var cc interfaces.CallCloser
	if t.autoReconnectEnabled {
		cc = newReconnectingTransport(t)
	} else {
		cc, err = newSimpleTransport(t)
	}
	if err != nil {
		return nil, err
	}
	t.cc = cc

	// Return the new transport.
	return t, nil
}

// Call implements interfaces.CallCloser.
func (t *Transport) Call(method string, params, response interface{}) error {
	return t.cc.Call(method, params, response)
}

// Close implements interfaces.CallCloser.
func (t *Transport) Close() error {
	return t.cc.Close()
}

// dial establishes a WebSocket connection according to the transport configuration.
func (t *Transport) dial(cancel <-chan struct{}) (*websocket.Conn, error) {
	// Prepare a WebSocket config.
	urlString := t.url.String()
	config, err := websocket.NewConfig(urlString, "http://localhost")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create WebSocket config")
	}

	// Establish the underlying TCP connection.
	// We need to do this manually so that we can set up the timeout and the cancel channel.
	var conn net.Conn
	dialer := &net.Dialer{
		Timeout: t.dialTimeout,
		Cancel:  cancel,
	}
	switch t.url.Scheme {
	case "ws":
		conn, err = dialer.Dial("tcp", toHostPort(t.url))

	case "wss":
		conn, err = tls.DialWithDialer(dialer, "tcp", toHostPort(t.url), nil)

	default:
		err = errors.Wrapf(websocket.ErrBadScheme, "invalid WebSocket URL scheme: %v", t.url.Scheme)
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to establish TCP connection")
	}

	// Establish the WebSocket connection.
	ws, err := websocket.NewClient(config, conn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to establish WebSocket connection")
	}
	return ws, nil
}

func (t *Transport) updateDeadline(ws *websocket.Conn) error {
	// Set deadline in case read timeout is the same as write timeout.
	if t.readTimeout != 0 && t.writeTimeout == t.readTimeout {
		if err := ws.SetDeadline(time.Now().Add(t.readTimeout)); err != nil {
			return errors.Wrap(err, "failed to set connection deadline")
		}
		return nil
	}

	// Set read deadline.
	if t.readTimeout != 0 {
		if err := ws.SetReadDeadline(time.Now().Add(t.readTimeout)); err != nil {
			return errors.Wrap(err, "failed to set connection read deadline")
		}
	}

	// Set write deadline.
	if t.writeTimeout != 0 {
		if err := ws.SetWriteDeadline(time.Now().Add(t.writeTimeout)); err != nil {
			return errors.Wrap(err, "failed to set connection write deadline")
		}
	}
	return nil
}

var portMap = map[string]string{
	"ws":  "80",
	"wss": "443",
}

func toHostPort(u *url.URL) string {
	if _, ok := portMap[u.Scheme]; ok {
		if _, _, err := net.SplitHostPort(u.Host); err != nil {
			return net.JoinHostPort(u.Host, portMap[u.Scheme])
		}
	}
	return u.Host
}
