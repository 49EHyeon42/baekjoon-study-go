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
		var R int
		fmt.Fscan(reader, &R)

		var str string
		fmt.Fscan(reader, &str)

		for _, r := range str {
			for j := 0; j < R; j++ {
				fmt.Fprint(writer, string(r))
			}
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
}
