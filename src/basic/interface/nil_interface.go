package main

import "fmt"

func main() {
	var v interface{}

	fmt.Printf("%T, %v %v\n", v, v, v == nil)
}
