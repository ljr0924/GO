package server

import (
	"fmt"
	"io"
	"net"
	"sync"
)

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
	s.CMessage <- NewMsg(u, fmt.Sprintf("[%s]%s\n", u.Name, msg))
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
		data := make([]byte, 1024)
		n, err := u.conn.Read(data)
		if n == 0 {
			s.Offline(u)
			return
		}
		if err != nil && err != io.EOF {
			fmt.Printf("[%s]读取数据错误，error：%s", u.Addr, err)
			return
		}
		data = data[:n-1]
		msg := string(data)
		msgLen := len(msg)
		if msgLen > 8 && msg[:8] == "set name" {
			u.SetName(msg[9:])
		} else if msg == "who" {
			s.onlineMapLock.RLock()
			for _, onlineUser := range s.OnlineMap {
				if onlineUser.Addr == u.Addr {
					continue
				}
				row := fmt.Sprintf("[%s]%s，在线...\n", onlineUser.Addr, onlineUser.Name)
				_, err := u.conn.Write([]byte(row))
				if err != nil {
					fmt.Println("发送在线用户名单失败，error：", err)
					continue
				}
			}
			s.onlineMapLock.RUnlock()

		} else {
			s.Broadcast(u, msg)
		}
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
