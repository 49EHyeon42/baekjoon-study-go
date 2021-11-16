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

	var answer []byte
	var stack []int
	input, value, flag := 0, 1, true
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &input)
		for value <= input {
			stack, answer, value = append(stack, value), append(answer, '+'), value+1
		}
		if stack[len(stack)-1] == input {
			stack, answer = stack[:len(stack)-1], append(answer, '-')
		} else {
			flag = false
			break
		}
	}
	if flag {
		for _, b := range answer {
			fmt.Fprintf(writer, "%c\n", b)
		}
	} else {
		fmt.Fprintln(writer, "NO")
	}
	writer.Flush()
}
