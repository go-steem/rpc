package types

import (
	// Stdlib
	"encoding/json"
	"strconv"

	// RPC
	"github.com/asuleymanov/golos-go/encoding/transaction"

	// Vendor
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

type UInt uint

func (num *UInt) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt(v)
	return nil
}

func (num UInt) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(uint(num))
}

type UInt8 uint8

func (num *UInt8) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt8(v)
	return nil
}

func (num UInt8) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(uint8(num))
}

type UInt16 uint16

func (num *UInt16) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt16(v)
	return nil
}

func (num UInt16) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(uint16(num))
}

type UInt32 uint32

func (num *UInt32) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt32(v)
	return nil
}

func (num UInt32) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(uint32(num))
}

type UInt64 uint64

func (num *UInt64) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt64(v)
	return nil
}

func (num UInt64) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(uint64(num))
}
