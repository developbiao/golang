package main

import "fmt"

func main() {
	var slice []int
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)

	// use make function create slice
	s1 := make([]int, 0, 5)
	fmt.Println(s1)
	s1 = append(s1, 1, 2)
	fmt.Println(s1)

	// auto increase cap
	s1 = append(s1, 3, 4, 5, 6, 7)
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))

	// Add slice to other slice
	s2 := make([]int, 0, 3)
	s2 = append(s2, s1...)
	fmt.Println(s2)

	// test new slice
	slice1 := new([]int)
	fmt.Println(slice1)

	// test make slice
	slice2 := make([]int, 5)
	fmt.Println(slice2)

}
