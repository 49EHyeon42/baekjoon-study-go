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

	var str []byte
	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &str)
		var stack []byte
		flag := true
		for _, b := range str {
			if b == 40 {
				stack = append(stack, b)
			} else {
				if len(stack) != 0 {
					stack = stack[:len(stack)-1]
				} else {
					flag = false
					break
				}
			}
		}
		if len(stack) == 0 && flag {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
	writer.Flush()
}
