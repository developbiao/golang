package main

import "fmt"

func main() {
	// 创建一个普通的数组
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// 创建一个指针用来存储数组的地址 即: 数组指针
	var p *[3]int
	p = &arr // 将数组arr的地址，存在到数组指针p上
	fmt.Println(p)

	fmt.Println(*p)
	fmt.Println(&p) // 该指针自己的地址

	// 修改数组指针中的数据
	(*p)[0] = 300
	fmt.Println(arr)

	// 简化写法
	p[1] = 304
	fmt.Println(arr)

}
