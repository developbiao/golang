package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Prevent off by one error
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)
}



