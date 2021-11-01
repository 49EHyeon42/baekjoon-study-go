package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var A, B, C int
	fmt.Fscan(reader, &A, &B, &C)

	if B < C {
		fmt.Fprint(writer, A/(C-B)+1)
	} else {
		fmt.Fprint(writer, -1)
	}

	writer.Flush()
}
