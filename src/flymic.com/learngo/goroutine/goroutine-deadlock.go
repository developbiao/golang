package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(4) // 设置同步等待组的数量
	go sale1()
	go sale2()
	go sale3()
	wg.Wait()

}

func sale1() {
	fmt.Println("send sale01")
}

func sale2() {
	fmt.Println("send sale02")
}
func sale3() {
	fmt.Println("send sale03")
}
