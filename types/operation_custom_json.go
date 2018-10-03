package types

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"

	"github.com/asuleymanov/steem-go/encoding/transaction"
	"github.com/pkg/errors"
)

var (
	TypeFollow         = "follow"
	TypeReblog         = "reblog"
	TypeLogin          = "login"
	TypePrivateMessage = "private_message"
)

var customJSONDataObjects = map[string]interface{}{
	TypeFollow:         &FollowOperation{},
	TypeReblog:         &ReblogOperation{},
	TypeLogin:          &LoginOperation{},
	TypePrivateMessage: &PrivateMessageOperation{},
}

// FC_REFLECT( steemit::chain::custom_json_operation,
//             (required_auths)
//             (required_posting_auths)
//             (id)
//             (json) )

//CustomJSONOperation represents custom_json operation data.
type CustomJSONOperation struct {
	RequiredAuths        []string `json:"required_auths"`
	RequiredPostingAuths []string `json:"required_posting_auths"`
	ID                   string   `json:"id"`
	JSON                 string   `json:"json"`
}

//Type function that defines the type of operation.
func (op *CustomJSONOperation) Type() OpType {
	return TypeCustomJSON
}

//Data returns the operation data.
func (op *CustomJSONOperation) Data() interface{} {
	return op
}

//UnmarshalData unpacking the JSON parameter in the CustomJSONOperation type.
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
			return nil, errors.Wrapf(err, "failed to unmarshal CustomJSONOperation.JSON: \n%v", op.JSON)
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
		return nil, errors.Wrapf(err, "failed to unmarshal CustomJSONOperation.JSON: \n%v", op.JSON)
	}

	return opData, nil
}

//MarshalTransaction is a function of converting type CustomJSONOperation to bytes.
func (op *CustomJSONOperation) MarshalTransaction(encoder *transaction.Encoder) error {
	enc := transaction.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(TypeCustomJSON.Code()))
	enc.EncodeArrString(op.RequiredAuths)
	enc.EncodeArrString(op.RequiredPostingAuths)
	enc.Encode(op.ID)
	enc.Encode(op.JSON)
	return enc.Err()
}

//FollowOperation the structure for the operation CustomJSONOperation.
type FollowOperation struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}

//ReblogOperation the structure for the operation CustomJSONOperation.
type ReblogOperation struct {
	Account  string `json:"account"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
}

//LoginOperation the structure for the operation CustomJSONOperation.
type LoginOperation struct {
	Account string `json:"account"`
}

//PrivateMessageOperation the structure for the operation CustomJSONOperation.
type PrivateMessageOperation struct {
	From             string `json:"from"`
	To               string `json:"to"`
	FromMemoKey      string `json:"from_memo_key"`
	ToMemoKey        string `json:"to_memo_key"`
	SentTime         uint64 `json:"sent_time"`
	Checksum         uint32 `json:"checksum"`
	EncryptedMessage string `json:"encrypted_message"`
}

//MarshalCustomJSON generate a row from the structure fields.
func MarshalCustomJSON(v interface{}) (string, error) {
	var tmp []interface{}

	typeInterface := reflect.TypeOf(v).Name()
	switch typeInterface {
	case "FollowOperation":
		tmp = append(tmp, TypeFollow)
	case "ReblogOperation":
		tmp = append(tmp, TypeReblog)
	case "LoginOperation":
		tmp = append(tmp, TypeLogin)
	case "PrivateMessageOperation":
		tmp = append(tmp, TypePrivateMessage)
	default:
		return "", errors.New("Unknown type")
	}

	tmp = append(tmp, v)

	b, err := json.Marshal(tmp)
	if err != nil {
		return "", err
	}

	return string(b), nil //strings.Replace(string(b), "\"", "\\\"", -1), nil
}
