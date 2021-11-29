package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, K int
	fmt.Fscan(reader, &N, &K)

	array := make([]int, N)
	for i := N; i > 0; i-- {
		fmt.Fscan(reader, &array[i-1])
	}

	var count int
	for i := 0; i < N; i++ {
		count += K / array[i]
		K %= array[i]
	}

	fmt.Fprint(writer, count)

	writer.Flush()
}
