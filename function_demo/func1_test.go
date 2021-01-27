package function_demo

import "fmt"
import "testing"


func TestFunc1(t *testing.T) {

    // 无返回值函数
    funcWithoutReturn("hello", "world")
    // 多个返回值函数
    a1, a2 := funcMultiReturn("hello", "world")
    fmt.Printf("funcMultiReturn returns: [%s, %s]\n", a1, a2)

    // 匿名函数
    lambdaFunc := func (n1, n2 int)  {
        fmt.Printf("这个是匿名函数 func(%d, %d)\n", n1, n2)
    }
    lambdaFunc(1, 2)

    a3 := 1
    fmt.Println("内存地址为 ", &a3)
    testCopy(a3)
    a4 := &a3
    fmt.Println("内存地址为 ", &a4)
    testCopy(a4)

    fmt.Println("add(1, 2) => ", add(1, 2))
    fmt.Println("sub(1, 2) => ", sub(1, 2))
    fmt.Println("mul(1, 2) => ", mul(1, 2))
    fmt.Println("div(1, 2) => ", div(1, 2))
    fmt.Println("div(3, 2) => ", div(3, 2))
    fmt.Println("div(1, 2) => ", div(1, 0))
    
}

/*
函数传参均为值传递 传指针的话 拷贝指针  传目标值的话 拷贝目标值
传指针会比较节约内存
*/
func testCopy(arg interface{}) {
    fmt.Println("内存地址为 ", &arg)
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
        fmt.Println("分母不能为零")
        return 0
    }
    return num1 / num2
}


// 无返回值
func funcWithoutReturn (a1, a2 string) {

    fmt.Printf("arg1:%s arg2:%s\n", a1, a2)

}

// 多个返回值
func funcMultiReturn (a1, a2 string) (string, string) {

    fmt.Printf("arg1:%s arg2:%s\n", a1, a2)
    return a1, a2

}