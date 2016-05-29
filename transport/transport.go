package transports

type Transport interface {
	Call(method string, params, response interface{}) error
	Close() error
}
