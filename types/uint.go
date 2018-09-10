package types

import (
	"encoding/json"
	"strconv"

	"github.com/asuleymanov/steem-go/encoding/transaction"
	"github.com/pkg/errors"
)

func unmarshalUInt(data []byte) (uint64, error) {
	if len(data) == 0 {
		return 0, errors.New("types: empty data received when unmarshalling an unsigned integer")
	}

	var (
		i   uint64
		err error
	)
	if data[0] == '"' {
		d := data[1:]
		d = d[:len(d)-1]
		i, err = strconv.ParseUint(string(d), 10, 64)
	} else {
		err = json.Unmarshal(data, &i)
	}
	return i, errors.Wrapf(err, "types: failed to unmarshal unsigned integer: %v", data)
}

//UInt type from parameter JSON
type UInt uint

//UnmarshalJSON unpacking the JSON parameter in the UInt type.
func (num *UInt) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt(v)
	return nil
}

//MarshalTransaction is a function of converting type UInt to bytes.
func (num UInt) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(uint(num))
}

//UInt8 type from parameter JSON
type UInt8 uint8

//UnmarshalJSON unpacking the JSON parameter in the UInt8 type.
func (num *UInt8) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt8(v)
	return nil
}

//MarshalTransaction is a function of converting type UInt8 to bytes.
func (num UInt8) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(uint8(num))
}

//UInt16 type from parameter JSON
type UInt16 uint16

//UnmarshalJSON unpacking the JSON parameter in the UInt16 type.
func (num *UInt16) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt16(v)
	return nil
}

//MarshalTransaction is a function of converting type UInt16 to bytes.
func (num UInt16) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(uint16(num))
}

//UInt32 type from parameter JSON
type UInt32 uint32

//UnmarshalJSON unpacking the JSON parameter in the UInt32 type.
func (num *UInt32) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt32(v)
	return nil
}

//MarshalTransaction is a function of converting type UInt32 to bytes.
func (num UInt32) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(uint32(num))
}

//UInt64 type from parameter JSON
type UInt64 uint64

//UnmarshalJSON unpacking the JSON parameter in the UInt64 type.
func (num *UInt64) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt64(v)
	return nil
}

//MarshalTransaction is a function of converting type UInt64 to bytes.
func (num UInt64) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(uint64(num))
}
