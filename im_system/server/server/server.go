package server

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"
)

var CmdPrefix = []byte("call")

type Server struct {
	Ip   string
	Port int
	// 在线用户 addr -> User
	OnlineMap     map[string]*User
	onlineMapLock sync.RWMutex

	// 消息广播channel
	CMessage chan *Msg
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		CMessage:  make(chan *Msg),
	}
}

func (s *Server) Broadcast(u *User, msg string) {
	s.CMessage <- NewMsg(u, fmt.Sprintf("[%s]%s", u.Name, msg))
}

func (s *Server) MessageListener() {
	for {
		msg := <-s.CMessage
		s.onlineMapLock.RLock()
		for _, user := range s.OnlineMap {
			if user.Addr == msg.User.Addr {
				continue
			}
			user.C <- msg
		}
		s.onlineMapLock.RUnlock()
	}
}

func (s *Server) Online(u *User) {
	s.onlineMapLock.Lock()
	s.OnlineMap[u.Addr] = u
	s.onlineMapLock.Unlock()

	s.Broadcast(u, fmt.Sprintf("%s，上线了\n", u.Name))
}

func (s *Server) Offline(u *User) {
	s.onlineMapLock.Lock()
	delete(s.OnlineMap, u.Addr)
	s.onlineMapLock.Unlock()
	s.Broadcast(u, fmt.Sprintf("%s，下线了\n", u.Name))
}

func (s *Server) MessageHandler(u *User) {
	for {
		msg := make([]byte, 1024)
		n, err := u.conn.Read(msg)
		if n == 0 {
			s.Offline(u)
			return
		}
		if err != nil && err != io.EOF {
			fmt.Printf("[%s]读取数据错误，error：%s", u.Addr, err)
			return
		}
		msg = msg[:n-1]
		if bytes.Equal(msg[:4], CmdPrefix) {
			s.CmdHandler(u, msg[5:])
		} else {
			s.Broadcast(u, string(msg))
		}
	}
}

func (s *Server) CmdHandler(u *User, cmd []byte) {
	cmds := bytes.Split(cmd, []byte(" "))
	cmdName, cmdValue := cmds[0], cmds[1]
	// 设置用户名称
	fmt.Printf("%s %v", cmdName, bytes.Equal(cmdName, []byte("setname")))
	if bytes.Equal(cmdName, []byte("setname")) {
		u.SetName(string(cmdValue))
	}
}

func (s *Server) Handler(conn net.Conn) {
	fmt.Printf("连接成功，对方IP：%s\n", conn.RemoteAddr().String())
	user := NewUser(conn)

	s.Online(user)

	go s.MessageHandler(user)
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("listen server err: ", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("listen server err：", err)
		}
	}(listener)

	go s.MessageListener()

	fmt.Println("开始监听....")

	for {
		// 监听连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err: ", err)
			return
		}
		go s.Handler(conn)
	}

}
