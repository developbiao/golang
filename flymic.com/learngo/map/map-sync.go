package main

import (
	"fmt"
	"sync"
)

var syncmap sync.Map

func main() {
	// store key-value to sync.Map
	syncmap.Store("gongbiao", 100)
	syncmap.Store("lijujn", 97)
	syncmap.Store("LiYa", 98)

	// load method get sync.Map value by key
	fmt.Println(syncmap.Load("gongbiao"))

	// delete
	syncmap.Delete("gongbiao")

	syncmap.Store("huanghuang", 68)

	// traversing
	syncmap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}
