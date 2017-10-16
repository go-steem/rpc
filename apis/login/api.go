package login

import (
	// Stdlib
	"encoding/json"

	// RPC
	"github.com/asuleymanov/golos-go/interfaces"
	"github.com/asuleymanov/golos-go/internal/rpc"

	// Vendor
	"github.com/pkg/errors"
)

const APIID = "login_api"

var EmptyParams = []string{}

type API struct {
	id     int
	caller interfaces.Caller
}

type Version struct {
	BlockchainVersion string `json:"blockchain_version"`
	SteemRevision     string `json:"steem_revision"`
	FcRevision        string `json:"fc_revision"`
}

func NewAPI(caller interfaces.Caller) (*API, error) {
	id, err := rpc.GetNumericAPIID(caller, APIID)
	if err != nil {
		return nil, err
	}
	return &API{id, caller}, nil
}

func (api *API) Raw(method string, params interface{}) (*json.RawMessage, error) {
	var resp json.RawMessage
	if err := api.caller.Call("call", []interface{}{api.id, method, params}, &resp); err != nil {
		return nil, errors.Wrapf(err, "golos-go: %v: failed to call %v\n", APIID, method)
	}
	return &resp, nil
}

//login
/*func (api *API) Login(username, password string) (bool, error) {
	raw, err := api.Raw("login", []interface{}{username, password})
	if err != nil {
		return false, err
	}
	var resp bool
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return false, errors.Wrap(err, "golos-go: login_api: failed to unmarshal login response")
	}
	return resp, nil
}*/

//get_api_by_name
func (api *API) GetAPIByName(apiName string) (int, error) {
	raw, err := api.Raw("get_api_by_name", []interface{}{apiName})
	if err != nil {
		return 0, err
	}
	var resp int
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, errors.Wrap(err, "golos-go: login_api: failed to unmarshal get_api_by_name response")
	}
	return resp, nil
}

//get_version
func (api *API) GetVersion() (*Version, error) {
	raw, err := api.Raw("get_version", EmptyParams)
	if err != nil {
		return nil, err
	}
	var resp *Version
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, errors.Wrap(err, "golos-go: login_api: failed to unmarshal get_version response")
	}
	return resp, nil
}
