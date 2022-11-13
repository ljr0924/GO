package client

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn

	// 当前客户端模式
	flag int
}

func NewClient(serverIp string, serverPort int) *Client {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("服务器连接错误..", err)
		return nil
	}

	fmt.Println("服务器连接成功")

	return &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		conn:       conn,
		flag:       0,
	}
}

func (c *Client) CloseConn() {
	err := c.conn.Close()
	if err != nil {
		fmt.Println("关闭连接失败，error：", err)
		return
	}
}

func (c *Client) ReceiveMessage() {
	io.Copy(os.Stdout, c.conn)
}

func (c *Client) Menu() bool {
	fmt.Println("1: 广播模式")
	fmt.Println("2: 私聊模式")
	fmt.Println("3: 更改用户名")
	fmt.Println("0: 退出")

	var flag int

	fmt.Scanln(&flag)

	if 0 <= flag && flag <= 3 {
		c.flag = flag
		return true
	} else {
		return false
	}
}

func (c *Client) Run() {
	for {
		for c.Menu() {
			switch c.flag {
			case 0:
				fmt.Println("退出客户端")
				c.CloseConn()
				return
			case 1: // 广播模式
				c.BroadCast()
			case 2: // 私聊模式
				c.PrivateChat()
			case 3: // 更改用户名
				c.UpdateName()
			}
		}
	}
}

func (c *Client) BroadCast() {
	fmt.Print("请输入广播内容：")
	var msg string
	_, err := fmt.Scanln(&msg)
	if err != nil {
		fmt.Println("error：", err)
		return
	}
	c.conn.Write([]byte(msg + "\n\n"))
}

func (c *Client) UpdateName() {
	fmt.Print("请输入新名字：")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("error：", err)
		return
	}

	c.conn.Write([]byte("set name " + name + "\n"))
}

func (c *Client) PrivateChat() {
	// 查询在线用户
	c.conn.Write([]byte("who\n"))
	fmt.Print("请输入私聊对象名字：")
	var name string
	fmt.Scanln(&name)

	fmt.Print("请输入要发送的消息：")
	var msg string
	fmt.Scanln(&msg)

	c.conn.Write([]byte("to " + name + " " + msg + "\n"))

}
