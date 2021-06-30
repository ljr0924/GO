package redis_demo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var testCh = "test_ch"

func TestPub(t *testing.T) {

	for i := 0; i < 10; i++ {
		c.Publish(testCh, fmt.Sprintf("这是第%d条消息", i))
		time.Sleep(time.Second)
	}

}

func TestSub(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(2)
	for i := 0; i < 2; i++ {

		go func(index int) {

			for {
				receive, err := c.Subscribe(testCh).ReceiveMessage()
				if err != nil {
					t.Log(err)
					break
				}

				t.Logf("这是协程%d收到的信息：%v", index, receive)
			}
			wg.Done()
		}(i)

	}

	wg.Wait()

}