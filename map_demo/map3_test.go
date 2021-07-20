package map_demo

import (
    "sync"
    "testing"
    "time"
)

func TestMapReadWrite(t *testing.T) {

    m := make(map[string]int)

    // 运行时会对字典并发操作做出检测，如果存在某个任务正在对字典进行写操作
    // 那么其他任务就不能对该字典执行并发操作（读，写，删除）
    // 否则会导致进程崩溃

    // 写操作
    go func() {
        for {
            m["a"] = 1
            time.Sleep(time.Microsecond)
        }
    }()

    // 读操作
    go func() {
        for true {
            _ = m["a"]
            time.Sleep(time.Microsecond)
        }
    }()

    select{}

}

func TestMapLock(t *testing.T) {

    var lock sync.RWMutex   // 使用读写锁 获取最佳性能
    m := make(map[string]int)

    // 写操作
    go func() {
        for {
            lock.Lock()
            m["a"] = 1
            lock.Unlock()
            time.Sleep(time.Microsecond)
        }
    }()

    // 读操作
    go func() {
        for {
            lock.RLock()
            _ = m["a"]
            lock.RUnlock()
            time.Sleep(time.Microsecond)
        }
    }()

}
