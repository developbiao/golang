package main

import "fmt"

func main() {
    num := 10
    Reply(5, 7, &num)
    fmt.Println("Multiply:", num)

}
func Reply(a int, b int, reply *int) {
    *reply = a * b
}

