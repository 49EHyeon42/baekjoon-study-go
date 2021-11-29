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

	var testCase string
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &testCase)
		var testScore, point int = 0, 1
		for _, result := range testCase {
			if result == 'O' {
				testScore, point = testScore+point, point+1
			} else if result == 'X' {
				point = 1
			}
		}
		fmt.Fprintln(writer, testScore)
	}

	writer.Flush()
}
