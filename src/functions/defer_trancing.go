package main

import "fmt"

func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func main() {
    b()
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}
