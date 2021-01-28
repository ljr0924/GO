package  method_demo

import "sync"
import "testing"
import "fmt"

type MySyncMap struct {
    sync.Mutex
    myMap map[interface{}]interface{}
}

func NewMySyncMap() *MySyncMap {
    return &MySyncMap{
        myMap: map[interface{}]interface{}{},
    }
}

func (m *MySyncMap) Store(k, v interface{}) {
    m.Lock()
    m.myMap[k] = v
    m.Unlock()
}

func (m *MySyncMap) Delete(k interface{}) {
    m.Lock()
    delete(m.myMap, k)
    m.Unlock()
}

func (m *MySyncMap) Load(k interface{}) (interface{}, bool) {
    m.Lock()
    v, ok := m.myMap[k]
    m.Unlock()
    return v, ok
}

func TestMySyncMap(t *testing.T) {
    m := NewMySyncMap()
    m.Store("name", "m")
    fmt.Printf("%+v\n", m)
    if v, ok := m.Load("name");ok {
        fmt.Printf("name: %s\n", v)
    } else {
        fmt.Print("获取name失败\n")
    }

    if v, ok := m.Load("age");ok {
        fmt.Printf("age: %d\n", v)
    } else {
        fmt.Print("获取age失败\n")
    }

    m.Store("age", 18)
    fmt.Println("存储后重新获取")
    if v, ok := m.Load("age");ok {
        fmt.Printf("age: %d\n", v)
    } else {
        fmt.Print("获取失败\n")
    }

}