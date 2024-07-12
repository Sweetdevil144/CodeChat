package server

import (
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func StartServer() error {
	red := color.New(color.FgRed).PrintlnFunc()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			red("error upgrading connection:  ", err)
			return
		}
		defer conn.Close()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				red("error reading message:  ", err)
				break
			}

			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				red("error writing message:  ", err)
				break
			}
		}
	})

	return http.ListenAndServe(":8080", nil)
}
