package types

import (
	"encoding/json"
	"strconv"

	"github.com/asuleymanov/steem-go/encoding/transaction"
	"github.com/pkg/errors"
)

func unmarshalInt(data []byte) (int64, error) {
	if len(data) == 0 {
		return 0, errors.New("types: empty data received when unmarshalling an integer")
	}

	var (
		i   int64
		err error
	)
	if data[0] == '"' {
		d := data[1:]
		d = d[:len(d)-1]
		i, err = strconv.ParseInt(string(d), 10, 64)
	} else {
		err = json.Unmarshal(data, &i)
	}
	return i, errors.Wrapf(err, "types: failed to unmarshal integer: %v", data)
}

//Int8 type from parameter JSON
type Int8 int8

//UnmarshalJSON unpacking the JSON parameter in the Int8 type.
func (num *Int8) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int8(v)
	return nil
}

//MarshalTransaction is a function of converting type Int8 to bytes.
func (num Int8) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(int(num))
}

//Int16 type from parameter JSON
type Int16 int16

//UnmarshalJSON unpacking the JSON parameter in the Int16 type.
func (num *Int16) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int16(v)
	return nil
}

//MarshalTransaction is a function of converting type Int16 to bytes.
func (num Int16) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(int16(num))
}

//Int32 type from parameter JSON
type Int32 int32

//UnmarshalJSON unpacking the JSON parameter in the Int32 type.
func (num *Int32) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int32(v)
	return nil
}

//MarshalTransaction is a function of converting type Int32 to bytes.
func (num Int32) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(int32(num))
}

//Int64 type from parameter JSON
type Int64 int64

//UnmarshalJSON unpacking the JSON parameter in the Int64 type.
func (num *Int64) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int64(v)
	return nil
}

//MarshalTransaction is a function of converting type Int64 to bytes.
func (num Int64) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(int64(num))
}
