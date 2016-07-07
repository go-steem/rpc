package interfaces

type Caller interface {
	Call(method string, params, response interface{}) error
}
