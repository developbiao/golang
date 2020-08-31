package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main () {
	f := adder()
	for i := 0; i < 10; i ++ {
		fmt.Println(f(i))
	}
	fmt.Println("ok")
}