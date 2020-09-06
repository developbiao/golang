package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:"name"`
	Sex  int    `json:"sex"`
	Age  int    `json:"age"`
}

func main() {
	a := 42

	fmt.Println(reflect.TypeOf(a))  // 通过反射获取变量a的type类型对象
	fmt.Println(reflect.ValueOf(a)) // 通过反射获取变量a的value类型对象

	atype := reflect.TypeOf(a)
	fmt.Println(atype.Name()) // 获取类型名称为int
	avalue := reflect.ValueOf(a)
	fmt.Println(avalue.Int())

	fmt.Println("----------")

	typeofstruct := reflect.TypeOf(person{})
	fmt.Println(typeofstruct.Name()) // 获取反射类型对象 person
	fmt.Println(typeofstruct.Kind()) // 获取反射类型种类 struct

	fieldnum := typeofstruct.NumField() // 获取结构体成员字段的数量

	for i := 0; i < fieldnum; i++ {
		fieldname := typeofstruct.Field(i) // 索引对应字段信息
		fmt.Println(fieldname)
		name, err := typeofstruct.FieldByName("Name") // 根据字符串返回对应的字段信息
		fmt.Println(name, err)
	}

}
