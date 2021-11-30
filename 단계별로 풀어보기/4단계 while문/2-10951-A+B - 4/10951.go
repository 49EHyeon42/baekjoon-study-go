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
		_, err := fmt.Fscan(reader, &A, &B)

		if err == nil {
			fmt.Fprintln(writer, A+B)
		} else {
			break
		}
	}

	writer.Flush()
}
