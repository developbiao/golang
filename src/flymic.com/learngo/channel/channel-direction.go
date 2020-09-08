package main

import "fmt"

func main() {
	ch1 := make(chan int)   // 双向通道
	ch2 := make(chan<- int) // 只写通道
	ch3 := make(<-chan int) // 只读通道

	// ========[1]=========
	// 如果创建的时候创建的就是双向通道
	// 则在子协程内部写入数据，读取是不受影响的
	go writeOnly(ch1)
	data2 := <-ch1
	fmt.Println("获取到只写通道中的数据是:", data2)

	// ========[2]=========
	// 如果将定向通道ch2只写通道,作为参数传递。
	// 则不能读取到写回来的数据。
	go writeOnly(ch2)

	go readOnly(ch1) // 这里只可以传入ch1 双向通道
	ch1 <- 20        // 向ch1中写入数据

	// ========[3]=========
	go readOnly(ch3) // 传递单向通道ch3 就无法向通道中写入数据

	fmt.Println("Over!")

}

func readOnly(ch <-chan int) {
	data := <-ch
	fmt.Println("读取到通道的数据是: ", data)
}

func writeOnly(ch chan<- int) {
	// 如果传入进来的是双向通道
	// 但是函数本身接收的是一个只写通道，则在此函数只允许写入不允许读取
	ch <- 10
	fmt.Println("只写通道结束")

}
