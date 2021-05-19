package main

import (
	"fmt"
)

// 定义结构体实现封装
type Person struct {
	Name string
	Age  int
}

func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) GetAge() int {
	return p.Age
}

func main() {
	rose := NewPerson("Rose")
	rose.SetAge(18)
	fmt.Println(rose.Name, rose.GetAge())

}
