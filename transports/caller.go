package transports

import (
	"fmt"
	"io"
)

//Caller interface for sending a request to a network transport
type Caller interface {
	Call(method string, params, response interface{}) error
}

//CallCloser network transport interface
type CallCloser interface {
	Caller
	io.Closer
}

type RPCError struct {
	Code    int64      `json:"code"`
	Message string     `json:"message"`
	Datas   RPCErrData `json:"data"`
}

type RPCErrData struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Stack   []struct {
		Context struct {
			Level      string `json:"level"`
			File       string `json:"file"`
			Line       int    `json:"line"`
			Method     string `json:"method"`
			Hostname   string `json:"hostname"`
			ThreadName string `json:"thread_name"`
			Timestamp  string `json:"timestamp"`
		} `json:"context"`
		Format string      `json:"format"`
		Data   interface{} `json:"data"`
	} `json:"stack"`
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
