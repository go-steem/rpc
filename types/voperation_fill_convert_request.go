package types

//FillConvertRequestOperation represents fill_convert_request operation data.
type FillConvertRequestOperation struct {
	Owner     string `json:"owner"`
	Requestid uint32 `json:"requestid"`
	AmountIn  *Asset `json:"amount_in"`
	AmountOut *Asset `json:"amount_out"`
}

//Type function that defines the type of operation FillConvertRequestOperation.
func (op *FillConvertRequestOperation) Type() OpType {
	return TypeFillConvertRequest
}

//Data returns the operation data FillConvertRequestOperation.
func (op *FillConvertRequestOperation) Data() interface{} {
	return op
}
