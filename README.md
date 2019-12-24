# Go 简易聊天室

Go 语言练手 Demo。

## 设计

### 客户端

- 与服务端建立连接 `net.Dial("tcp", address)`

### 服务端

功能：

- 监听某端口连接 `listener, err := net.Listen("tcp", address)`
- 循环等待接收请求 `listener.Accept()`

## 涉及知识点

- 接口
- 并发
- 网络编程
  - [Go socket编程实践: TCP服务器和客户端实现](https://colobu.com/2014/12/02/go-socket-programming-TCP/)
  - [tcp 服务器](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.1.md)
- 终端 UI