package websocket

import (
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn      *websocket.Conn
	writeChan chan message
	name      string
	once      sync.Once
	ws        *WS
}

type message struct {
	data string
	code int
}

func NewClient(conn *websocket.Conn, ws *WS) *Client {
	client := &Client{
		conn:      conn,
		writeChan: make(chan message, 1),
		name:      uuid.New().String(),
		once:      sync.Once{},
		ws:        ws,
	}

	conn.SetPingHandler(func(appData string) error {
		client.writeMessage("pong", websocket.PongMessage)
		return nil
	})

	ws.Counter.Inc()

	log.Printf("client_name: %s, counter: %d", client.name, ws.Counter.Load())
	return client
}

func (c *Client) close() error {
	c.ws.Counter.Dec()

	log.Printf("client_name: %s, counter: %d", c.name, c.ws.Counter.Load())
	c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	c.ws.Clients.Delete(c.name)
	close(c.writeChan)
	return c.conn.Close()
}

func (c *Client) Close() (err error) {
	c.once.Do(func() {
		err = c.close()
	})

	return
}

func (c *Client) Read() {
	defer c.Close()
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("client_name: %s, err: %s", c.name, err.Error())
			return
		}

		log.Printf("client_name: %s, mt: %d, message: %s", c.name, mt, message)
		c.WriteMessage("pong!")
	}
}

func (c *Client) writeMessage(data string, code int) {
	c.writeChan <- message{data, code}
}

func (c *Client) WriteMessage(data string) {
	c.writeChan <- message{data, websocket.TextMessage}
}

func (c *Client) Write() {
	for msg := range c.writeChan {
		err := c.conn.WriteMessage(msg.code, []byte(msg.data))
		if err != nil {
			log.Printf("write err: %s", err.Error())
			c.Close()
			return
		}
	}
}

func (c *Client) Name() string {
	return c.name
}
