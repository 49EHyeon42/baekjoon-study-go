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
	fmt.Fscanln(reader, &N)

	for i := 1; i <= 1000000; i++ {
		sum, temp := i, i
		for temp != 0 {
			sum += temp % 10
			temp /= 10
		}
		if N == sum {
			fmt.Fprint(writer, i)
			break
		} else if N == i {
			fmt.Fprint(writer, 0)
			break
		}
	}
	writer.Flush()
}
