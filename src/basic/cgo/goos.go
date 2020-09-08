package main

import (
    "fmt"
    "runtime"
    "os"
)

func main() {
    var goos string = runtime.GOOS
    fmt.Print("The operating system is: %s\n", goos)
    home := os.Getenv("HOME")
    fmt.Printf("Path is %s\n", home)
}
