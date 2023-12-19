package api

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nihankhan/chatz/internal/models"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool) //stores all the connected clients uniquely
var broadcast = make(chan models.Message)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var m models.Message
		err := conn.ReadJSON(&m)
		if err != nil {
			log.Println(err)
			delete(clients, conn)
			break
		}
		broadcast <- m
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

/*

project-root
|-- handler
|   |-- chat.go
|-- models
|   |-- model.go
|-- websocket
|   |-- websocket.go
|-- routers
|   |-- router.go
|-- main.go
|--web
    |-- static
	|   |-- chat.js
	|	|-- style.css
	|-- templates
	    |-- index.html

*/
