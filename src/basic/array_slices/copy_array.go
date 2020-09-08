package main
import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := &arr1

	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", *arr2)

	arr1[3] = 800 

	fmt.Println("\narr1: ", arr1)
	fmt.Println("arr2: ", *arr2)

}