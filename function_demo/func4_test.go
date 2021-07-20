package function_demo

import (
    "fmt"
    "testing"
)

func Print(args ...interface{}) {
    fmt.Printf("%+v\n", args...)
}

func TestPrint(t *testing.T) {
    Print([]int{1,2,3,4,5})
    Print(map[string]interface{}{
        "name":"ljr",
        "age": 18,
    })
}

func Fix(p [2]int) {
    p[1] = 2
    fmt.Println(&p)
    fmt.Println(p)
}

func TestArg(t *testing.T) {

    p := [2]int{}
    fmt.Println(&p)
    fmt.Println(p)
    Fix(p)
    fmt.Println(&p)
    fmt.Println(p)
}
