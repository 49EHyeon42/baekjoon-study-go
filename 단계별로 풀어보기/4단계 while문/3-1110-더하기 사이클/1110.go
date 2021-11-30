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

	var count, newN int = 0, N
	for {
		count += 1
		newN = ((newN % 10) * 10) + ((newN/10)+(newN%10))%10
		if newN == N {
			break
		}
	}
	fmt.Fprint(writer, count)

	writer.Flush()
}
