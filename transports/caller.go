package transports

import (
	"encoding/json"
	"io"
)

var (
	EmptyParams = []struct{}{}
)

//Caller interface for sending a request to a network transport
type Caller interface {
	Call(method string, args []interface{}, reply interface{}) error
	SetCallback(api string, method string, callback func(raw json.RawMessage)) error
}

//CallCloser network transport interface
type CallCloser interface {
	Caller
	io.Closer
}
