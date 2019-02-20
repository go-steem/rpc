package types

//CustomBinaryOperation represents custom_binary operation data.
type CustomBinaryOperation struct {
	RequiredOwnerAuths   []string    `json:"required_owner_auths"`
	RequiredActiveAuths  []string    `json:"required_active_auths"`
	RequiredPostingAuths []string    `json:"required_posting_auths"`
	RequiredAuths        []Authority `json:"required_auths"`
	ID                   string      `json:"id"`
	Datas                []byte      `json:"data"`
}

//Type function that defines the type of operation CustomBinaryOperation.
func (op *CustomBinaryOperation) Type() OpType {
	return TypeCustomBinary
}

//Data returns the operation data CustomBinaryOperation.
func (op *CustomBinaryOperation) Data() interface{} {
	return op
}
