package rpc

import (
	"github.com/go-steem/rpc/transports/websocket"
)

//
// Types
//

type Transport interface {
	Call(method string, params, response interface{}) error
	Close() error
}

type TransportConstructor func(address string) (Transport, error)

//
// Transport Registry
//

var registeredTransportConstructors = map[string]TransportConstructor{}

func RegisterTransport(scheme string, constructor TransportConstructor) {
	registeredTransportConstructors[scheme] = constructor
}

func AvailableTransports() []string {
	schemes := make([]string, 0, len(registeredTransportConstructors))
	for scheme := range registeredTransportConstructors {
		schemes = append(schemes, scheme)
	}
	return schemes
}

//
// Transport Registry Init
//

func init() {
	RegisterTransport("ws", func(address string) (Transport, error) {
		return websocket.Dial(address)
	})
}
