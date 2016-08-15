package database

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

const (
	CustomJSONOperationIDFollow = "follow"
)

var customJSONOpBodyObjects = map[string]interface{}{
	CustomJSONOperationIDFollow: &FollowOperation{},
}

// FC_REFLECT( steemit::chain::custom_json_operation,
//             (required_auths)
//             (required_posting_auths)
//             (id)
//             (json) )

type CustomJSONOperation struct {
	RequiredAuths        []string `json:"required_auths"`
	RequiredPostingAuths []string `json:"required_posting_auths"`
	ID                   string   `json:"id"`
	JSON                 string   `json:"json"`
}

func (op *CustomJSONOperation) UnmarshalBody() (interface{}, error) {
	// Get the corresponding operation object template.
	bodyTemplate, ok := customJSONOpBodyObjects[op.ID]
	if !ok {
		// In case there is no corresponding template, return unquoted data.
		return op.JSON, nil
	}

	// Clone the template.
	body := reflect.New(reflect.Indirect(reflect.ValueOf(bodyTemplate)).Type()).Interface()

	// Unmarshal into the new object instance.
	if err := json.NewDecoder(strings.NewReader(op.JSON)).Decode(body); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal CustomJSONOperation.JSON")
	}

	return body, nil
}

type FollowOperation struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}
