package gnet

import (
	"fmt"
	"gce/giface"
	"net"
)

type Connection struct {
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	//当前连接的关闭状态
	isClosed bool
	//当前链接所绑定的处理业务方法的API
	handleAPI giface.HandleFunc
	//该连接的处理方法router
	//Router  giface.IRouter
	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
	//Router
	Router giface.IRouter
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		//读取我们最大的数据到buf中
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChan <- true
			continue
		}
		//调用当前链接所绑定的handleAPI
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("ConnId", c.ConnID, "handle is error", err)
		//	break
		//}
		//创建IRequest
		req := Request {
			conn:c,
			data:buf,
		}
		go func(request giface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

//启动链接 让当前的链接准备开始工作
func (c *Connection) Start() {
	fmt.Println("connetion start() connID:", c.ConnID)
	go c.StartReader()

}

//停止链接 结束当前链接的工作
func (c *Connection) Stop() {
	fmt.Println("connection stop() connID:", c.ConnID)
	if c.isClosed == true {
		c.isClosed = true
		//关闭sockct链接
		c.Conn.Close()
		//回收资源
		close(c.ExitBuffChan)
	}

}

//获取当前链接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端的TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据 将数据发送给远程客户端
func (c *Connection) Send(data []byte) error {
	return nil
}

//初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api giface.HandleFunc) giface.IConnection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		handleAPI:    callback_api,
		isClosed:     false,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}