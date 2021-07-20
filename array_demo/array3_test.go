package array_demo

import (
	"fmt"
	"testing"
)

func TestArrayPtr(t *testing.T) {

	// 字符指针数组，存储字符串类型指针的数组
	stringPtrArray := [2]*string{}

	// 数组指针
	p := &stringPtrArray

	fmt.Printf("%T\n", stringPtrArray)
	fmt.Printf("%T\n", p)

	s1 := "123"
	stringPtrArray[1] = &s1

	fmt.Printf("%v", stringPtrArray)

}