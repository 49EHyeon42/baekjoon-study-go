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

	var count, space int = 1, 1
	for N > space {
		count, space = count+1, space+(6*count)
	}

	fmt.Fprint(writer, count)

	writer.Flush()
}
