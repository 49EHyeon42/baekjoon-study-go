package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		str, _ := reader.ReadBytes('\n')
		if str[0] == '.' {
			break
		}
		var stack []byte
		flag := true
		for _, b := range str {
			if b == '(' || b == '[' {
				stack = append(stack, b)
			} else if b == ')' {
				if len(stack) != 0 && '(' == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					flag = false
					break
				}
			} else if b == ']' {
				if len(stack) != 0 && '[' == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					flag = false
					break
				}
			}

		}
		if len(stack) == 0 && flag {
			fmt.Fprintln(writer, "yes")
		} else {
			fmt.Fprintln(writer, "no")
		}
	}
	writer.Flush()
}
