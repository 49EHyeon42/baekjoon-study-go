package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var first, second int
	for {
		fmt.Fscan(reader, &first, &second)
		if first != 0 && second != 0 {
			if second%first == 0 {
				fmt.Fprintln(writer, "factor")
			} else if first%second == 0 {
				fmt.Fprintln(writer, "multiple")
			} else {
				fmt.Fprintln(writer, "neither")
			}
		} else {
			break
		}
	}
	writer.Flush()
}
