package main

import "fmt"

func main() {
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{7, true},
	}

	fmt.Println(s)
}