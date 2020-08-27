package main

import "fmt"

func main() {
	s1 := make([]int, 0, 3)
	fmt.Printf("Address: %p, len: %d, Cap:%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 1, 2)
	fmt.Printf("Address: %p, len: %d, Cap:%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 3, 4, 5)
	fmt.Printf("Address: %p, len: %d, Cap:%d\n", s1, len(s1), cap(s1))

	s2 := []int{1, 2, 3, 4}
	s3 := []int{7, 8, 9}

	// copy s3 to s2
	//copy(s2, s3)
	//fmt.Println(s2)
	//fmt.Println(s3)

	//copy(s3, s2[2:])
	//fmt.Println(s2)
	//fmt.Println(s3)

	// copy s2 to s3
	copy(s3, s2)
	// [1, 2 ,3]
	fmt.Println(s2)
	fmt.Println(s3)

}
