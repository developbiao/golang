package main

import (
    "fmt"
)

func main() {
    _, num, str := numbers() // 获取函数返回最后两个
    fmt.Println(num, str)
}

func numbers()(int, int, string) {
    a, b, c := 3, 8, "gogo"
    return a, b, c

}
