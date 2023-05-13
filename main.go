package main

import "fmt"

func main() {
	board := make([][]int, 10)
	for i := range board {
		board[i] = make([]int, 10)
	}
	fmt.Println(board)
}
