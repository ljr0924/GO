package interface_demo

import "testing"

type tester interface {
    test()
    string() string
}

type data struct {}

func (*data) test() {}

func (data) string() string {
    return ""
}

func TestInterface1(t *testing.T) {
    var d data

    // var t tester = d  // data does not implement tester (test method has pointer receiver)

    var mt tester = &d

    mt.test()
    mt.string()

}

func TestInterface2(t *testing.T) {
    var t1, t2 interface{}
    // 空接口的默认值为nil
    t.Log(t1 == nil, t2 == nil)

    t1 = 100
    t2 = 100

    t.Log(t1 == t2)

}