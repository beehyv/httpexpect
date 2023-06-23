package examples

import (
	"net/http"

	"github.com/fasthttp/websocket"

	"github.com/valyala/fasthttp"
)

// WsHttpHandler is a simple http.Handler that implements WebSocket echo server.
func WsHTTPHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

// WsFastHandler is a simple fasthttp.RequestHandler that implements
// WebSocket echo server.
func WsFastHTTPHandler(ctx *fasthttp.RequestCtx) {
	var upgrader websocket.FastHTTPUpgrader
	err := upgrader.Upgrade(ctx, func(c *websocket.Conn) {
		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				break
			}
			err = c.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	})
	if err != nil {
		panic(err)
	}
}
