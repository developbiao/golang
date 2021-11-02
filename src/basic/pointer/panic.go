package main

import "fmt"

func main() {
	var nowhere *int
	fmt.Println(nowhere)
	// panic
	fmt.Println(*nowhere)
}
