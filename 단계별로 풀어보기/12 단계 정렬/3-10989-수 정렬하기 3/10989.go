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

	array, number := make([]int, 10001), 0
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &number)
		array[number] += 1
	}

	for i := 1; i <= 10000; i++ {
		for j := 0; j < array[i]; j++ {
			fmt.Fprintln(writer, i)
		}
	}

	writer.Flush()
}
