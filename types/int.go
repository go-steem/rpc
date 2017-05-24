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

/*
type Int int

func (num *Int) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int(v)
	return nil
}
*/

type Int8 int8

func (num *Int8) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int8(v)
	return nil
}

func (num Int8) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(int(num))
}

type Int16 int16

func (num *Int16) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int16(v)
	return nil
}

func (num Int16) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(int16(num))
}

type Int32 int32

func (num *Int32) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int32(v)
	return nil
}

func (num Int32) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(int32(num))
}

type Int64 int64

func (num *Int64) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int64(v)
	return nil
}

func (num Int64) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.EncodeNumber(int64(num))
}
