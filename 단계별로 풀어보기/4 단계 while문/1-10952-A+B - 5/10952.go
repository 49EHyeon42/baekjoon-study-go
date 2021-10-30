package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		var A, B int
		fmt.Fscan(reader, &A, &B)

		if A != 0 && B != 0 {
			fmt.Fprintln(writer, A+B)
		} else {
			break
		}
	}

	writer.Flush()
}
