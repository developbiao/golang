package main

import (
	"fmt"
	"time"
)

func main() {
	go testgo1()
	go testgo2()

	for i := 0; i <= 5; i++ {
		fmt.Println("main function exuecte", i)
	}

	time.Sleep(3000 * time.Millisecond) // 加上休眠让主程序休眠3秒钟
	fmt.Println("Main function is over!")
}

func testgo1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Test goroutine1:", i)
	}
}

func testgo2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Test goroutine2:", i)
	}
}
