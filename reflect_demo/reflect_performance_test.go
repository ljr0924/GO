package reflect_demo

import (
    "reflect"
    "testing"
)


type Data struct {
    X int
}

var d = new(Data)

func set (x int) {
    d.X = x
}

func reflectSet(x int) {
    v := reflect.ValueOf(d).Elem()
    f := v.FieldByName("X")
    f.Set(reflect.ValueOf(x))
}

/*
BenchmarkSet
BenchmarkSet-8    	1000000000	         0.366 ns/op
*/
func BenchmarkSet(b *testing.B) {
    for i := 0; i < b.N; i++ {
        set(100)
    }
}

/*
BenchmarkRSet
BenchmarkRSet-8   	12136863	        91.2 ns/op
*/
func BenchmarkRSet(b *testing.B) {
    for i := 0; i < b.N; i++ {
        reflectSet(100)
    }
}
