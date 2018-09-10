package transports

import "io"

//Caller interface for sending a request to a network transport
type Caller interface {
	Call(method string, params, response interface{}) error
}

//CallCloser network transport interface
type CallCloser interface {
	Caller
	io.Closer
}
