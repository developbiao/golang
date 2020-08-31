package main

import (
	"bufio"
	"flymic.com/learngo/functional/fib"
	"fmt"
	"io"
	"strings"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//fmt.Println(fib.Fibonacci()())
	//fmt.Println(fib.Fibonacci()())
	//fmt.Println(fib.Fibonacci()())

	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
