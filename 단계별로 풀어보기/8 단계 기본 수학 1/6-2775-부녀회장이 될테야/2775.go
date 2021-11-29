package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var T int
	fmt.Fscan(reader, &T)

	for i := 0; i < T; i++ {
		var k, n int
		fmt.Fscan(reader, &k, &n)

		fmt.Fprintf(writer, "%d\n", solution(k, n))
	}

	writer.Flush()
}

func solution(k, n int) int {
	if n == 1 {
		return 1
	} else if k == 0 {
		return n
	}
	return solution(k-1, n) + solution(k, n-1)
}
