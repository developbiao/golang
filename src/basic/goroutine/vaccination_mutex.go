package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义全局变量，表示疫苗总量10支
var vaccine = 10

// 同步等待组对象
var wg sync.WaitGroup

// 创建一把锁
var mutex sync.Mutex

func main() {

	// 定义一组要打疫苗的人
	names := [4]string{
		"隔壁老王1",
		"隔壁老王2",
		"武候区-吴彦祖1",
		"武候区-吴彦祖2",
	}

	// 加入同步等待的人数
	peopleNum := len(names)
	wg.Add(peopleNum)

	// 开启4个协程抢疫苗
	for _, name := range names {
		go Injection(name)
	}

	// 阻塞主协程, 等待子协程执行结束
	wg.Wait()

}

// 打疫苗方法
func Injection(name string) {
	defer wg.Done()
	for {
		// 上锁 Lock
		mutex.Lock()
		if vaccine > 0 {
			// 随机休睡时间
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

			// 抢到后减少支数
			vaccine--
			fmt.Println(name, "已接种疫苗, 还剩下", vaccine, "支")
		} else {
			// 条件不满足也需要解锁，否则会造成死锁deadlock其它不能执行
			mutex.Unlock()
			fmt.Println(name, "打个锤子，些别抢了，没疫苗了!")
			break
		}

		// 执行完成解锁，让其它协程也能够进来 Unlock
		mutex.Unlock()

	}
}
