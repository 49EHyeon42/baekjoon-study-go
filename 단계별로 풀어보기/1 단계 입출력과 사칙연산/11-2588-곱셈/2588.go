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
	fmt.Fscanln(reader, &A)
	fmt.Fscan(reader, &B)

	fmt.Fprintf(writer, "%d\n%d\n%d\n%d",
		A*((B%100)%10), A*((B%100)/10), A*(B/100), A*B)

	writer.Flush()
}
