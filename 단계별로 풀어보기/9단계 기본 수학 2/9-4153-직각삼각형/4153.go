package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b, c int
	for {
		fmt.Fscanln(reader, &a, &b, &c)
		if a == 0 && b == 0 && c == 0 {
			break
		}
		a, b, c = a*a, b*b, c*c
		if a+b == c || a+c == b || b+c == a {
			fmt.Fprintln(writer, "right")
		} else {
			fmt.Fprintln(writer, "wrong")
		}
	}

	writer.Flush()
}
