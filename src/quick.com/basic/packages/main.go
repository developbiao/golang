package main

import (
	"fmt"
)

func main() {

	//packagetest.TestStrings()
	//packagetest.TestTime()
	//packagetest.TestTimestamp()
	func () {
		fmt.Println("立即运行的匿名函数")
	}()

	func(a int, b int) {
		fmt.Println(a, b)
	}(1, 2)

	res := func(a int, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(res)

	fmt.Println("Let running golang")
}

