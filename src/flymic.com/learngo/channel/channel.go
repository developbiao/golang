package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Woker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	var channles [10]chan<- int
	for i := 0; i < 10; i++ {
		channles[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channles[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channles[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	time.Sleep(time.Millisecond)
}
