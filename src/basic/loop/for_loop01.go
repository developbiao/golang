package main

import "fmt"

func main() {
    var b int = 15
    var a int

    numbers := [6]int{1, 2, 3, 4, 5, 6}

    for a := 0; a < 10; a ++ {
        fmt.Printf(" a value is :%d\n", a)
    }

    for a < b {
         a++
         fmt.Printf("a value is :%d\n", a)
    }

    for i, x := range numbers {
        fmt.Printf("the %d bit x value =%d\n", i, x)
    }
}
