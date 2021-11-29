package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	N, input, max, min := 0, 0, 1, 1000001
	fmt.Fscanln(reader, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &input)
		if max < input {
			max = input
		}
		if min > input {
			min = input
		}
	}
	fmt.Fprint(writer, max*min)
	writer.Flush()
}
