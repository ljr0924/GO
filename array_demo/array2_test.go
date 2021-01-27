package array_demo

import "testing"

func TestArray2(t *testing.T) {

    // 数组初始化，指定长度
    arr1 := [5]int{1,2,3,4,5}
     t.Logf("arr1 length %d  capcity %d\n", len(arr1), cap(arr1))
    
    // 自动识别长度
    arr2 := [...]int{1,2,3,4,5}
     t.Logf("arr2 length %d  capcity %d\n", len(arr2), cap(arr2))

    // 数组访问
     t.Log("下标访问数组，", arr1[1])
    
    // range循环
     t.Log("range循环方式")
    for index, value := range arr1 {
         t.Logf("index: %d  value: %d\n", index, value)
    }

     t.Log("下标循环方式")
    // 下标循环
    for i := 0; i < len(arr1); i++ {
         t.Logf("index: %d  value: %d\n", i, arr1[i])
    }

}