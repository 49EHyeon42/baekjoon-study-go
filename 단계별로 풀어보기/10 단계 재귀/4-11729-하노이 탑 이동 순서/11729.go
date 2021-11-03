package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	var N int
	fmt.Fscan(reader, &N)
	fmt.Fprintln(writer, int(math.Pow(2, float64(N)))-1)
	hanoi(N, 1, 2, 3)
	writer.Flush()
}

func hanoi(N, from, by, to int) {
	if N == 1 {
		fmt.Fprintln(writer, from, to)
	} else {
		hanoi(N-1, from, to, by)
		fmt.Fprintln(writer, from, to)
		hanoi(N-1, by, from, to)
	}
}
