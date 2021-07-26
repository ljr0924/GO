package goroutine_demo

import (
    "fmt"
    "math"
    "runtime"
    "sync"
    "testing"
    "time"
)

func TestWaitGroup(t *testing.T) {
    // 使用sync.WaitGroup控制协程调度
    var wg sync.WaitGroup

    wg.Add(5)
    for i := 1; i <= 5; i++ {
        go func(index int) {
            t.Logf("goroutine %d start\n", index)
            time.Sleep(time.Second * 3)
            t.Logf("goroutine %d done\n", index)
            wg.Done()
        }(i)
    }
    wg.Wait()

}

func count() {
    x := 0
    for i := 0; i < math.MaxUint32; i++ {
        x += 1
    }

    fmt.Println(x)
}

// 执行n遍
func test1(n int) {
    for i := 0; i < n; i++ {
        count()
    }
}

func test2(n int) {

    var wg sync.WaitGroup

    wg.Add(n)

    for i := 0; i < n; i++ {
        go func() {
            count()
            wg.Done()
        }()
    }

    wg.Wait()

}

/*
=== RUN   TestTest1
    goroutine2_test.go:64: CPU逻辑处理核心数: 8
4294967295
4294967295
4294967295
4294967295
4294967295
4294967295
4294967295
4294967295
--- PASS: TestTest1 (9.81s)
PASS
*/
func TestTest1(t *testing.T) {
    n := runtime.GOMAXPROCS(0)
    t.Logf("CPU逻辑处理核心数: %d", n)

    test1(n)

}

/*
=== RUN   TestTest2
    goroutine2_test.go:72: CPU逻辑处理核心数: 8
4294967295
4294967295
4294967295
4294967295
4294967295
4294967295
4294967295
4294967295
--- PASS: TestTest2 (2.40s)
PASS
*/
func TestTest2(t *testing.T) {
    n := runtime.GOMAXPROCS(0)
    t.Logf("CPU逻辑处理核心数: %d", n)

    test2(n)
}