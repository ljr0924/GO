package client

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var Client redis.Conn

func init() {
	var err error
	Client, err = redis.Dial("tcp", "192.168.1.181:6379")

	if err != nil {
		fmt.Println(err)
	}

}