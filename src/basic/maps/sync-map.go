package main

import (
	"fmt"
	"sync"
)

var syncmap sync.Map

func main() {
	// Store 方法将键值对保存到sync.Map
	syncmap.Store("zhangsan", 97)
	syncmap.Store("gongbiao", 93)
	syncmap.Store("liuzhehui", 93)

	// Load 方法获取sync.Map 键所对应的值
	fmt.Println(syncmap.Load("liuzhehui"))

	// Range 遍历所有sync.Map中的键值对
	syncmap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}
