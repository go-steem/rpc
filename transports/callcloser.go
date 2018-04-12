package transports

import "io"

type CallCloser interface {
	Caller
	io.Closer
}
