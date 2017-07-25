package types

import (
	// Stdlib
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"

	// Vendor
	"github.com/pkg/errors"

	// RPC
	"github.com/asuleymanov/golos-go/encoding/transaction"
)

const (
	TypeFollow = "follow"
	TypeReblog = "reblog"
)

var customJSONDataObjects = map[string]interface{}{
	TypeFollow: &FollowOperation{},
	TypeReblog: &ReblogOperation{},
}

// FC_REFLECT( steemit::chain::custom_json_operation,
//             (required_auths)
//             (required_posting_auths)
//             (id)
//             (json) )

// CustomJSONOperation represents custom_json operation data.
type CustomJSONOperation struct {
	RequiredAuths        []string `json:"required_auths"`
	RequiredPostingAuths []string `json:"required_posting_auths"`
	ID                   string   `json:"id"`
	JSON                 string   `json:"json"`
}

func (op *CustomJSONOperation) Type() OpType {
	return TypeCustomJSON
}

func (op *CustomJSONOperation) Data() interface{} {
	return op
}

func (op *CustomJSONOperation) UnmarshalData() (interface{}, error) {
	// Get the corresponding data object template.
	template, ok := customJSONDataObjects[op.ID]
	if !ok {
		// In case there is no corresponding template, return nil.
		return nil, nil
	}

	// Clone the template.
	opData := reflect.New(reflect.Indirect(reflect.ValueOf(template)).Type()).Interface()

	// Prepare the whole operation tuple.
	var bodyReader io.Reader
	if op.JSON[0] == '[' {
		rawTuple := make([]json.RawMessage, 2)
		if err := json.NewDecoder(strings.NewReader(op.JSON)).Decode(&rawTuple); err != nil {
			return nil, errors.Wrapf(err,
				"failed to unmarshal CustomJSONOperation.JSON: \n%v", op.JSON)
		}
		if rawTuple[1] == nil {
			return nil, errors.Errorf("invalid CustomJSONOperation.JSON: \n%v", op.JSON)
		}
		bodyReader = bytes.NewReader([]byte(rawTuple[1]))
	} else {
		bodyReader = strings.NewReader(op.JSON)
	}

	// Unmarshal into the new object instance.
	if err := json.NewDecoder(bodyReader).Decode(opData); err != nil {
		return nil, errors.Wrapf(err,
			"failed to unmarshal CustomJSONOperation.JSON: \n%v", op.JSON)
	}

	return opData, nil
}

func (op *CustomJSONOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCustomJSON.Code()))
	enc.EncodeArrString(op.RequiredAuths)
	enc.EncodeArrString(op.RequiredPostingAuths)
	enc.Encode(op.ID)
	enc.Encode(op.JSON)
	return enc.Err()
}
