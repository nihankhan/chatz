// routers/router.go
package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nihankhan/chatz/internal/api"
)

func Routes() http.Handler {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("web")) //set folder as file handler
	r.Handle("/", fs)

	// WebSocket endpoint
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		api.WebSocketHandler(w, r)
	})

	go api.HandleMessages()

	// Static files (CSS, JS, etc.)
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	return r
}
