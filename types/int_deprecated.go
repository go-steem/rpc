package types

import (
	"encoding/json"
	"math/big"
)

//Int type from parameter JSON
type Int struct {
	*big.Int
}

//UnmarshalJSON unpacking the JSON parameter in the Int type.
func (num *Int) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		data = data[1:]
		data = data[:len(data)-1]
		var value big.Int
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		num.Int = &value
		return nil
	}

	var value int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	num.Int = big.NewInt(value)
	return nil
}
