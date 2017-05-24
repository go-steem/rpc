package call

import (
	// Stdlib
	"encoding/json"

	// Vendor
	"github.com/asuleymanov/golos-go/interfaces"
)

var EmptyParams = []string{}

func Raw(caller interfaces.Caller, method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := caller.Call(method, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
