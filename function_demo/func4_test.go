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
