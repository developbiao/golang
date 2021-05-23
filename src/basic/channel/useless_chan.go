package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Example 01
	// 只能发不能收通道
	var uselessChan = make(chan<- int, 1)

	// 只能收不能发的通道
	var anotherUselessChan = make(<-chan int, 1)

	fmt.Printf("The useless channles: %v, %v\n", uselessChan, anotherUselessChan)

	// Example 02
	intChan1 := make(chan int, 3)
	SendInt(intChan1)

	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	_ = GetIntChan(getIntChan)

}

// Example02
func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

type GetIntChan func() <-chan int
