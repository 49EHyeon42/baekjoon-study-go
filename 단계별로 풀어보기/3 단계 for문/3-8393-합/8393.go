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
	fmt.Fscanln(reader, &n)

	var total int
	for i := 1; i <= n; i++ {
		total += i
	}
	fmt.Fprint(writer, total)

	writer.Flush()
}
