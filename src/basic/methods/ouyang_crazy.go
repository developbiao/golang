package main

import "fmt"

type Ouyangcrazy struct {
	Name    string
	Age     int
	Ability string
}

type YangGuo struct {
	Ouyangcrazy        // 包含父类的所有属性
	Address     string // 单独子类有的字段
}

// 父类的方法
func (o *Ouyangcrazy) ToadKongfu() {
	fmt.Println(o.Name, "的哈蟆功")
}

// 子类的方法
func (y *YangGuo) NewKongfu() {
	fmt.Println(y.Name, "子类自己的新功夫！")
}

// 子类重写父类的方法
func (y *YangGuo) ToadKongfu() {
	fmt.Println(y.Name, "的新式飞龙在天蛤蟆功！")
}

func main() {
	o := &Ouyangcrazy{Name: "欧阳疯", Age: 70} // 创建父类
	o.ToadKongfu()

	y := &YangGuo{Ouyangcrazy{Name: "杨过", Age: 18}, "古墓"}
	fmt.Println(y.Name)
	fmt.Println(y.Address)

	//y.ToadKongfu() // 子类对象访问父类方法
	y.ToadKongfu() // 子类访问重写后的方法
	y.NewKongfu()  // 子类访问自己的方法
}
