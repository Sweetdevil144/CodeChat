package client

import (
    "fmt"
    "github.com/gorilla/websocket"
    "os"
    "bufio"
)

func StartClient() error {
    conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
    if err != nil {
        return err
    }
    defer conn.Close()

    go func() {
        for {
            _, message, err := conn.ReadMessage()
            if err != nil {
                fmt.Println("Error reading message:", err)
                return
            }
            fmt.Println("Received:", string(message))
        }
    }()

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        msg := scanner.Text()
        err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
        if err != nil {
            return err
        }
    }
    return nil
}
