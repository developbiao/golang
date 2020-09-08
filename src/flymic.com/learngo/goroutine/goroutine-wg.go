package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(3) // 设置同步等待组的数量
	go relief1()
	go relief2()
	go relief3()
	wg.Wait() // 主gorountine进入阻塞状态
	fmt.Println("Main end...")

}

func relief1() {
	fmt.Println("fun1...")
	wg.Done() // 执行完成同步等待数量减1
}

func relief2() {
	fmt.Println("fun2...")
	wg.Done()
}
func relief3() {
	fmt.Println("fun3...")
	wg.Done()
}
