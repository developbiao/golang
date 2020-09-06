package main

import "fmt"

// 定义接口
type Kongfu interface {
	Toad()      // 蛤蟆功
	SixSwords() // 六脉神剑
}

// 实现类型
type Haojiahuo struct {
	name string
}

// 实现类型
type Laolitou struct {
	name string
}

// 老李头实现方法
func (f Laolitou) Toad() {
	fmt.Println(f.name, "实现了蛤蟆功...")
}

// 老李头实现了六脉神剑
func (f Laolitou) SixSwords() {
	fmt.Println(f.name, "实现了六脉神剑...")
}

// 老李头实现自己的方法
func (f Laolitou) PlayGame() {
	fmt.Println(f.name, "玩LOL游戏...")
}

func (h Haojiahuo) Toad() {
	fmt.Println(h.name, "也实现了蛤蟆功~~~")
}

func (h Haojiahuo) SixSwords() {
	fmt.Println(h.name, "也实现了六脉神剑~~~")
}

// 测试方法
func testInterface(k Kongfu) {
	k.Toad()
	k.SixSwords()
}

func main() {
	h := Haojiahuo{name: "好家伙"}
	fmt.Println(h.name)

	l := Laolitou{name: "老李头"}
	fmt.Println(l.name)

	// testInterface 需要参数类型为Kongfu接口类型的参数
	// h实现了Kongfu接口方法
	// h就是这个接口的实现 就可以作为这个函数的参数
	testInterface(h)

	var kf Kongfu
	kf = h
	kf.Toad()

	l.PlayGame()

}
