package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	"github.com/asuleymanov/steem-go/types"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

var (
	ErrShutdown = errors.New("connection is shut down")
	writeWait   = 10 * time.Second
	pongWait    = 60 * time.Second
	pingPeriod  = (pongWait * 9) / 10
)

type Transport struct {
	conn *websocket.Conn

	reqMutex  sync.Mutex
	requestID uint64
	pending   map[uint64]*callRequest

	callbackMutex sync.Mutex
	callbackID    uint64
	callbacks     map[uint64]func(args json.RawMessage)

	closing  bool // user has called Close
	shutdown bool // server has told us to stop

	mutex sync.Mutex
}

// Represent an async call
type callRequest struct {
	Error error            // after completion, the error status.
	Done  chan bool        // strobes when call is complete.
	Reply *json.RawMessage // reply message
}

func NewTransport(url string) (*Transport, error) {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	client := &Transport{
		conn:      ws,
		pending:   make(map[uint64]*callRequest),
		callbacks: make(map[uint64]func(args json.RawMessage)),
	}

	go ping(ws)
	go client.input()
	return client, nil
}

func (caller *Transport) Call(method string, args []interface{}, reply interface{}) error {
	caller.reqMutex.Lock()
	defer caller.reqMutex.Unlock()

	caller.mutex.Lock()
	if caller.closing || caller.shutdown {
		caller.mutex.Unlock()
		return ErrShutdown
	}

	// increase request id
	if caller.requestID == math.MaxUint64 {
		caller.requestID = 0
	}
	caller.requestID++
	seq := caller.requestID

	c := &callRequest{
		Done: make(chan bool, 1),
	}
	caller.pending[seq] = c
	caller.mutex.Unlock()

	request := types.RPCRequest{
		Method: method,
		JSON:   "2.0",
		ID:     caller.requestID,
		Params: args,
	}

	// send Json Rcp request
	if err := caller.WriteJSON(request); err != nil {
		caller.mutex.Lock()
		delete(caller.pending, seq)
		caller.mutex.Unlock()
		return err
	}

	// wait for the call to complete
	<-c.Done
	if c.Error != nil {
		return c.Error
	}

	if c.Reply != nil {
		if err := json.Unmarshal(*c.Reply, reply); err != nil {
			return err
		}
	}
	return nil
}

func (caller *Transport) SetCallback(api string, method string, notice func(args json.RawMessage)) error {
	var ans map[string]interface{}
	// increase callback id
	caller.callbackMutex.Lock()
	if caller.callbackID == math.MaxUint64 {
		caller.callbackID = 0
	}
	//caller.callbackID++
	caller.callbackID = caller.requestID + 1
	caller.callbacks[caller.callbackID] = notice
	caller.callbackMutex.Unlock()

	return caller.Call("call", []interface{}{api, method, []interface{}{caller.callbackID}}, ans)
}

func (caller *Transport) input() {
	caller.conn.SetPongHandler(func(string) error { _ = caller.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := caller.conn.ReadMessage()
		if err != nil {
			caller.stop(err)
			return
		}

		var response types.RPCResponse
		if err := json.Unmarshal(message, &response); err != nil {
			caller.stop(err)
			return
		} else {
			if call, ok := caller.pending[response.ID]; ok {
				caller.onCallResponse(response, call)
			} else {
				//the message is not a pending call, but probably a callback notice
				var incoming types.RPCIncoming
				if err := json.Unmarshal(message, &incoming); err != nil {
					caller.stop(err)
					return
				}
				if _, ok := caller.callbacks[incoming.ID]; ok {
					if err := caller.onNotice(incoming); err != nil {
						caller.stop(err)
						return
					}
				} else {
					log.Printf("protocol error: unknown message received: %+v\n", incoming)
					log.Printf("Answer: %+v\n", string(message))
				}
			}
		}
	}
}

// Return pending clients and shutdown the client
func (caller *Transport) stop(err error) {
	caller.reqMutex.Lock()
	caller.shutdown = true
	for _, call := range caller.pending {
		call.Error = err
		call.Done <- true
	}
	caller.reqMutex.Unlock()
}

// Call response handler
func (caller *Transport) onCallResponse(response types.RPCResponse, call *callRequest) {
	caller.mutex.Lock()
	delete(caller.pending, response.ID)
	if response.Error != nil {
		call.Error = response.Error
	}
	call.Reply = response.Result
	call.Done <- true
	caller.mutex.Unlock()
}

// Incoming notice handler
func (caller *Transport) onNotice(incoming types.RPCIncoming) error {
	notice := caller.callbacks[incoming.ID]
	if notice == nil {
		return fmt.Errorf("callback %d is not registered", incoming.ID)
	}

	// invoke callback
	notice(incoming.Result)

	return nil
}

// Close calls the underlying web socket Close method. If the connection is already
// shutting down, ErrShutdown is returned.
func (caller *Transport) Close() error {
	caller.mutex.Lock()
	if caller.closing {
		caller.mutex.Unlock()
		return ErrShutdown
	}
	caller.closing = true
	caller.mutex.Unlock()
	return caller.conn.Close()
}

func ping(ws *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		<-ticker.C
		if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
			log.Println("ping:", err)
		}
	}
}

func (caller *Transport) WriteJSON(v interface{}) error {
	w, err := caller.conn.NextWriter(1)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	err1 := enc.Encode(v)
	err2 := w.Close()
	if err1 != nil {
		return err1
	}
	return err2
}
