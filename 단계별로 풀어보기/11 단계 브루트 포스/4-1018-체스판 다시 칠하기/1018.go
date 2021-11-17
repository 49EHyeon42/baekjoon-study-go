package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, M int
	fmt.Fscanln(reader, &N, &M)

	chessboard := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &chessboard[i])
	}
	min := 9223372036854775807
	for i := 0; i < N-7; i++ {
		for j := 0; j < M-7; j++ {
			countChangeB, countChangeW := 0, 0
			for k := i; k < i+8; k++ {
				for l := j; l < j+8; l++ {
					if (k+l)%2 == 0 {
						if chessboard[k][l] == 'B' {
							countChangeB += 1
						} else {
							countChangeW += 1
						}
					} else {
						if chessboard[k][l] == 'B' {
							countChangeW += 1
						} else {
							countChangeB += 1
						}
					}
				}
			}
			if min > getMin(countChangeB, countChangeW) {
				min = getMin(countChangeB, countChangeW)
			}
		}
	}
	fmt.Fprint(writer, min)
	writer.Flush()
}

func getMin(first, second int) int {
	if first >= second {
		return second
	}
	return first
}
