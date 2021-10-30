package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var C int
	fmt.Fscan(reader, &C)

	for i := 0; i < C; i++ {
		var N int
		fmt.Fscan(reader, &N)

		var array []float32 = make([]float32, N)

		var avg, sum float32
		for j := 0; j < N; j++ {
			fmt.Fscan(reader, &array[j])
			sum += array[j]
		}

		avg = sum / float32(N)

		var count int
		for j := 0; j < N; j++ {
			if avg < array[j] {
				count += 1
			}
		}

		fmt.Fprintf(writer, "%.3f%%\n", float32(count)/float32(N)*100)
	}

	writer.Flush()
}
