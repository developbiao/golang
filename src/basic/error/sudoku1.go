package main

import (
	"errors"
	"fmt"
)

const rows, columns = 9, 9

// Grid is a Suodku grid
type Grid [rows][columns]int8

// set digit on a Sudoku grid

func init() {
	fmt.Println("init function is call")
}

func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return errors.New("out of bounds")
	}
	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}

	if column < 0 || column >= columns {
		return false
	}

	return true
}

func main() {
	var g Grid
	err := g.Set(2, 6, 8)
	// trigger error
	err = g.Set(9, 3, 99)
	if err != nil {
		fmt.Printf("An error occurred: %v.\n", err)
	}
	fmt.Println("grid data", g)
}
