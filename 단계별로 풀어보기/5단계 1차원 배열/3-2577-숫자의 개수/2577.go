package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var A, B, C int
	fmt.Fscan(reader, &A, &B, &C)
	ABC := A * B * C

	array := make([]int, 10)
	for ABC > 0 {
		array[ABC%10]++
		ABC /= 10
	}
	for _, item := range array {
		fmt.Fprintln(writer, item)
	}

	writer.Flush()
}
