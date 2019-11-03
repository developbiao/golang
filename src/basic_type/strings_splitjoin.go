package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "The quick brown fox jumps over the lazy dog"
    sl := strings.Fields(str)
    fmt.Printf("Splitted in slice: %v\n", sl)
    for _, val := range sl {
        fmt.Printf("%s -", val)
    }
    fmt.Println()

    str2 := "GOlang|THE PHP of world|35"
    sl2 := strings.Split(str2, "|")
    fmt.Printf("Splitted in slice: %v\n", sl2)
    for _, val := range sl2 {
        fmt.Printf("%s - ", val)
    }
    fmt.Println()
    str3 := strings.Join(sl2, "&")
    fmt.Printf("sl2 joined by ;:%s\n", str3)
}
