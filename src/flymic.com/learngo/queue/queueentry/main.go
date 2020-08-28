package main

import (
	"flymic.com/learngo/queue"
	"fmt"
)

func main() {

	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop(), q.IsEmpty())
	fmt.Println(q.Pop(), q.IsEmpty())
	fmt.Println(q.Pop(), q.IsEmpty())

	q.Push(233)
	fmt.Println(q)

}
