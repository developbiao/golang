package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go Relief1()
	go Relief2()
	go Relief3()
	wg.Wait() // 主goroutine 进入阻塞状态
}

func Relief1() {
	fmt.Println("func1...")
	wg.Done() // 执行完成 同步等待数量减少 1
}

func Relief2() {
	defer wg.Done()
	fmt.Println("func2...")
}

func Relief3() {
	fmt.Println("func3...")
	wg.Done()
}
