package types

//FillTransferFromSavingsOperation represents fill_transfer_from_savings operation data.
type FillTransferFromSavingsOperation struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    *Asset `json:"amount"`
	RequestID uint32 `json:"request_id"`
	Memo      string `json:"memo"`
}

//Type function that defines the type of operation FillTransferFromSavingsOperation.
func (op *FillTransferFromSavingsOperation) Type() OpType {
	return TypeFillTransferFromSavings
}

//Data returns the operation data FillTransferFromSavingsOperation.
func (op *FillTransferFromSavingsOperation) Data() interface{} {
	return op
}
