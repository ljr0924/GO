package slice_demo

import (
    "fmt"
    "testing"
)

func TestSlice3(t *testing.T) {

	a := make([]int, 1)
	b := make([]int, 1)

	// fmt.Println(a == b)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
}