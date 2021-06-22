package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 派出所证明
	var psCertificate sync.Mutex
	// 物业证明
	var propertyCertificate sync.Mutex

	// 派出所和物业都要处理，加入等待组
	var wg sync.WaitGroup
	wg.Add(2)

	// 派出所的goroutine
	go func() {
		defer wg.Done() // 派出所处理完成

		psCertificate.Lock()
		defer psCertificate.Unlock()

		// 检查材料
		time.Sleep(3 * time.Second)
		// 请求物业的证明
		propertyCertificate.Lock()
		propertyCertificate.Unlock()

	}()

	// 物业处理goroutine
	go func() {
		defer wg.Done() // 物业处理完成

		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()

		// 检查材料
		time.Sleep(5 * time.Second)
		// 请求派出所的证明
		psCertificate.Lock()
		psCertificate.Unlock()
	}()

	wg.Wait()
	fmt.Println("执行完成!")

}
