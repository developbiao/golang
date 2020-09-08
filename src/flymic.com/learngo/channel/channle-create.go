package main

import "fmt"

func main() {
	var channel chan int
	fmt.Printf("通道的数据类型：%T, 通道的值:%v\n", channel, channel)

	if channel == nil {
		channel = make(chan int)
		fmt.Printf("通过make创建的通道数据类型:%T, 通道的值:%v\n", channel, channel)
	}

	go func() {
		fmt.Println("===== 子协程执行 =====")
		data := <-channel
		fmt.Println("读取到时通道中的数据是：", data)
	}()

	channel <- 3 // 往通道里放数据
	fmt.Println("===== 主协程结束 =====")

}
