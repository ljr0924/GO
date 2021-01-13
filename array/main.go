package main

import "fmt"

func main() {
	// 数组初始化，指定长度
	arr1 := [5]int{1,2,3,4,5}
	fmt.Printf("arr1 length %d  capcity %d\n", len(arr1), cap(arr1))
	
	// 自动识别长度
	arr2 := [...]int{1,2,3,4,5}
	fmt.Printf("arr2 length %d  capcity %d\n", len(arr2), cap(arr2))

	// 数组访问
	fmt.Println("下标访问数组，", arr1[1])
	
	// range循环
	fmt.Println("range循环方式")
	for index, value := range arr1 {
		fmt.Printf("index: %d  value: %d\n", index, value)
	}

	fmt.Println("下标循环方式")
	// 下标循环
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("index: %d  value: %d\n", i, arr1[i])
	}

}