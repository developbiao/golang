package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 2020
	valof := reflect.ValueOf(a) // 先通过reflect.ValueOf 获取反射的值对像
	fmt.Println(valof)
	// 再通过值对像通过类型断言转换为指定类型
	fmt.Println(valof.Interface()) // 转换为interface{} 类型
	fmt.Println(valof.Int())       // 将值以int返回返回

}
