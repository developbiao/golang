package main
import "fmt"
import "runtime"

func main() {
    system := runtime.GOOS
    if system == "windows" {
        fmt.Printf("The value is true windows\n")
    } else {
        fmt.Printf("The value is false, %s\n", system)
    }
}
