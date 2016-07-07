package websocket

import (
	"fmt"
	"time"
)

// ConnectingEvent is emitted when a new connection is being established.
type ConnectingEvent struct {
	URL string
}

func (e *ConnectingEvent) String() string {
	return fmt.Sprintf("CONNECTING [url=%v]", e.URL)
}

// ConnectedEvent is emitted when the WebSocket connection is established.
type ConnectedEvent struct {
	URL string
}

func (e *ConnectedEvent) String() string {
	return fmt.Sprintf("CONNECTED [url=%v]", e.URL)
}

// DisconnectedEvent is emitted when the WebSocket connection is lost.
type DisconnectedEvent struct {
	URL string
	Err error
}

func (e *DisconnectedEvent) String() string {
	return fmt.Sprintf("DISCONNECTED [url=%v, err=%v]", e.URL, e.Err)
}

// DialTimeoutEvent is emitted when establishing a new connection times out.
type DialTimeoutEvent struct {
	URL     string
	Err     error
	Timeout time.Duration
}

func (e *DialTimeoutEvent) String() string {
	return fmt.Sprintf("DIAL_TIMEOUT [url=%v, err=%v, timeout=%v]", e.URL, e.Err, e.Timeout)
}
