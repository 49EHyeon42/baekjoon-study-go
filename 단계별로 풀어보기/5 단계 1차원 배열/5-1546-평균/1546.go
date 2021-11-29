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
	var input, M float32
	fmt.Fscan(reader, &N)
	array := make([]float32, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &input)
		array[i] = input
		if input > M {
			M = input
		}
	}

	var total float32
	for i := 0; i < N; i++ {
		total += array[i] / M * 100
	}

	fmt.Fprint(writer, total/float32(N))

	writer.Flush()
}
