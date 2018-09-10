package types

//FillOrderOperation represents fill_order operation data.
type FillOrderOperation struct {
	CurrentOwner   string `json:"current_owner"`
	CurrentOrderid uint32 `json:"current_orderid"`
	CurrentPays    *Asset `json:"current_pays"`
	OpenOwner      string `json:"open_owner"`
	OpenOrderid    uint32 `json:"open_orderid"`
	OpenPays       *Asset `json:"open_pays"`
}

//Type function that defines the type of operation FillOrderOperation.
func (op *FillOrderOperation) Type() OpType {
	return TypeFillOrder
}

//Data returns the operation data FillOrderOperation.
func (op *FillOrderOperation) Data() interface{} {
	return op
}
