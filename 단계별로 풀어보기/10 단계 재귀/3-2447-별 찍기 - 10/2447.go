package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	var N int
	fmt.Fscan(reader, &N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			solution(i, j, N)
		}
		fmt.Fprintf(writer, "\n")
	}
	writer.Flush()
}

func solution(i, j, N int) {
	if i/N%3 == 1 && j/N%3 == 1 {
		fmt.Fprint(writer, " ")
	} else if N/3 == 0 {
		fmt.Fprint(writer, "*")
	} else {
		solution(i, j, N/3)
	}
}
