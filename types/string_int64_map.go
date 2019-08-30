package types

import (
	"encoding/json"
	"errors"
)

//StringInt64Map type from parameter JSON
type StringInt64Map map[string]int64

//MarshalJSON function for packing the StringInt64Map type in JSON.
func (m StringInt64Map) MarshalJSON() ([]byte, error) {
	xs := make([]interface{}, 0, len(m))
	for k, v := range m {
		xs = append(xs, []interface{}{k, v})
	}
	return JSONMarshal(xs)
}

//UnmarshalJSON unpacking the JSON parameter in the StringInt64Map type.
func (m *StringInt64Map) UnmarshalJSON(data []byte) error {
	var xs [][]interface{}
	if err := json.Unmarshal(data, &xs); err != nil {
		return err
	}

	var invalid bool
	mp := make(map[string]int64, len(xs))
	for _, kv := range xs {
		if len(kv) != 2 {
			invalid = true
			break
		}

		k, ok := kv[0].(string)
		if !ok {
			invalid = true
			break
		}

		var v int64
		switch t := kv[1].(type) {
		case float64:
			v = int64(t)
		case int64:
			v = t
		default:
			invalid = true
		}

		mp[k] = v
	}
	if invalid {
		return errors.New("invalid map encoding")
	}

	*m = mp
	return nil
}
