package main

import "fmt"

func main () {
    var a int = 30
    var ip *int

    ip = &a

    fmt.Printf("a value address:%x\n", ip)
    fmt.Printf("*ip value address:%d\n", *ip)
}
