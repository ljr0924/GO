package goroutine_demo

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

type data struct {
    sync.Mutex
}

// 将Mutex作为匿名字段时，相关方法必须实现为pointer-receiver
// 否则会因为复制导致锁机制失效
func (d *data) test(s string) {
    d.Lock()
    defer d.Unlock()

    for i := 0; i < 3; i++ {
        fmt.Println(s)
        time.Sleep(time.Second / 2)
    }
}

func (d data) test1(s string) {
    d.Lock()
    defer d.Unlock()

    for i := 0; i < 3; i++ {
        fmt.Println(s)
        time.Sleep(time.Second / 2)
    }
}

func TestMutex1(t *testing.T) {

    var wg sync.WaitGroup
    wg.Add(2)

    var d data
    go func() {
        defer wg.Done()
        d.test("write")
    }()
    go func() {
        defer wg.Done()
        d.test("read")
    }()
    wg.Wait()
}

func TestMutex2(t *testing.T) {

    var wg sync.WaitGroup
    wg.Add(2)

    var d data
    go func() {
        defer wg.Done()
        d.test1("write")
    }()
    go func() {
        defer wg.Done()
        d.test1("read")
    }()
    wg.Wait()
}
