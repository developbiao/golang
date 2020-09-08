package main

var a = "G"

func main() {
    n() // "G"
    m() // "0"
    n() // "0"
}

func n() {
    print(a)
}

func m() {
    a = "0"
    print(a)
}
