package slice

import "testing"
import "fmt"

func TestSlice1(t *testing.T) {

    // slice初始化
    s1 := []int{1,2,3,4,5}
    fmt.Printf("type: %T\n", s1)

    fmt.Printf("追加元素前：%v\n", s1)
    // 追加元素
    s1 = append(s1, 6, 7, 8)
    fmt.Printf("追加元素后：%v\n", s1)

    // 删除第一个元素
    s1 = s1[1:]
    fmt.Printf("删除第一个元素后：%v\n", s1)

    // 删除最后n个元素
    s1 = s1[:len(s1)-2]
    fmt.Printf("删除后面两个元素后：%v\n", s1)

    // 删除指定位置元素
    // s1 = append(s1[:i], s1[i+1:])
    s1 = append(s1[:3], s1[4:]...)
    fmt.Printf("删除中间元素后：%v\n", s1)
}
