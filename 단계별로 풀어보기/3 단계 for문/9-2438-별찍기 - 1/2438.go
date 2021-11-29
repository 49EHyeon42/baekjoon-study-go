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

	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer, "")
	}

	writer.Flush()
}
