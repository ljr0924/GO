package reflect_demo

import (
    "fmt"
    "reflect"
    "testing"
)

type X int

func TestReflect1(t *testing.T) {

    var a X = 100

    ta := reflect.TypeOf(a)
    fmt.Println(ta.Name(), ta.Kind())

}
