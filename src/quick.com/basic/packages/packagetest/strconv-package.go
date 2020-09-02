package packagetest

import (
	"fmt"
	"strconv"
)

func TestStrConv() {
	// string 转 boolean
	s1 := "true"
	b, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T, %t", b, b)
	// string 转int
	var s2 string = "100"
	b2, err2 := strconv.ParseInt(s2, 16, 64)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("%T, %d", b2, b2)
}
