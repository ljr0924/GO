package goroutine_demo

import (
    "testing"
    "time"
)

func TestGoroutine1(t *testing.T) {

    // 使用chan阻塞主goroutine，当子goroutine执行完毕，主goroutine继续执行
    exit := make(chan struct{})

    go func() {
        time.Sleep(time.Second * 3)
        t.Log("goroutine done!!!!!!")
        close(exit)
        // exit <- struct{}{}
    }()

    t.Log("main goroutine")
    <- exit
    t.Log("main goroutine done!!!!!")


}
