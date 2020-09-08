package main

func main() {
	// 定义一个缓冲区大小为5的通道
	ch1 := make(chan int, 5)

	ch1 <- 1 // 向缓冲区放入数据
	ch1 <- 2
	ch1 <- 3
	ch1 <- 4
	ch1 <- 5 // 此时缓冲区已满，如果再加入则进入阻塞状态
	ch1 <- 6 // fatal error: all goroutines are asleep - deadlock!

}
