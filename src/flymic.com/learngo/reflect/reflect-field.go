package main

import (
	"fmt"
	"reflect"
)

type haojiahuo struct {
	Name string
	Age  int
}

func main() {
	h := haojiahuo{"好家伙", 20}
	fmt.Println(h)
	hofvalue := reflect.ValueOf(h) // 获取结构体reflect.Value对象
	for i := 0; i < hofvalue.NumField(); i++ {
		// 获取结构体内的索引为i的字段值
		fmt.Println(hofvalue.Field(i).Interface())
	}
	fmt.Println(hofvalue.Field(1).Type()) // 获取结构体内索引为1的字段的类型

}
