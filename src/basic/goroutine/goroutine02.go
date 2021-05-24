package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Get current GOROOT directory
	fmt.Println("GOROOT:", runtime.GOROOT())
	// Get current os
	fmt.Println("OS:", runtime.GOOS)
	// Get logic CPU number
	fmt.Println("Logic CPU:", runtime.NumCPU())

	// Set maximum CPU core
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)

	// goexit stop current goroutine

	go func() {
		fmt.Println("start...")
		runtime.Goexit() // stop
		fmt.Println("end...")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("main end...")
}
