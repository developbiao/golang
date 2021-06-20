package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义全局变量，表示疫苗总量10支
var vaccine = 10

func main() {
	// 开启4个协程抢疫苗
	go Injection("隔壁老王1")
	go Injection("隔壁老王2")
	go Injection("武候区-吴彦祖1")
	go Injection("武候区-吴彦祖2")

	// 让程序休息5秒等待所有的子协程执行完毕
	time.Sleep(5 * time.Second)

}

// 打疫苗方法
func Injection(name string) {
	for {
		if vaccine > 0 {
			// 随机休睡时间
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// 抢到后减少支数
			vaccine--
			fmt.Println(name, "已接种疫苗, 还剩下", vaccine, "支")
		} else {
			fmt.Println(name, "打个锤子，些别抢了，没疫苗了!")
			break
		}

	}
}
