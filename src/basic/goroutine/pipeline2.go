package main

import (
	"fmt"
	"strings"
)

// 流水线地鼠
func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	// 装资源地鼠
	go sourceGopher(c0)
	// 过滤地鼠
	go filterGopher(c0, c1)

	// 包装地鼠
	printGopher(c1)

}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for {
		item := <-upstream
		if item == "" {
			close(downstream)
			return
		}

		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}

}

func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}
