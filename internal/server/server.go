package server

import (
    "sync"
    "net/http"
    "github.com/fatih/color"
    "github.com/gorilla/websocket"
    "CodeChat/pkg/common"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:      1024,
    WriteBufferSize:     1024,
    CheckOrigin:         func(r *http.Request) bool { return true },
    EnableCompression:   true,
}

var clients = make(map[*websocket.Conn]bool)
var mutex = &sync.Mutex{}

func StartServer() error {
    red := color.New(color.FgRed).PrintlnFunc()
    purple := color.New(color.FgHiMagenta).PrintlnFunc()
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            red("error upgrading connection:", err)
            return
        }
        purple("Connection established")

        mutex.Lock()
        clients[conn] = true
        mutex.Unlock()

        defer func() {
            mutex.Lock()
            delete(clients, conn)
            mutex.Unlock()
            conn.Close()
        }()

        for {
            _, message, err := conn.ReadMessage()
            if err != nil {
                red("error reading message:", err)
                break
            }

            var msg common.Message
            msg, err = common.DecodeMessage(message)
            if err != nil {
                red("Error decoding message:", err)
                continue
            }
            purple("Received message from " + msg.Sender)

            encodedMsg, err := common.EncodeMessage(msg)
            if err != nil {
                red("Error encoding message:", err)
                continue
            }

            mutex.Lock()
            for client := range clients {
                if client != conn {
                    if err := client.WriteMessage(websocket.TextMessage, encodedMsg); err != nil {
                        red("Error sending message to a client:", err)
                        return
                    }
                }
            }
            mutex.Unlock()
        }
    })

    return http.ListenAndServe(":8080", nil)
}
