package main

import (
	"fmt"
	"reflect"
)

type person2 struct {
	Name string `json:"name"`
	Sex  int    `json:"sex"`
	Age  int    `json:"age"`
}

func main() {
	typeofstruct := reflect.TypeOf(&person2{})
	fmt.Println(typeofstruct.Elem().Name()) // 获取指针类型指向的元素类型的名称
	fmt.Println(typeofstruct.Elem().Kind()) // 获取指针类型指向的元素的种类

}
