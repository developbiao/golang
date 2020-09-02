package main

import "fmt"

// 将漂亮女孩定义为int 类型
type BeautyGirl int

// 使用Clear方法将BeautyGirl值清空
func (b BeautyGirl) Clear() bool {
	b = 0
	return b == 0
}

// 使用Add方法给BeautyGirl增加值
func (b BeautyGirl) Add(num int) int {
	return int(b) + num
}

func main() {
	var lzl BeautyGirl
	fmt.Println(lzl.Clear())
	fmt.Println(lzl.Add(3))
	fmt.Println(lzl.Clear())
	fmt.Println(lzl.Add(6))
	fmt.Println(lzl.Add(8))
	fmt.Println(lzl.Clear())

}
