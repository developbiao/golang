package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string = "Hi, I'm Biao, Hi."

    fmt.Printf("The posiion of \"Biao\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "Marc"))

    fmt.Printf("The positions of the first indestance of \"Hi\" is:")
    fmt.Printf("%d\n", strings.Index(str, "Hi"))

    fmt.Printf("Teh position of the last instance of \"Hi\" is:")
    fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))


}
