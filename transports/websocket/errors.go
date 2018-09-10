package websocket

import "errors"

//ErrClosing error returned about a closed channel
var ErrClosing = errors.New("closing")
