package types

import (
	"encoding/json"
	"fmt"
)

type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params,omitempty"`
	JSON   string      `json:"jsonrpc"`
	ID     uint64      `json:"id"`
}

type RPCResponse struct {
	Result *json.RawMessage `json:"result,omitempty"`
	Error  *RPCError        `json:"error,omitempty"`
	JSON   string           `json:"jsonrpc,omitempty"`
	ID     uint64           `json:"id"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
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
	} `json:"data"`
}

type RPCIncoming struct {
	ID     uint64          `json:"id"`
	JSON   string          `json:"jsonrpc"`
	Result json.RawMessage `json:"result"`
}

/*
Old Version
type RPCIncoming struct {
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}
*/

func (e *RPCError) Error() string {
	return fmt.Sprintf("%d: %s\n %#v", e.Code, e.Message, e.Data)
}
