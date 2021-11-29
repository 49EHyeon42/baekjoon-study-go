package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(reader, &n)
	fmt.Fprint(writer, fibonacci(n))
	writer.Flush()
}

func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
