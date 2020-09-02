package main

import "fmt"

func main() {
	res := closure()
	fmt.Println(res)
	r1 := res()
	fmt.Println(r1)
	r2 := res()
	fmt.Println(r2)

	res2 := closure()
	fmt.Println(res2)
	r3 := res2()
	fmt.Println(r3)
}

func closure() func() int {
	a := 0

	return func() int {
		a++
		return a
	}
}
