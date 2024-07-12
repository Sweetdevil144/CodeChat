package client

import (
    "bufio"
    "os"
    "github.com/fatih/color"
    "github.com/gorilla/websocket"
    "CodeChat/pkg/common"
)

func StartClient() error {
    conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
    if err != nil {
        return err
    }
    defer conn.Close()

    blue := color.New(color.FgBlue).PrintlnFunc()
    red := color.New(color.FgRed).PrintlnFunc()

    // Input username
    scanner := bufio.NewScanner(os.Stdin)
    blue("Enter your username:")
    scanner.Scan()
    username := scanner.Text()

    go func() {
        for {
            _, message, err := conn.ReadMessage()
            if err != nil {
                red("Error reading message:", err)
                return
            }
            var msg common.Message
            msg, err = common.DecodeMessage(message)
            if err != nil {
                red("Error decoding message:", err)
                continue
            }
            blue(msg.Sender + " : " + msg.Content)
        }
    }()

    for scanner.Scan() {
        content := scanner.Text()
        msg := common.Message{Sender: username, Content: content}
        encodedMsg, err := common.EncodeMessage(msg)
        if err != nil {
            red("Error encoding message:", err)
            continue
        }
        if err := conn.WriteMessage(websocket.TextMessage, encodedMsg); err != nil {
            red("Error sending message:", err)
            return err
        }
    }
    return nil
}
