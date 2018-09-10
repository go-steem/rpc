package types

//POW2Operation represents pow2 operation data.
type POW2Operation struct {
	Input      *POW2Input `json:"input"`
	PowSummary uint32     `json:"pow_summary"`
}

//Type function that defines the type of operation POW2Operation.
func (op *POW2Operation) Type() OpType {
	return TypePOW2
}

//Data returns the operation data POW2Operation.
func (op *POW2Operation) Data() interface{} {
	return op
}
