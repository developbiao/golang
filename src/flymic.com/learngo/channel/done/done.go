package main

import (
	"fmt"
)

func doWorker(id int, c chan int, done chan bool) {

	go func() {
		for n := range c {
			fmt.Printf("Woker %d received %c\n", id, n)
			done <- true
		}
	}()
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	// send
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	// received
	for _, worker := range workers {
		<-worker.done
	}

	// send
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// received
	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
}
