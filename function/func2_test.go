package function

import "fmt"
import "testing"

func TestFunc2(* testing.T) {

    // 可变长参数需要传入切片类型
    a1 := []int{1, 2, 3}
    test(a1...)
    // 如果类型为数组 需要转成切片类型
    a2 := [3]int{1, 2, 3}
    test(a2[:]...)	

    // 值拷贝问题  只会拷贝最外层值
    a3 := []int{1, 2, 3}
    fmt.Printf("传入前 %v\n", a3)
    test1(a3...)	
    fmt.Printf("处理后 %v\n", a3)
    // 值拷贝问题  只会拷贝最外层值
    a4 := []int{1, 2, 3}
    fmt.Printf("传入前 %v\n", a4)
    test2(a4...)	
    fmt.Printf("处理后 %v\n", a4)
}

/**
可变长参数
**/
func test(a... int) {

    for _, v := range a{
        fmt.Println(v)
    }

}

func test1(a... int) {

    for i := 0; i<len(a); i++ {
        a[i] += 10
    }

}

func test2(a... int) {
    dst := []int{}
    // 重新copy一份 不影响原来的值
    copy(dst, a)
    for i := 0; i<len(dst); i++ {
        dst[i] += 10
    }

}