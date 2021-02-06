package goroutine_demo

import (
    "fmt"
    "os"
    "os/signal"
    "strconv"
    "sync"
    "syscall"
    "testing"
    "time"
)

func TestNotify1(t *testing.T) {

    done := make(chan struct{})
    var wg sync.WaitGroup

    wg.Add(3)
    for i := 1; i <= 3; i++ {
        name := "goroutine" + strconv.Itoa(i)
        go func(n string, d chan struct{}) {
            defer wg.Done()
            for true {
                select {
                case <-d:
                    fmt.Println(n + " exit!!!!!")
                    return
                default:
                    fmt.Println(n + " running")
                    time.Sleep(time.Second * 2)
                }
            }
        }(name, done)
    }

    sig := make(chan os.Signal, 1)
    signal.Notify(sig, os.Interrupt, syscall.SIGINT)
    switch <- sig {
    case syscall.SIGINT:
        fmt.Println("intercept")
    default:
        fmt.Println("default")
    }

    close(done)

    wg.Wait()
}