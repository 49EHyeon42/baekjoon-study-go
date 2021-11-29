package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var str string
	fmt.Fscan(reader, &str)

	for i := 97; i < 123; i++ {
		var count int
		for _, r := range str {
			if r == rune(i) {
				fmt.Fprintf(writer, "%d ", count)
				break
			}
			count += 1
		}
		if count == len(str) {
			fmt.Fprint(writer, "-1 ")
		}
	}

	writer.Flush()
}
