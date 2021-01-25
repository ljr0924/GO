package main

import "fmt"


/*
字符串是一个不可变字节序列，是一个复合数据结构

默认使用utf-8编码存储unicode字符
*/
func main() {

	name := "香蕉、\x31\x32\x33\x34\x35"
	fmt.Printf("%s\n", name)
	fmt.Printf("% x  长度%d\n", name, len(name))

	name1 := name[2:]
	fmt.Printf("%s\n", name1)
	fmt.Printf("% x  长度%d\n", name, len(name1))

	// 字符串零值为 ""
	var name2 string
	fmt.Println("字符串零值： ", name2)

	// `` 反引号字符串为不转义字符串，支持跨行
	name3 := `香蕉
	banana`
	fmt.Printf("字符串：%s\n原始形状 %v", name3, name3)

}