package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明变量a
	a := 100
	fmt.Printf("a的内存地址为：%p\n", &a)

	// 获取变量a的反射类型reflect.Value地址
	rf := reflect.ValueOf(&a)
	fmt.Println("通过反射获取变量a的地址:", rf)

	// 获取变量a地址的值
	rval := rf.Elem()
	fmt.Println("反射a的值: ", rval)

	// 修改a的值
	rval.SetInt(200)
	fmt.Println("修改之后反射类型的值为:", rval.Int())

	// 原始值也被修改了
	fmt.Println("原始a的值也被修改为: ", a)
}
