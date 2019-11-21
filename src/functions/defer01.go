package main
import "fmt"

func main() {
    foo()
}

func foo() {
    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }
}
