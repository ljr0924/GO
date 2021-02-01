package redis_demo

import (
	"fmt"
	"go_demo/redis_demo/client"
	"testing"
)

func TestCommand(t *testing.T) {

	reply, _ := client.Client.Do("setnx", "123", "123")
	fmt.Printf("%+v\n", reply)


}
