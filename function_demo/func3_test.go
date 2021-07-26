package function_demo

import "sync"
import "testing"
import "fmt"


func TestDefer(t *testing.T) {
    // 先输出b 再输出a
    defer fmt.Println("a")
    defer fmt.Println("b")
}


var mtx sync.Mutex

func call() {
    mtx.Lock()
    mtx.Unlock()
}

func deferCall() {
    mtx.Lock()
    defer mtx.Unlock()
}

func BenchmarkCall(b *testing.B) {
    call()
}

func BenchmarkDeferCall(b *testing.B) {
    deferCall()
}