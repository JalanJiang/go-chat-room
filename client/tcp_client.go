package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TCPChatClient TCP客户端
type TCPChatClient struct {
	conn net.Conn
}

// Start 启动
func (c *TCPChatClient) Start() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		// clientName, _ := inputReader.ReadString('\n')
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")

		if trimmedInput == "Q" {
			return
		}

		c.conn.Write([]byte("says: " + trimmedInput))
	}
}

// Dial 建立连接
func (c *TCPChatClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)

	if err == nil {
		c.conn = conn
	}

	// TODO: 字节流处理

	return err
}

// Close 关闭
func (c *TCPChatClient) Close() {}
