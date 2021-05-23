package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	example01()
	example02()
}

func example01() {
	// Prepare some channel
	intChannles := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	// Random chose one channel, send element for it
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannles[index] <- index

	// Which one channel exists can fetch element, whice branch will be exected
	select {
	case <-intChannles[0]:
		fmt.Println("The first candiate case is selected.")
	case <-intChannles[1]:
		fmt.Println("The second  candiate case is selected.")
	case elem := <-intChannles[2]:
		fmt.Printf("The third candidate case is selected, the element is %d\n", elem)
	default:
		fmt.Println("No candiate case is selected")
	}
}

// Example 02
func example02() {
	intChan := make(chan int, 1)
	// One second close channle
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})

	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}
