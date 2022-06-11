package websocket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/atomic"
)

type WS struct {
	Upgrader websocket.Upgrader
	Counter  atomic.Int32
	Clients  sync.Map
}

func NewWS() *WS {
	return &WS{
		Upgrader: websocket.Upgrader{},
		Counter:  atomic.Int32{},
		Clients:  sync.Map{},
	}
}

func (w *WS) stopClients() {
	w.Clients.Range(func(key, value interface{}) bool {
		log.Printf("stop client key: %s", key)
		value.(*Client).Close()
		return true
	})
}
