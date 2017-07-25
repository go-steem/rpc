package transaction

import (
	// Stdlib
	"encoding/binary"
	"io"
	"regexp"
	"strconv"
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

func (encoder *Encoder) EncodeArrString(v []string) error {
	if err := encoder.EncodeUVarint(uint64(len(v))); err != nil {
		return errors.Wrapf(err, "encoder: failed to write string: %v", v)
	}
	for _, val := range v {
		if err := encoder.EncodeUVarint(uint64(len(val))); err != nil {
			return errors.Wrapf(err, "encoder: failed to write string: %v", val)
		}
		if _, err := io.Copy(encoder.w, strings.NewReader(val)); err != nil {
			return errors.Wrapf(err, "encoder: failed to write string: %v", val)
		}
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

func (encoder *Encoder) EncodeBool(b bool) error {
	if b {
		return encoder.EncodeNumber(byte(1))
	} else {
		return encoder.EncodeNumber(byte(0))
	}
}

func (encoder *Encoder) EncodeMoney(s string) error {
	r, _ := regexp.Compile("^[0-9]+\\.?[0-9]* [A-Za-z0-9]+$")
	if r.MatchString(s) {
		asset := strings.Split(s, " ")
		amm, _ := strconv.ParseInt(strings.Replace(asset[0], ".", "", -1), 10, 64)
		ind := strings.Index(asset[0], ".")
		var perc int
		if ind == -1 {
			perc = 0
		} else {
			perc = len(asset[0]) - ind - 1
		}
		if err := binary.Write(encoder.w, binary.LittleEndian, amm); err != nil {
			return errors.Wrapf(err, "encoder: failed to write number: %v", amm)
		}
		if err := binary.Write(encoder.w, binary.LittleEndian, byte(perc)); err != nil {
			return errors.Wrapf(err, "encoder: failed to write number: %v", perc)
		}

		if _, err := io.Copy(encoder.w, strings.NewReader(asset[1])); err != nil {
			return errors.Wrapf(err, "encoder: failed to write string: %v", asset[1])
		}

		for i := byte(len(asset[1])); i < 7; i++ {
			if err := binary.Write(encoder.w, binary.LittleEndian, byte(0)); err != nil {
				return errors.Wrapf(err, "encoder: failed to write number: %v", 0)
			}
		}
		return nil
	} else {
		return errors.New("Expecting amount like '99.000 SYMBOL'")
	}
}
