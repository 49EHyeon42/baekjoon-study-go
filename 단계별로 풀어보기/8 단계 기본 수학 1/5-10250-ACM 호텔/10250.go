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
	fmt.Fscan(reader, &T)

	for i := 0; i < T; i++ {
		var H, W, N int
		fmt.Fscan(reader, &H, &W, &N)

		var floor int = (N-1)%H + 1
		var number int = (N-1)/H + 1

		fmt.Fprintf(writer, "%d%02d\n", floor, number)
	}

	writer.Flush()
}
