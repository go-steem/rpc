package types

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

//StringSlice type from parameter JSON
type StringSlice []string

//UnmarshalJSON unpacking the JSON parameter in the StringSlice type.
func (ss *StringSlice) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if data[0] == '[' {
		var v []string
		if err := json.Unmarshal(data, &v); err != nil {
			return errors.Wrap(err, "failed to unmarshal string slice")
		}
		*ss = v
	} else {
		var v string
		if err := json.Unmarshal(data, &v); err != nil {
			return errors.Wrap(err, "failed to unmarshal string slice")
		}
		*ss = strings.Split(v, " ")
	}
	return nil
}
