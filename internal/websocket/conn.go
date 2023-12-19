// websocket/websocket.go
package websocket

import (
	"net/http"

	"github.com/nihankhan/chatz/internal/models"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
