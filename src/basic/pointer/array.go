package main

import "fmt"

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}

func main() {
	var board [8][8]rune
	fmt.Println("Before board")
	fmt.Println(board)

	reset(&board)

	fmt.Println("After board")
	fmt.Println(board)

	fmt.Printf("%c\n", board[0][0])
}
