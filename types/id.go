package types

import (
	"bytes"
	"encoding/json"
)

//ID type from parameter JSON
type ID struct {
	ValueInt    *Int
	ValueString string
}

var dot = []byte{'.'}

//UnmarshalJSON unpacking the JSON parameter in the ID type.
func (id *ID) UnmarshalJSON(data []byte) error {
	if bytes.Contains(data, dot) {
		id.ValueString = string(data)
		return nil
	}

	var value Int
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	id.ValueInt = &value
	return nil
}
