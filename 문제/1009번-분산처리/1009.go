package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	T, a, b int
)

func main() {
	fmt.Fscanln(reader, &T)
	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &a, &b)
		var answer, b = a, b%4 + 4
		for j := 1; j < b; j++ {
			answer = answer * a % 10
		}
		if answer == 0 {
			answer = 10
		}
		fmt.Fprintln(writer, answer)
	}
	writer.Flush()
}
