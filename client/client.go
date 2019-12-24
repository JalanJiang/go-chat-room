package client

// ChatClient 客户端接口
type ChatClient interface {
	SetName(clientName string) error // 设置客户端名称
	Start()                          // 开启
	Close()                          // 关闭
}
