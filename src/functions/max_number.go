package main

import "fmt"

func main() {
   var a int = 100
   var b int = 200
   var ret int = max(a, b)
   fmt.Println("max value is :", ret)
}

func max(num1, num2 int) int {
    var result int
    if num1 > num2 {
        result = num1
    } else {
        result = num2
    }

   return result
}

