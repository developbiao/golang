package main

import "fmt"

// 创建一个结构体起名 Ouyangcrazy 代表父类
type Ouyangcrazy struct {
	Name    string
	Age     int
	Ability string
}

// 创建一个结构体代表子类
type YangGuo struct {
	Ouyangcrazy        // 包含父类的所有属性
	Address     string // 单独子类有的字段
}

// 父类的方法
func (o *Ouyangcrazy) ToadKongfu() {
	fmt.Println(o.Name, "的蛤蟆功！")
}

// 子类的方兴
func (y *YangGuo) NewKongfu() {
	fmt.Println(y.Name, "子类自己的功夫!")
}

// 子类重写父类的方法
func (y *YangGuo) ToadKongfu() {
	fmt.Println(y.Name, "改造后的蛤蟆功！")
}

func main() {
	o := &Ouyangcrazy{Name: "欧阳疯", Age: 70} // 创建父类
	fmt.Println(o.Name)
	o.ToadKongfu()

	// 创建子类
	y := &YangGuo{Ouyangcrazy{Name: "杨过", Age: 18}, "古墓"}
	fmt.Println(y.Name)
	fmt.Println(y.Address)

	y.ToadKongfu() // 子类访问父类的方法
	y.NewKongfu()  // 子类访问自己的方法
	//y.ToadKongfu() // 如果存在自己的方法， 访问自己重写的方法

}
