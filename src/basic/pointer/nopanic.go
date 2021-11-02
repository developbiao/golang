package main

import "fmt"

func main() {
	var nowhere *int
	fmt.Println(nowhere)
	// panic
	if nowhere != nil {
		fmt.Println(*nowhere)
	}
}
