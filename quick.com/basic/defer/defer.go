package main

import "fmt"

func main() {
	tryDefer()
	fmt.Println("New project can be running?")
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
