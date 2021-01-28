package method_demo

import "testing"
import "fmt"


type N int 

func (n N) Add() {
    n++
}

type People struct {
    name string
}

func (p *People) SayHi()  {
    fmt.Println("大家好，我是"+p.name)
}


func TestMethod1(t *testing.T) {
    var n1 N
    n1 = 1
    n1.Add()
    t.Log(n1)

    p1 := &People{"p1"}
    p1.SayHi()
}