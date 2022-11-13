package main

import (
	"flag"

	"go_demo/im_system/client/client"
)

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "服务器IP（默认127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "服务器端口（默认8888）")
}

func main() {
	flag.Parse()
	c := client.NewClient(serverIp, serverPort)
	if c == nil {
		return
	}

	go c.ReceiveMessage()

	c.Run()
}
