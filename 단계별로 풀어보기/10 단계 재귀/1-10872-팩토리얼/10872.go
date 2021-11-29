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
	fmt.Fprint(writer, factorial(N))
	writer.Flush()
}

func factorial(n int) int {
	if n > 1 {
		return n * factorial(n-1)
	} else {
		return 1
	}
}
