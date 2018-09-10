package transaction

//RollingEncoder structure for the chain of converters
type RollingEncoder struct {
	next *Encoder
	err  error
}

//NewRollingEncoder initializing the chain of converters
func NewRollingEncoder(next *Encoder) *RollingEncoder {
	return &RollingEncoder{next, nil}
}

//EncodeVarint adding int64 to the converted value
func (encoder *RollingEncoder) EncodeVarint(i int64) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeVarint(i)
	}
}

//EncodeUVarint adding uint64 to the converted value
func (encoder *RollingEncoder) EncodeUVarint(i uint64) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeUVarint(i)
	}
}

//EncodeNumber adding number to the converted value
func (encoder *RollingEncoder) EncodeNumber(v interface{}) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeNumber(v)
	}
}

//EncodeBool adding bool to the converted value
func (encoder *RollingEncoder) EncodeBool(v bool) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeBool(v)
	}
}

//EncodeMoney adding Asset to the converted value
func (encoder *RollingEncoder) EncodeMoney(v string) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeMoney(v)
	}
}

//EncodeString adding string to the converted value
func (encoder *RollingEncoder) EncodeString(v string) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeString(v)
	}
}

//EncodePubKey adding PubKey to the converted value
func (encoder *RollingEncoder) EncodePubKey(v string) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodePubKey(v)
	}
}

//EncodeArrString adding []string to the converted value
func (encoder *RollingEncoder) EncodeArrString(v []string) {
	if encoder.err == nil {
		encoder.err = encoder.next.EncodeArrString(v)
	}
}

//Encode adding to a chain of other values
func (encoder *RollingEncoder) Encode(v interface{}) {
	if encoder.err == nil {
		encoder.err = encoder.next.Encode(v)
	}
}

//Err function that returns an error (if any) from the cup of converters
func (encoder *RollingEncoder) Err() error {
	return encoder.err
}
