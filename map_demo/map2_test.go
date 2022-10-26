package map_demo

import "testing"

func TestMapRange(t *testing.T) {

    m := make(map[string]int)

    for i := 0; i < 10; i++ {
        m[string(rune('a'+i))] = i
    }

    // 每次打印顺序不一样
    for i := 0; i < 4; i++ {
        for k, v := range m {
            t.Logf("k: %s  v: %d", k, v)
        }
        t.Log("------------------------------")
    }

    // 使用len获取map的元素个数
    t.Logf("map 长度：%d", len(m))

}


type user struct {
    name string
    age byte
}

/*
=== RUN   TestMapStruct
    map2_test.go:36: m1: map[1:{name:m1 age:1}]
    map2_test.go:41: m1: map[1:{name:m1 age:2}]
    map2_test.go:47: m2: map[1:0xc00000c220]
    map2_test.go:49: m2: map[1:0xc00000c220]
--- PASS: TestMapStruct (0.00s)
PASS
*/
func TestMapStruct(t *testing.T) {
    m1 := map[int]user {
        1: {"m1", 1},
    }
    t.Logf("m1: %+v", m1)

    u := m1[1]
    u.age += 1
    m1[1] = u
    t.Logf("m1: %+v", m1)

    m2 := map[int]*user {
        1: {"m2", 1},
    }

    t.Logf("m2: %+v", m2)
    m2[1].age++
    t.Logf("m2: %+v", m2)
}

/*
=== RUN   TestMapInit
    map2_test.go:57: false true
--- PASS: TestMapInit (0.00s)
PASS
*/
func TestMapInit(t *testing.T) {

    m1 := make(map[int]string)
    var m2 map[int]string

    t.Log(m1 == nil, m2 == nil)

}

/*
=== RUN   TestMapDeleteInRange
    map2_test.go:90: 删除key：9
    map2_test.go:90: 删除key：2
    map2_test.go:90: 删除key：3
    map2_test.go:90: 删除key：4
    map2_test.go:90: 删除key：5
    map2_test.go:90: 删除key：6
    map2_test.go:90: 删除key：7
    map2_test.go:90: 删除key：0
    map2_test.go:90: 删除key：1
    map2_test.go:90: 删除key：8
    map2_test.go:93: map[100:1000]
--- PASS: TestMapDeleteInRange (0.00s)
PASS
*/
func TestMapDeleteInRange(t *testing.T) {

    m := make(map[int]int)

    for i := 0; i < 10; i++ {
        m[i] = i
    }

    // 遍历的时候删除或新值键值对是安全的
    for k := range m {
        if k == 5 {
            m[100] = 1000
        }
        delete(m, k)
        t.Logf("删除key：%d", k)
    }

    t.Logf("%+v", m)

}