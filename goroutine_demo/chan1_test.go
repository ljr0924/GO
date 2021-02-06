package goroutine_demo

import (
    "fmt"
    "strconv"
    "testing"
    "time"
)

func TestChan1(t *testing.T) {

    ch := make(chan string, 5)
    done := make(chan struct{})

    // 生产者
    go func() {
        for i := 1; i <= 50; i++ {
            if i % 5 == 0 {
                time.Sleep(time.Second * 2)
            }
            ch <- "string" + strconv.Itoa(i)
        }
        close(ch)
    }()

    // 消费者
    go func() {
        for s := range ch {
            fmt.Println(s)
        }
        // 通知已经没有消息
        done <- struct{}{}
    }()

    <- done
}
