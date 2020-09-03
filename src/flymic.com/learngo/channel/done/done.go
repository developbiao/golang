package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, c chan int, wg *sync.WaitGroup) {

	go func() {
		for n := range c {
			fmt.Printf("Woker %d received %c\n", id, n)
			wg.Done()
		}
	}()
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w.in, w.wg)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	// send
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	// send
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()

}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
}
