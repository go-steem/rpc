package types

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

type StringSlice []string

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
