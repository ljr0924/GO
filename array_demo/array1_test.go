package array_demo

import "testing"

func TestArray1(t *testing.T) {

    // 初始化
    var a [4]int
     t.Logf("a: %v\n", a)

    // 指定长度
    b := [4]int{1,2}
     t.Logf("b: %v\n", b)

    // 指定索引位置初始化
    c := [4]int{1, 3:2}
     t.Logf("c: %v\n", c)

    // 长度自适应
    d := [...]int{1, 3:2}
     t.Logf("d: %v\n", d)
    e := [...]int{1,2,3}
     t.Logf("e: %v\n", e)

    // 结构体数组初始化可省略名字
    type user struct {
        name string
        age int
    }
    f := [...]user{
        {"p1", 18},
        {"p2", 19},
    }
     t.Logf("f: %v\n", f)
     t.Logf("f: %+v\n", f)
     t.Logf("f: %#v\n", f)

    // 多维数组 初始化
    g := [2][2]int{
        {1, 2},
        {3, 4},
    }
     t.Logf("g: %v\n", g)

    // 多维数组 自适应长度
    h := [...][2]int{
        {1, 2},
        {3, 4},
    }
     t.Logf("h: %v\n", h)

    // 数组地址

}