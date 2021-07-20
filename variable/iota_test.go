package variable

import (
    "fmt"
    "testing"
)

const (
    var1 = iota
    var2
    var3
    var4
)

func TestIota(t *testing.T) {
    fmt.Println("var1 = ", var1)
    fmt.Println("var2 = ", var2)
    fmt.Println("var3 = ", var3)
    fmt.Println("var4 = ", var4)
}
