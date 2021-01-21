package main

import "fmt"

func main() {
    funcWithoutReturn("hello", "world")
    a1, a2 := funcMultiReturn("hello", "world")
    fmt.Printf("funcMultiReturn returns: [%s, %s]\n", a1, a2)
	fmt.Println("add(1, 2) => ", add(1, 2))
	fmt.Println("sub(1, 2) => ", sub(1, 2))
	fmt.Println("mul(1, 2) => ", mul(1, 2))
    fmt.Println("div(1, 2) => ", div(1, 2))
    fmt.Println("div(3, 2) => ", div(3, 2))
    fmt.Println("div(1, 2) => ", div(1, 0))
    
}

func add(num1, num2 int) int {
    return num1 + num2
}

func sub(num1, num2 int) int {
    return num1 - num2
}

func mul(num1, num2 int) int {
    return num1 * num2
}

func div(num1, num2 float64) float64 {
    if num2 == 0 {
        panic("分母不能为零")
    }
    return num1 / num2
}

func funcWithoutReturn (a1, a2 string) {

    fmt.Printf("arg1:%s arg2:%s\n", a1, a2)

}

func funcMultiReturn (a1, a2 string) (string, string) {

    fmt.Printf("arg1:%s arg2:%s\n", a1, a2)
    return a1, a2

}