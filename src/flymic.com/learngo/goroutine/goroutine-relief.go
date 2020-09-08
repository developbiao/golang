package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义全局变量 表示救济粮食总是
var food = 10

// 同步等待组对象
var wg sync.WaitGroup

// 创建一把锁
var mutex sync.Mutex

func main() {
	wg.Add(4)
	// 开启4个协程报粮食
	go relief("灾民好家伙1")
	go relief("灾民好家伙2")
	go relief("灾民老李头1")
	go relief("灾民老李头2")

	wg.Wait() // 阻塞主协程，让子协程完成任务
	fmt.Println("ok")

}

// 定义一个发放的救灾发放的方法
func relief(name string) {
	defer wg.Done()
	for {
		mutex.Lock() // 上锁
		if food > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			food--
			fmt.Println(name, "抢到了救济粮，还剩下", food, "个")
			mutex.Unlock() // 解锁
		} else {
			fmt.Println(name, "抢锤子抢，没有粮食了!")
			mutex.Unlock() // 条件不满足取需要解锁，不然其它协程进不来
			break
		}
	}
}
