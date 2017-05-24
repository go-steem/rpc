package login

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/golos-go/interfaces"

	// Vendor
	"github.com/pkg/errors"
)

const (
	APIID         = "login_api"
	NumbericAPIID = 1
)

type API struct {
	caller interfaces.Caller
}

func NewAPI(caller interfaces.Caller) *API {
	return &API{caller}
}

func (api *API) call(method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{NumbericAPIID, method, params}, resp)
}

func (api *API) LoginRaw(username, password string) (*json.RawMessage, error) {
	var resp json.RawMessage
	params := []interface{}{username, password}
	if err := api.call("login", params, &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: login_api: failed to call login")
	}
	return &resp, nil
}

func (api *API) Login(username, password string) (bool, error) {
	var resp bool
	params := []interface{}{username, password}
	if err := api.call("login", params, &resp); err != nil {
		return false, errors.Wrap(err, "golos-go: login_api: failed to call login")
	}
	return resp, nil
}

func (api *API) GetAPIByNameRaw(apiName string) (*json.RawMessage, error) {
	var resp json.RawMessage
	params := []interface{}{apiName}
	if err := api.call("get_api_by_name", params, &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: login_api: failed to call get_api_by_name")
	}
	return &resp, nil
}

func (api *API) GetAPIByName(apiName string) (int, error) {
	var resp int
	params := []interface{}{apiName}
	if err := api.call("get_api_by_name", params, &resp); err != nil {
		return 0, errors.Wrap(err, "golos-go: login_api: failed to call get_api_by_name")
	}
	return resp, nil
}
