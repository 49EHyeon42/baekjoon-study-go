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

	for i := 1; i <= N; i++ {
		for j := N - i; j > 0; j-- {
			fmt.Fprint(writer, " ")
		}
		for j := 0; j < i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "\n")
	}

	writer.Flush()
}
