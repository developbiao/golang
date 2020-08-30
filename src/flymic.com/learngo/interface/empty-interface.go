package main

import "fmt"

func main() {
	var i interface{}

	i = 42
	describe(i)

	i = "you are"
	describe(i)

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
