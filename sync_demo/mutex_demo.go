package main

import (
	"fmt"
	"sync"
)


func main (){
	// 同步 Map
	syncMap := sync.Map{}

	// 存储key value
	syncMap.Store("name", "ljr")

	// 获取key value
	name, ok := syncMap.Load("name")
	if !ok {
		fmt.Println("不存在key：name")
	} else{
		fmt.Printf("key %s value %s\n", "name", name)
	}

	// 如果存在key 返回value
	// 否则存储新值，并返回
	name, loaded := syncMap.LoadOrStore("name", "msh")
	if loaded {
		fmt.Println("成功获取name")
	}
	fmt.Printf("key: name value: %s\n", name)

	// 删除key
	fmt.Println("删除key value")
	syncMap.Delete("name")
	name, ok = syncMap.Load("name")
	if !ok {
		fmt.Println("不存在key：name")
	} else{
		fmt.Printf("key %s value %s\n", "name", name)
	}

	// 遍历key value
	syncMap.Store("name", "ljr")
	syncMap.Store("sex", 1)
	syncMap.Store("birthday", "1996/09/24")

	fmt.Println("遍历key value")
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key %s value %v\n", key, value)
		return true
	})

}
