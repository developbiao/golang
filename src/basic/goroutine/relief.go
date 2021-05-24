package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义全局变量，表示救济粮食总量
var food = 10

func main() {
	// 开启4个协程抢粮食
	go Relief("灾民好家伙1")
	go Relief("灾民好家伙2")
	go Relief("灾民老王头1")
	go Relief("灾民老王头2")

	// 让程序休息5秒等待所有的子协程执行完毕
	time.Sleep(5 * time.Second)

}

func Relief(name string) {
	for {
		if food > 0 {
			// 随机休睡时间
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			food--
			fmt.Println(name, "抢到了救济粮, 还剩下", food, "个")
		} else {
			fmt.Println(name, "大哥些别抢了，没粮食了!")
			break
		}

	}
}
