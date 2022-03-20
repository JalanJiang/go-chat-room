package main

import (
	"github.com/JalanJiang/go-chat-room/server"
)

func main() {
	// var s server.ChatServer
	s := &server.TCPChatServer{}

	// 监听
	s.Listen("127.0.0.1:9999")

	// 开启服务
	s.Start()
}
