package server

import (
	"fmt"
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan *Msg
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()

	u := &User{
		Name: addr,
		Addr: addr,
		C:    make(chan *Msg),
		conn: conn,
	}

	go u.Sender()

	return u
}

func (u *User) Sender() {
	for {
		msg := <-u.C
		_, err := u.conn.Write([]byte(msg.Msg))
		if err != nil {
			fmt.Printf("发送消息给%s失败，error：%s", u.Name, err.Error())
			return
		}
	}
}

func (u *User) SetName(name string) {
	u.Name = name
}
