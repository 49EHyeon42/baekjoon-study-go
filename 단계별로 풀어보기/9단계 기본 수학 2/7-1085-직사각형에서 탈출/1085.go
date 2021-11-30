package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var x, y, w, h int
	fmt.Fscanln(reader, &x, &y, &w, &h)

	fmt.Fprintln(writer, min(min(x, w-x), min(y, h-y)))

	writer.Flush()
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}
