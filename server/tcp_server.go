package server

import (
	"fmt"
	"log"
	"net"
)

// TCPChatServer TCP服务端结构体
type TCPChatServer struct {
	listener net.Listener
}

// Start 开启服务
func (s *TCPChatServer) Start() {
	// 监听并接受来自 Client 的连接
	for {
		conn, err := s.listener.Accept()

		if err != nil {
			log.Print(err)
		} else {
			go doServerStuff(conn)
		}
	}
}

// Listen 监听
func (s *TCPChatServer) Listen(address string) error {
	// 创建 listener
	listener, err := net.Listen("tcp", address)

	if err == nil {
		s.listener = listener
	}

	log.Printf("Listening on %v", address)

	return err
}

// Close 关闭服务
func (s *TCPChatServer) Close() {
	s.listener.Close()
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error Reading", err.Error())
			return // 终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}
