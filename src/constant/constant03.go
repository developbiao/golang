package main 

import "fmt"

func main() {
        const (
               a = iota  //0
               b 
               c 
               d = "ha"
               e
               f = 100
               g
               h = iota
               i
        )

        fmt.Println(a, b, c, d, e, f, g, h, i)
}

