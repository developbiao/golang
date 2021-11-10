package main

import (
	"fmt"
	"time"
)

func main() {

	go sleepGohper()
	time.Sleep(4 * time.Second)
	fmt.Println("Everything is ok")
}

func sleepGohper() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}
