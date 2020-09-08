package main

import (
	"fmt"
	"sync"
	"time"
)

// 创建一把读写锁,可以是他的指针类型
var rwmutex sync.RWMutex

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go readTest(1)
	go writeTest(2)
	go readTest(3)
	wg.Wait()
	fmt.Println("Main End")

}

func readTest(i int) {
	defer wg.Done()
	rwmutex.RLock() // 读上锁
	fmt.Println("===== 准备读取数据 =====")
	fmt.Println("正在读取数据....", i)

	rwmutex.RUnlock() // 读解锁
	fmt.Println("===== 读取完成 =====")
}

func writeTest(i int) {
	defer wg.Done()
	rwmutex.Lock()
	fmt.Println("===== 准备写取数据 =====")
	fmt.Println("正在写入数据....", i)
	time.Sleep(3 * time.Second)
	rwmutex.Unlock()
	fmt.Println("===== 写操作结束 =====")

}
