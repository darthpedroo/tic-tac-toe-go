package main

import "fmt"



func main() {

	board := make([][]int, 3)

	for i := range board {
		board[i] = make([]int, 3)
	}
	board[0][0] = 1
	fmt.Print(board)

}
