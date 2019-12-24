package client

// ChatClient 客户端接口
type ChatClient interface {
	Dial(address string) error       // 与服务端建立连接
	SetName(clientName string) error // 设置客户端名称
	Start()                          // 开启
	Close()                          // 关闭
}