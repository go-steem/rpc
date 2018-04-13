package types

import (
	"encoding/json"
	"errors"
	"reflect"
)

type FollowOperation struct {
	Follower  string   `json:"follower"`
	Following string   `json:"following"`
	What      []string `json:"what"`
}

type ReblogOperation struct {
	Account  string `json:"account"`
	Author   string `json:"author"`
	Permlink string `json:"permlink"`
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
