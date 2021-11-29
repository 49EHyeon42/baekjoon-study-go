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
	fmt.Fscanln(reader, &A, &B, &C)

	fmt.Fprintf(writer, "%d\n%d\n%d\n%d", (A+B)%C, ((A%C)+(B%C))%C, (A*B)%C, ((A%C)*(B%C))%C)

	writer.Flush()
}
