package main

import "fmt"
import "strconv"
import "os"

func main() {
    var orig string = "ACD"
    // var an int
    var new_str string

    fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

    // anInt, err = strconv.Atoi(origStr)
    an, err := strconv.Atoi(orig)
    if err != nil {
        fmt.Printf("Orig %s is not an integer existing with error\n", orig)
        os.Exit(1)
    }
    fmt.Printf("The integer is %d\n", an)
    an += 5
    new_str = strconv.Itoa(an)
    fmt.Printf("The new string is: %s\n", new_str)
}
