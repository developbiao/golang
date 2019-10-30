package main

import "fmt"

const (
    a = iota
    b = iota
    c = iota
)

func main() {
    fmt.Println("ok")
    fmt.Print("a:%v\n", a)
    fmt.Print("b:%v\n", b)
    fmt.Print("c:%v\n", c)
}
