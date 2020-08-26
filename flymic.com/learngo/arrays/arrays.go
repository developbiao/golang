package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 3, 5, 7, 11}
	var grid [4][5]int

	fmt.Println("array definitions:")
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr1)
	
}

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
