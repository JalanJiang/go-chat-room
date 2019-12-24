package server

// ChatServer 定义服务端接口
type ChatServer interface {
	Listen()    // 监听
	Boardcast() // 广播
	Start()     // 开启
	Close()     // 关闭
}
