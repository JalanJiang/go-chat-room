package server

// ChatServer 定义服务端接口
type ChatServer interface {
	Listen(address string) error // 监听
	Boardcast()                  // 广播
	Start()                      // 开启服务
	Close()                      // 关闭服务
}
