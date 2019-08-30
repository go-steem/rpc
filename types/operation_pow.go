package types

//POWOperation represents pow operation data.
type POWOperation struct {
	WorkerAccount string           `json:"worker_account"`
	BlockID       string           `json:"block_id"`
	Nonce         *Int             `json:"nonce"`
	Work          *POW             `json:"work"`
	Props         *ChainProperties `json:"props"`
}

//Type function that defines the type of operation POWOperation.
func (op *POWOperation) Type() OpType {
	return TypePOW
}

//Data returns the operation data POWOperation.
func (op *POWOperation) Data() interface{} {
	return op
}
