package goroutine_demo

import (
	"sync"
	"testing"
	"time"
)

func TestSpeedControl(t *testing.T) {

	type task struct {
		A int
		B int
	}


	var ch1 = make(chan *task, 10)
	go func() {
		t1 := time.Now().Unix()
		for i := 0; i < 30; i++ {
			ch1 <- &task{i, i+1}
		}
		close(ch1)
		t.Logf("生产者 %ds", time.Now().Unix() - t1)
	}()

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 1; i <= 2; i++ {
		go func(index int) {
			t.Logf("协程%d开始执行", index)
			for tmp := range ch1 {
				t.Logf("协程%d执行结果%d", index, tmp.A+tmp.B)
				time.Sleep(time.Second)
			}
			t.Logf("协程%d结束任务", index)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
