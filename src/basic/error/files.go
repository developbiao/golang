package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Erros are values.
// Dont't just check errors, handle them gracefully.
// Don't panic.
// Make the zero value useful.
// The bigger the interface, the weaker the abstraction.
// interface{} says nothing.
// Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
// Documentation is for users.
// A little coying is better than a little dependency.
// Clear is better than clever.
// Concurrency is not paralleism.
// Don't communicate by sharing memory, share memory by comunicating.
// Channels orchestrate; mutexe serialize.
func main() {
	files, err := ioutil.ReadDir("/unrooted")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
