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
	fmt.Fscanln(reader, &N)

	var stack []int
	command, value := "", 0
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &command)
		switch command {
		case "push":
			fmt.Fscanln(reader, &value)
			stack = append(stack, value)
		case "pop":
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				value = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fmt.Fprintln(writer, value)
			}
		case "size":
			fmt.Fprintln(writer, len(stack))
		case "empty":
			if len(stack) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case "top":
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
			}
		}
	}
	writer.Flush()
}
