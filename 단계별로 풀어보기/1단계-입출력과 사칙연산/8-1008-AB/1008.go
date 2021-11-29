package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var A, B float64
	fmt.Fscanln(reader, &A, &B)

	fmt.Fprint(writer, A/B)

	writer.Flush()
}
