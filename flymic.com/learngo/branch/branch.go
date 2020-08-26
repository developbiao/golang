package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println( grade(50) )
	fmt.Println( grade(90) )
	fmt.Println( grade(60) )

	readFile()

}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "f"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func readFile() {
	const filename = "branch/hello.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}