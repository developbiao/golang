package main

import "fmt"

type Role struct {
	Name    string
	Ability string
	Level   int
	Kill    float64
}

// 创建一个方法只要是Role结构体就能调用
func (r Role) Kungfu() {
	fmt.Printf("我是:%s, 我会的武功: %s, 已经练到了%d级了， 杀伤力:%.1f\n",
		r.Name, r.Ability, r.Level, r.Kill)
}

// 指针类型方法
func (r *Role) Kungfu2() {
	fmt.Printf("Kungfu2-我是:%s, 我会的武功: %s, 已经练到了%d级了， 杀伤力:%.1f\n",
		r.Name, r.Ability, r.Level, r.Kill)

}

func main() {

	// 使用Role结构体创建一个角色代表"派达星"
	pai := Role{"派达星", "吸星大法", 10, 9.9}
	pai.Kungfu()

	// 使用Role结构体创建角色"张无忌"
	zwj := &Role{"张无忌", "九阳神功", 9, 13}
	zwj.Kungfu2()

}
