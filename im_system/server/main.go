package main

import "go_demo/im_system/server/server"

func main() {
	s := server.NewServer("127.0.0.1", 8888)
	s.Start()
}
