package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	sudoku [9][9]int
)

func main() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fscan(reader, &sudoku[i][j])
		}
	}
	dfs(0, 0)
	writer.Flush()
}

func dfs(row, col int) {
	if col == 9 {
		dfs(row+1, 0)
		return
	}
	if row == 9 {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Fprintf(writer, "%d ", sudoku[i][j])
			}
			fmt.Fprintln(writer)
		}
		return
	}
	if sudoku[row][col] == 0 {
		for i := 1; i <= 9; i++ {
			if isPossible(row, col, i) {
				sudoku[row][col] = i
				dfs(row, col+1)
			}
		}
		sudoku[row][col] = 0
		return
	}
	dfs(row, col+1)
}

func isPossible(row, col, value int) bool {
	// row line check
	for i := 0; i < 9; i++ {
		if sudoku[row][i] == value {
			return false
		}
	}
	// col line check
	for i := 0; i < 9; i++ {
		if sudoku[i][col] == value {
			return false
		}
	}
	// square check
	squareRow, squareCol := row/3*3, col/3*3
	for i := squareRow; i < squareRow+3; i++ {
		for j := squareCol; j < squareCol+3; j++ {
			if sudoku[i][j] == value {
				return false
			}
		}
	}
	return true
}

// ! error 전체가 0 이라면?
// https://cryptosalamander.tistory.com/59
