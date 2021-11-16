package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var K int
	fmt.Fscanln(reader, &K)

	var stack []int
	value := 0
	for i := 0; i < K; i++ {
		fmt.Fscanln(reader, &value)
		if value != 0 {
			stack = append(stack, value)
		} else {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	answer := 0
	for _, value := range stack {
		answer += value
	}
	fmt.Fprint(writer, answer)
	writer.Flush()
}
