package types

import (
	"encoding/json"
)

//UnknownOperation represents Unknown operation data.
type UnknownOperation struct {
	kind OpType
	data *json.RawMessage
}

//Type function that defines the type of operation UnknownOperation.
func (op *UnknownOperation) Type() OpType {
	return op.kind
}

//Data returns the operation data UnknownOperation.
func (op *UnknownOperation) Data() interface{} {
	return op.data
}
