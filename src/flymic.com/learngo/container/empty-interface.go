package main

import "fmt"

// 空接口
type T interface{}

func foo(t T) {
	fmt.Println(t)
}

func bar(t interface{}) {
	fmt.Println(t)
}

func main() {
	foo(0)           // 使用int类型做为参数传入到函数
	foo("Hello美女老师") // 使用string类型做为参数传入到函数

	bar(7)
	bar("Look")
}
