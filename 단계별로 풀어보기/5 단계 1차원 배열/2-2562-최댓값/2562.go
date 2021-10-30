package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var data, max, idx int
	for i := 0; i < 9; i++ {
		fmt.Fscan(reader, &data)
		if data > max {
			max = data
			idx = i
		}
	}
	fmt.Fprintf(writer, "%d\n%d", max, idx+1)

	writer.Flush()
}
