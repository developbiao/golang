package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// (*.p).X, that notation is cumbersome, so the language permits us instead to write just p.X
	p.X = 2e3
	fmt.Println(v)

	v2 := Vertex{X: 3}
	fmt.Println(v2)

}
