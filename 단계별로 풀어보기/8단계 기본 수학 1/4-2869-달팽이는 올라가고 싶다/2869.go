package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var A, B, V int
	fmt.Fscan(reader, &A, &B, &V)

	var day int = (V - B) / (A - B)
	if (V-B)%(A-B) != 0 {
		day += 1
	}
	fmt.Fprint(writer, day)

	writer.Flush()
}
