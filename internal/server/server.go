package server

import (
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin:     func(r *http.Request) bool { return true },
}

func StartServer() error {
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            return
        }
        defer conn.Close()
        for {
            _, message, err := conn.ReadMessage()
            if err != nil {
                break
            }
            // Echo the message back
            err = conn.WriteMessage(websocket.TextMessage, message)
            if err != nil {
                break
            }
        }
    })

    return http.ListenAndServe(":8080", nil)
}
