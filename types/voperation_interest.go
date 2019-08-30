package types

//InterestOperation represents interest operation data.
type InterestOperation struct {
	Owner    string `json:"owner"`
	Interest *Asset `json:"interest"`
}

//Type function that defines the type of operation InterestOperation.
func (op *InterestOperation) Type() OpType {
	return TypeInterest
}

//Data returns the operation data InterestOperation.
func (op *InterestOperation) Data() interface{} {
	return op
}
