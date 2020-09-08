package main

import (
	"fmt"
	"time"
)

func testgo1() {
	for i := 0; i < 5; i++ {
		fmt.Println(" test goroutine1 output", i)
	}
}

func testgo2() {
	for i := 0; i < 5; i++ {
		fmt.Println(" test goroutine2 output", i)
	}
}

func main() {

	go testgo1()
	go testgo2()
	for i := 0; i < 5; i++ {
		fmt.Println("Main func output", i)
	}

	time.Sleep(3000 * time.Millisecond)

}
