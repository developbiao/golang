package main

import (
	"fmt"
	"sync"
)

// 定义全局变量，表示救济粮问量
var food = 10

// 同步等到组对象
var wg sync.WaitGroup

// 创建一把锁
var mutex sync.Mutex

func main() {
	peopleNum := 4
	wg.Add(peopleNum)
	for i := 1; i <= peopleNum; i++ {
		go Relief(fmt.Sprintf("灾民#%d号", i))
	}
	wg.Wait() // 阻塞主协程, 等待子协程执行结束
}

func Relief(name string) {
	defer wg.Done()
	for {
		// 上锁
		mutex.Lock()
		if food > 0 {
			food--
			fmt.Println(name, "抢到了救济粮🍔", "还剩下", food, "个")
		} else {
			// 条件不满足也需要解锁，否则会造成死锁其它不能执行
			mutex.Unlock()
			fmt.Println(name, "别抢了，没有粮食了。🐣")
			break
		}
		// 执行完成解锁, 让其它协程也能够进来
		mutex.Unlock()
	}
}
