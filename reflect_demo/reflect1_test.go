package reflect_demo

import (
	"fmt"
	"reflect"
	"testing"
)

type X int

type Order struct {
	Id   int    `json:"id"`
	Desc string `json:"desc"`
}

func TestReflect1(t *testing.T) {

	var a X = 100

	ta := reflect.TypeOf(a)
	va := reflect.ValueOf(a)

	fmt.Println(ta.Name(), ta.Kind(), ta.String(), va)

	var b = &Order{
		Id:   1,
		Desc: "desc",
	}

	tb := reflect.TypeOf(b)
	vb := reflect.ValueOf(b)

	fmt.Println(tb.Name(), tb.Kind(), tb.String(), tb.FieldAlign(), vb)

}
