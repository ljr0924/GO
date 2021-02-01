package redis_demo

import (
	"fmt"
	"go_demo/redis_demo/client"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {

	key := "key"

	var wg sync.WaitGroup
	var mtx sync.Mutex

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(index int) {
			name := fmt.Sprintf("协程%d", index)
			fmt.Printf("启动%s\n", name)
			time.Sleep(time.Second)
			mtx.Lock()
			reply, err := client.Client.Do("setnx", key, name)
			mtx.Unlock()
			if err != nil {
				fmt.Printf("%s %+v\n", name, err)
				wg.Done()
				return
			}
			fmt.Printf("%s %d\n", name, reply)
			if reflect.ValueOf(reply).Int() == 1 {
				fmt.Printf("%s 获取锁成功\n", name)
				mtx.Lock()
				time.Sleep(1)
				_, _ = client.Client.Do("del", key)
				mtx.Unlock()
			} else {
				fmt.Printf("%s 获取锁失败\n", name)
			}

			wg.Done()
		}(i)
	}
	wg.Wait()

}
