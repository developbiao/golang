package main

import "fmt"

type ByteSize float64
const (
    _ = iota 
    KB ByteSize = 1<<(10*iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func main() {
    fmt.Printf("KB:%d\n", KB)
    fmt.Printf("MB:%d\n", MB)
    fmt.Printf("GB:%d\n", GB)
    fmt.Printf("TB:%d\n", TB)
    fmt.Printf("PB:%d\n", PB)
    fmt.Printf("EB:%d\n", EB)
    fmt.Printf("ZB:%d\n", ZB)
    fmt.Printf("YB:%d\n", YB)
}


