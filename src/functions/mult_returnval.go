package main

import "fmt"

var sum, mul, dif int

func main() {
    sum, mul, dif = computeNum(3, 7)
    PrintValues()

    sum, mul, dif = computeNum2(3, 7)
    PrintValues()

}

func computeNum(num1 int, num2 int) (sum int, mul int, dif int) {
    sum = num1 + num2
    mul = num1 * num2
    dif = num1 - num2
    return
}

func computeNum2(num1 int, num2 int) (int, int, int) {
    sum = num1 + num2
    mul = num1 * num2
    dif = num1 - num2
    return sum, mul, dif
}

func PrintValues() {
    fmt.Printf("sum = %d, mul = %d, dif = %d\n", sum, mul, dif)
}
