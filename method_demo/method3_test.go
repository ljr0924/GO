package method_demo

import "testing"
import "fmt"
import "reflect"

type S struct {}

type T struct {
    S
}

// 函数名首字母要大写才能被识别
func (S) SVal() {}
func (*S) SPtr() {}
func (T) TVal() {fmt.Println("tVal")}
func (*T) TPtr() {}

func MethodSet(a interface{}) {
    t := reflect.TypeOf(a)
    fmt.Println("Type: ", t.String())
    fmt.Println("Method num: ", t.NumMethod())
    for i, n := 0, t.NumMethod(); i < n; i++ {
        m := t.Method(i)
        fmt.Println(m.Name, m.Type)
    }

}

func TestCombineMethod(t *testing.T) {
    var mt T 
    MethodSet(mt)
    fmt.Println("------------------")
    MethodSet(&mt)
}