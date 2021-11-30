package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var X int
	fmt.Fscan(reader, &X)

	var i int
	for i = 0; i < X; i++ {
		X -= i
	}
	if i%2 == 0 {
		fmt.Fprintf(writer, "%d/%d", X, i+1-X)
	} else {
		fmt.Fprintf(writer, "%d/%d", i+1-X, X)
	}

	writer.Flush()
}
