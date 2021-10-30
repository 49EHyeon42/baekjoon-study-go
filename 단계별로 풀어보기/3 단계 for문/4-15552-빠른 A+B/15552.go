package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var T int
	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		var A, B int
		fmt.Fscanln(reader, &A, &B)
		fmt.Fprintln(writer, A+B)
	}

	writer.Flush()
}
