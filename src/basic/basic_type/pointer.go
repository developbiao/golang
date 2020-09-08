package main

import "fmt"

func main() {
    var i1 = 5
    fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)

    var int_p *int
    int_p = &i1
    fmt.Printf("The value at memory location %p is %d\n", int_p, *int_p)
}
