package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscanln(reader, &N)

	board := make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fscan(reader, &board[i][j])
		}
	}

	partition(board, 0, 0, N)

	fmt.Fprintf(writer, "%d\n%d", white, blue)

	writer.Flush()
}

var (
	white, blue int
)

func partition(board [][]int, row, col, size int) {
	if isSquare(board, row, col, size) {
		if board[row][col] == 0 {
			white += 1
		} else {
			blue += 1
		}
		return
	}
	newSize := size / 2
	partition(board, row, col, newSize)
	partition(board, row, col+newSize, newSize)
	partition(board, row+newSize, col, newSize)
	partition(board, row+newSize, col+newSize, newSize)
}

func isSquare(board [][]int, row, col, size int) bool {
	color := board[row][col]
	for i := row; i < row+size; i++ {
		for j := col; j < col+size; j++ {
			if board[i][j] != color {
				return false
			}
		}
	}
	return true
}
