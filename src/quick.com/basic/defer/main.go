package main

import (
	"fmt"
	"quick.com/basic/defer/packagetest"
)

func main() {
	tryDefer()
	fmt.Println("New project can be running?")
	packagetest.SayHello()

}

func tryDefer() {
	for i := 0; i < 6; i++ {
		defer fmt.Println(i)
		if i == 3 {
			// how it works with defer
			panic("printed too many")
		}

	}
}
