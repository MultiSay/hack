package server

import (
	"log"
	"sync"
	"text/template"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"go.uber.org/atomic"
)

var (
	upgrader = websocket.Upgrader{}
	counter  atomic.Int32
	clients  = sync.Map{}
)

type Client struct {
	conn      *websocket.Conn
	writeChan chan message
	name      string
	clients   *sync.Map
	counter   *atomic.Int32
	once      sync.Once
}

type message struct {
	data string
	code int
}

func NewClient(conn *websocket.Conn, counter *atomic.Int32, clients *sync.Map) *Client {
	client := &Client{
		conn:      conn,
		writeChan: make(chan message, 1),
		name:      uuid.New().String(),
		clients:   clients,
		counter:   counter,
		once:      sync.Once{},
	}

	conn.SetPingHandler(func(appData string) error {
		client.writeMessage("pong", websocket.PongMessage)
		return nil
	})

	counter.Inc()

	log.Printf("client_name: %s, counter: %d", client.name, client.counter.Load())
	return client
}

func (c *Client) close() error {
	c.counter.Dec()

	log.Printf("client_name: %s, counter: %d", c.name, c.counter.Load())
	c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	c.clients.Delete(c.name)
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

func stopClients() {
	clients.Range(func(key, value interface{}) bool {
		log.Printf("stop client key: %s", key)
		value.(*Client).Close()
		return true
	})
}

// handleWS Обрабатываем WebSocket
// handleWS godoc
// @Summary Обрабатываем WebSocket
// @Tags ws
// @Description Обрабатываем WebSocket
// @Produce	json
// @Success 200
// @Failure 500 {object} model.ResponseError
// @Router /ws/ [get]
func (s *server) handleWS(ctx echo.Context) error {
	c, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		// upgrade error
		log.Printf("write err: %s", err.Error())
		return err
	}

	client := NewClient(c, &counter, &clients)
	clients.Store(client.Name(), client)
	go client.Read()
	go client.Write()

	return nil
}

func (s *server) hello(ctx echo.Context) error {
	homeTemplate.Execute(ctx.Response(), "ws://"+ctx.Request().Host+"/ws")
	return nil
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server,
"Send" to send a message to the server and "Close" to close the connection.
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
