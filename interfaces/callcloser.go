package interfaces

import "io"

type CallCloser interface {
	Caller
	io.Closer
}
