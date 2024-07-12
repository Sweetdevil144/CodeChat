package client

import (
	"bufio"
	"os"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

func StartClient() error {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	blue := color.New(color.FgBlue).PrintlnFunc()
	red := color.New(color.FgRed).PrintlnFunc()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				red("Error reading message:", err)
				return
			}
			blue("Recieved : ", string(message))
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
