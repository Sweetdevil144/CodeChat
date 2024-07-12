package common

import "encoding/json"

// Message represents a chat message with a sender and content
type Message struct {
    Sender  string `json:"sender"`
    Content string `json:"content"`
}

// EncodeMessage encodes a Message into a JSON byte array
func EncodeMessage(m Message) ([]byte, error) {
    return json.Marshal(m)
}

// DecodeMessage decodes a JSON byte array into a Message
func DecodeMessage(data []byte) (Message, error) {
    var m Message
    err := json.Unmarshal(data, &m)
    return m, err
}
