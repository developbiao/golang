package main

import "fmt"
import "math"

func main() {
    t, ok := mySqrt(0)
    fmt.Println(t)
    fmt.Println(ok)

}

func mySqrt(f float64) (v float64, ok bool) {
    if f < 0 { return } // error case
    return math.Sqrt(f), true
}
