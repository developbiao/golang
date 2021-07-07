package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Wild goroutine demo running...")
	// Anywhere make wild goroutine
	Go(func() {
		// Assume goroutine is crashed
		panic("Wild goroutine is crashed! 💣")
	})

	// Prevent main goroutine is end waiting 3 seconds
	time.Sleep(3 * time.Second)

}

// Delegation wild goroutine we can caputre exception
func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		f()
	}()
}
