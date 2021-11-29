package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, X int
	fmt.Fscan(reader, &N, &X)

	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(reader, &a)

		if a < X {
			fmt.Fprint(writer, a, " ")
		}
	}

	writer.Flush()
}
