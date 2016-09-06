package transaction

import (
	// Stdlib
	"encoding/binary"
	"io"
	"strings"

	// Vendor
	"github.com/pkg/errors"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}

func (encoder *Encoder) EncodeVarint(i int64) error {
	if i >= 0 {
		return encoder.EncodeUVarint(uint64(i))
	}

	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(b, i)
	return encoder.writeBytes(b[:n])
}

func (encoder *Encoder) EncodeUVarint(i uint64) error {
	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(b, i)
	return encoder.writeBytes(b[:n])
}

func (encoder *Encoder) EncodeNumber(v interface{}) error {
	if err := binary.Write(encoder.w, binary.LittleEndian, v); err != nil {
		return errors.Wrapf(err, "encoder: failed to write number: %v", v)
	}
	return nil
}

func (encoder *Encoder) Encode(v interface{}) error {
	if marshaller, ok := v.(TransactionMarshaller); ok {
		return marshaller.MarshalTransaction(encoder)
	}

	switch v := v.(type) {
	case int:
		return encoder.EncodeNumber(v)
	case int8:
		return encoder.EncodeNumber(v)
	case int16:
		return encoder.EncodeNumber(v)
	case int32:
		return encoder.EncodeNumber(v)
	case int64:
		return encoder.EncodeNumber(v)

	case uint:
		return encoder.EncodeNumber(v)
	case uint8:
		return encoder.EncodeNumber(v)
	case uint16:
		return encoder.EncodeNumber(v)
	case uint32:
		return encoder.EncodeNumber(v)
	case uint64:
		return encoder.EncodeNumber(v)

	case string:
		return encoder.encodeString(v)

	default:
		return errors.Errorf("encoder: unsupported type encountered")
	}
}

func (encoder *Encoder) encodeString(v string) error {
	if err := encoder.EncodeUVarint(uint64(len(v))); err != nil {
		return errors.Wrapf(err, "encoder: failed to write string: %v", v)
	}

	return encoder.writeString(v)
}

func (encoder *Encoder) writeBytes(bs []byte) error {
	if _, err := encoder.w.Write(bs); err != nil {
		return errors.Wrapf(err, "encoder: failed to write bytes: %v", bs)
	}
	return nil
}

func (encoder *Encoder) writeString(s string) error {
	if _, err := io.Copy(encoder.w, strings.NewReader(s)); err != nil {
		return errors.Wrapf(err, "encoder: failed to write string: %v", s)
	}
	return nil
}
