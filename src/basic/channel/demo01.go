package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Println("The first element recevied from channel ch1: %v\n", elem1)
	elem2 := <-ch1
	fmt.Println("The second element recevied from channel ch1: %v\n", elem2)
}
