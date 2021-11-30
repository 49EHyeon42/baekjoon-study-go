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
	fmt.Fscanln(reader, &N)

	for i := 1; i < 10; i++ {
		fmt.Fprintf(writer, "%d * %d = %d\n", N, i, N*i)
	}

	writer.Flush()
}
