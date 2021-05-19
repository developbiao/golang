package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	// Sender
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	// Receiver
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: receive an elemnt: %v\n", elem)
	}

	fmt.Println("End.")
}
