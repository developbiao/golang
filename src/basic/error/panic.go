package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Catched error", e)
		}
	}()

	panic("I forgot my towel")

	fmt.Println("Programing is done!")
}
