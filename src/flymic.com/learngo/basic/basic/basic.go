package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "gogo"
	bb = true
)

func main() {
	fmt.Println("Hello world go again!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enums()
}

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 7, false, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 7, false, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		php
	)

	const (
		_ = iota // ignore first value by assigning to blank identifier
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, python, golang, php)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
