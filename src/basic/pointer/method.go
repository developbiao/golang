package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	p.age++
}

func main() {
	terry := &person{
		name: "Terry",
		age:  15,
	}
	fmt.Printf("%+v\n", terry)
	gongbiao := person{
		name: "GongBiao",
		age:  17,
	}
	// (&gongbiao).birthday()
	// Go 语言在变量通过点标记调用方法的时候会自动使用&取得变量的内存地址，所以就算不写出(&gongbiao).birthday(),代码可以正常运行
	gongbiao.birthday()
	fmt.Printf("%v\n", gongbiao)
}
