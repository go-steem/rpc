package transaction

type RollingEncoder struct {
	next *Encoder
	err  error
}

func NewRollingEncoder(next *Encoder) *RollingEncoder {
	return &RollingEncoder{next, nil}
}

func (encoder *RollingEncoder) EncodeVarint(i int64) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeVarint(i)
	}
}

func (encoder *RollingEncoder) EncodeUVarint(i uint64) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeUVarint(i)
	}
}

func (encoder *RollingEncoder) EncodeNumber(v interface{}) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeNumber(v)
	}
}

func (encoder *RollingEncoder) EncodeBool(v bool) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeBool(v)
	}
}

func (encoder *RollingEncoder) EncodeMoney(v string) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeMoney(v)
	}
}

func (encoder *RollingEncoder) EncodeString(v string) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeString(v)
	}
}

func (encoder *RollingEncoder) EncodePubKey(v string) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodePubKey(v)
	}
}

func (encoder *RollingEncoder) EncodeArrString(v []string) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeArrString(v)
	}
}

func (encoder *RollingEncoder) Encode(v interface{}) {
	if encoder.err == nil {
		encoder.err = encoder.next.Encode(v)
	}
}

func (encoder *RollingEncoder) Err() error {
	return encoder.err
}
