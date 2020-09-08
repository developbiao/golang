package main

import "fmt"

func main() {
    str := "Chinese: 中文传呼机"

    fmt.Println("index int(rune) rune   char bytes")
    for index, rune := range str {
        fmt.Printf("%-2d    %d  %U '%c' % X\n", index, rune, rune, []byte(string(rune)))
    } 
}
