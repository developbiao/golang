package main

import "fmt"

func main() {
    var a int = 4
    var b int32
    var c float32
    var ptr *int

    fmt.Printf("a=%T\n", a)
    fmt.Printf("b=%T\n", b)
    fmt.Printf("c=%T\n", c)

    ptr = &a
    fmt.Printf("a value %d\n", a)
    fmt.Printf("*ptr value value %d\n", *ptr)
}
