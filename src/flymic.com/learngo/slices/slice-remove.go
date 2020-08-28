package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	// remove first element method 1
	// slice = slice[1:]
	// fmt.Println(slice)

	// remove fist element method 2 not change address
	slice = append(slice[:0], slice[1:]...)
	fmt.Println(slice)

	// remove last element
	slice = []int{1, 2, 3, 4}
	slice = slice[:len(slice)-1]
	fmt.Println(slice)

}
