package redis_demo

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {

	r, _ := c.SetNX("nx", "123", 1).Result()
	fmt.Printf("%+v\n", r)

	r1, _ := c.Get("nx").Result()
	t.Log(r1)


}
