package types

//CustomOperation represents custom operation data.
type CustomOperation struct {
	RequiredAuths []string `json:"required_auths"`
	ID            uint16   `json:"id"`
	Datas         string   `json:"data"`
}

//Type function that defines the type of operation CustomOperation.
func (op *CustomOperation) Type() OpType {
	return TypeCustom
}

//Data returns the operation data CustomOperation.
func (op *CustomOperation) Data() interface{} {
	return op
}
