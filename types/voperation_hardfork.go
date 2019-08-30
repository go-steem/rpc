package types

//HardforkOperation represents hardfork operation data.
type HardforkOperation struct {
	HardforkID uint32 `json:"hardfork_id"`
}

//Type function that defines the type of operation HardforkOperation.
func (op *HardforkOperation) Type() OpType {
	return TypeHardfork
}

//Data returns the operation data HardforkOperation.
func (op *HardforkOperation) Data() interface{} {
	return op
}
