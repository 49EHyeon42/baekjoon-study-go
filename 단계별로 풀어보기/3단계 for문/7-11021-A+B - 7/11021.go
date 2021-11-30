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

	for i := 1; i <= N; i++ {
		var A, B int
		fmt.Fscanln(reader, &A, &B)
		fmt.Fprintf(writer, "Case #%d: %d\n", i, A+B)
	}

	writer.Flush()
}
