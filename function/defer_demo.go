package main

import "fmt"


/*
defer 遵循先进后出原则
*/

func main() {
	// 先输出b 再输出a
	defer fmt.Println("a")
	defer fmt.Println("b")
}