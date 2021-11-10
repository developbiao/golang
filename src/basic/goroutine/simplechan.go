package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		// Receive from goroutinea gopher ID
		gopherID := <-c
		fmt.Println("goher ", gopherID, "has finsihed sleeping")
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, " snore ....")
	// Send gopher ID to main function
	c <- id
}
