package main

import "testing"
import "strings"


/*性能测试结果
$ go test -bench=. -run=none -benchmem
goos: windows
goarch: amd64
pkg: std/my_demo/string
BenchmarkTest-2             1832            676609 ns/op          530338 B/op        999 allocs/op
BenchmarkTest1-2           35190             29221 ns/op            1024 B/op          1 allocs/op
PASS
ok      std/my_demo/string      4.417s

使用预分配内存 性能明显提升

*/



func test() string {
	var s string
	for i := 0; i < 1000; i++ {
        s += "a"
	}

	return s

}

func test1() string {
	s := make([]string, 1000)
	for i:=0;i<1000;i++ {
		s[i] = "a"
	}
	return strings.Join(s, "")
}


func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test1()
	}
}