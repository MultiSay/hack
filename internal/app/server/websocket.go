package server

import (
	"hack/internal/app/websocket"
	"log"
	"text/template"

	"github.com/labstack/echo/v4"
)

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
	c, err := s.ws.Upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		// upgrade error
		log.Printf("write err: %s", err.Error())
		return err
	}

	client := websocket.NewClient(c, s.ws)
	s.ws.Clients.Store(client.Name(), client)
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
