package client

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {

	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := Client.Ping().Err()
	if err != nil {
		fmt.Println(err)
	}

}