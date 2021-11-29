package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscan(reader, &N)
	var str string
	fmt.Fscan(reader, &str)

	var sum int
	for _, r := range str {
		sum += int(r - '0')
	}
	fmt.Fprintln(writer, sum)

	writer.Flush()
}
