package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 获取当前GOROOT的目录
	fmt.Println("GOROOT:", runtime.GOROOT())

	// 获取当前操作系统
	fmt.Println("操作系统:", runtime.GOOS)

	// 获取当前逻辑CPU数量
	fmt.Println("逻辑CPU数量:", runtime.NumCPU())

	// 设置最大的可同时使用的CPU的核数，取逻辑CPU数量
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)

	// goexit 终止当前goroutine
	// 创建一个goroutine
	go func() {
		fmt.Println("start...")
		runtime.Goexit()
		fmt.Println("end...")
	}()

	time.Sleep(3 * time.Second) // 主goroutine 休眠3秒让goroutine执行完成
	fmt.Println("main_end...")

}
