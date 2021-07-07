package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Prograning is running...")
	go func() {
		fmt.Println("I'am a wild goroutine😈")
		panic("Oh No~ programing is crashed WTF?")
	}()

	time.Sleep(3 * time.Second)
}
