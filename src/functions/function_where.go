package main

import "fmt"
import "runtime"
import "log"

func main() {
    where := func() {
        _, file, line, _ := runtime.Caller(1)
        log.Printf("%s:%d", file, line)
    }

    where()
    fmt.Printf("some code 1")
    where()
    fmt.Printf("some code 2")
    where()
}



