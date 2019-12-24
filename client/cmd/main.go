package main

import (
	"github.com/go-chat-room/client"
)

func main() {
	client := &client.TCPChatClient{}

	client.Dial("127.0.0.1:9999")
	client.Start()
}
