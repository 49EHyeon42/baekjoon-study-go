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

	fmt.Printf("%d\n%d\n%d\n%d\n%d", A+B, A-B, A*B, A/B, A%B)

	writer.Flush()
}
