package tests

import (
	"CodeChat/internal/client"
	"CodeChat/internal/server"
	"testing"
)

func TestStartServer(t *testing.T) {
	go server.StartServer()
}

func TestStartClient(t *testing.T) {
	go client.StartClient()
}

func TestChat(t *testing.T) {
	go server.StartServer()
	go client.StartClient()
}
