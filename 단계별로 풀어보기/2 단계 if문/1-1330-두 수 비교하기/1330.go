package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var A, B int
	fmt.Fscanln(reader, &A, &B)

	if A > B {
		fmt.Fprint(writer, ">")
	} else if A < B {
		fmt.Fprint(writer, "<")
	} else if A == B {
		fmt.Fprint(writer, "==")
	}

	writer.Flush()
}
