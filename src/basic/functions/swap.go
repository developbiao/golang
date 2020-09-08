package main

import "fmt"

func main() {
   var a int = 100
   var b int = 300

   fmt.Printf("swap before a=%d\n", a)
   fmt.Printf("swap before b=%d\n", b)

   swap(&a, &b)
   fmt.Printf("swap after a=%d\n", a)
   fmt.Printf("swap after b=%d\n", b)
}

func swap(x *int, y *int) {
     *x, *y = *y, *x
}
