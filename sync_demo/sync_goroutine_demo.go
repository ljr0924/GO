package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {

    var sMap sync.Map
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        sMap.Store(fmt.Sprintf("key%d", i), i)
    }

    wg.Add(2)
    for i := 1; i < 3; i++ {
        go func(index int) {
            name := fmt.Sprintf("goroutine%d", index)
            for j := 0; j < 100; j++ {
                time.Sleep(10)
                key := fmt.Sprintf("key%d", j)
                value, _ := sMap.Load(key)
                fmt.Printf("%s key: %s value: %v\n", name, key, value)
                sMap.Delete(key)
            }
            wg.Done()
        }(i)
    }
    wg.Wait()


}
