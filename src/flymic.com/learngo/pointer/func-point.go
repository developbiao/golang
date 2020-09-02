package main

import "fmt"

func main() {
	// 函数默认为指针中是不需要用 *
	a := func1
	fmt.Println(a)
	a1 := func1()
	fmt.Printf("a1的类型： %T, a1的地址:%p, 数值: %v \n", a1, &a1, a1)

	a2 := func2()
	fmt.Printf("a2的类型： %T, a2的地址:%p, 数值: %v \n", a2, &a2, a2)
	fmt.Printf("a2的值为 %p \n", a2)

}

// 普通函数
func func1() []int {
	c := []int{1, 2, 3}
	return c
}

// 指针函数 返回指针
func func2() *[]int {
	c := []int{1, 2, 3, 4}
	fmt.Printf("c的地址为%p: \n", &c)
	return &c
}
