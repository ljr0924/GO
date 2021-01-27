package main

import (
    "fmt"
    "sync"
    "time"
)


var mtx sync.Mutex

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
				v, ok := LoadAndDelete(&sMap, key)
				if ok {
					fmt.Printf("%s key: %s value: %v\n", name, key, v)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

}

func LoadAndDelete(m *sync.Map, key interface{}) (interface{}, bool) {
	mtx.Lock()
	v, ok := m.Load(key)
	if ok {
		m.Delete(key)
	}
	mtx.Unlock()
	return v, ok

}
