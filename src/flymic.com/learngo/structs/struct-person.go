package main

import "fmt"

// 定义一个结构体实现封闭
type Person struct {
	Name string
	Age  int
}

// 使用NewPerson方法来创建一个对象
func NewPerson(name string) *Person {
	return &Person{Name: name}
}

// 使用SetAge方法来设置结构体成员的Age
// 方法的接收者谁来调用方法(Person)
func (p *Person) SetAge(age int) {
	p.Age = age
}

// 使用GetAge方法获取成员现在的Age
func (p *Person) GetAge() int {
	return p.Age
}

func main() {
	// 创建一个对象
	lijun := NewPerson("李骏")
	lijun.SetAge(19)
	fmt.Println(lijun.Name, lijun.Age)
}
