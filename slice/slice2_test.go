package slice

import "fmt"
import "testing"

func TestSlice2(t *testing.T) {

    // 不指定切片长度和容量
    s1 := []int{1,2,3,4,5}
    fmt.Printf("s1: %v\n", s1)
    fmt.Printf("切片长度：%d，容量：%d\n", len(s1), cap(s1))

    // 使用make指定切片长度和容量
    s2 := make([]int, 10, 15)
    fmt.Printf("s2: %v\n", s2)
    fmt.Printf("切片长度：%d，容量：%d\n", len(s2), cap(s2))

    // 字符串类型转切片
    str := "hello world"
    str1 := []byte(str)
    fmt.Printf("转换成[]byte类型后， %v\n", str1)

    str2 := []rune(str)
    fmt.Printf("转换成[]rune类型后， %v\n", str2)
}
