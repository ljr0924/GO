package goroutine_demo

import (
    "strconv"
    "sync"
    "testing"
    "time"
)

// GO并发哲学，不要通过共享内存来通信，而应通过通信来共享内存
func TestChanControlGoroutine(t *testing.T) {

    done := make(chan struct{})   // 监听是否完成
    c := make(chan string)        // 与主goroutine通信的channel

    go func() {
        t.Log("waiting data")
        s := <- c             // 接收从goroutine传过来的数据
        t.Logf("fetch data: %s", s)
        done <- struct{}{}    // 告诉主goroutine已经执行完毕
    }()

    time.Sleep(time.Second * 3)
    c <- "I'm channel C"      // 发送数据到chanel，使子goroutine获取数据
    <- done                   // 阻塞主goroutine，监听子goroutine是否完成
    t.Log("main done!!!!!!!!")

}

func TestCacheChan(t *testing.T) {
    // 带缓冲的channel
    c := make(chan int, 3)

    c <- 1
    c <- 2

    t.Log(<- c)
    t.Log(<- c)
}

func TestCacheChan2(t *testing.T) {
    c := make(chan string, 3)
    done := make(chan struct{})
    go func() {
        for true {
            s, ok := <- c
            if ok {
                t.Log(s)
            } else {
                done <- struct{}{}
            }
        }
    }()

    c <- "string1"
    time.Sleep(time.Second)
    c <- "string2"
    time.Sleep(time.Second)
    c <- "string3"
    time.Sleep(time.Second)
    close(c)
    <- done
}

func TestRangeChan(t *testing.T) {

    ch := make(chan string, 10)
    t.Logf("ch len: %d cap: %d", len(ch), cap(ch))

    // 生产者
    go func() {
        for i := 1; i <= 10; i++ {
            ch <- "string" + strconv.Itoa(i)
            time.Sleep(time.Second * 1 / 2)
        }
        close(ch)
    }()

    // 消费者
    // 使用range循环监听更优雅
    // for s := range ch {
    //     t.Log(s)
    // }

    for {
        if s, ok := <- ch; ok {
            t.Log(s)
        } else {
            break
        }

    }

}

func TestBoard(t *testing.T) {

    var wg sync.WaitGroup
    ready := make(chan struct{})

    // 挂起3个goroutine，当ready关闭后，3个goroutine同时启动
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(index int) {
           name := "goroutine" + strconv.Itoa(index)
           t.Logf("%s wait to run", name)
           // 阻塞goroutine运行，等待信号再启动
           <- ready
           t.Logf("%s running", name)
           wg.Done()
        }(i)
    }

    for i := 3; i >= 1; i-- {
        time.Sleep(time.Second)
        t.Log(i)
    }

    close(ready)

    wg.Wait()

}