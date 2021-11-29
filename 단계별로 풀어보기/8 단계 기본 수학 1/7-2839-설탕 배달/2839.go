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
	fmt.Fscan(reader, &N)

	fmt.Fprint(writer, solution(N))

	writer.Flush()
}

func solution(N int) int {
	var count int
	for N >= 0 {
		if N%5 == 0 {
			count += N / 5
			return count
		}
		count, N = count+1, N-3
	}
	return -1
}
