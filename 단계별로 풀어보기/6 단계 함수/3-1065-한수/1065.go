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

	fmt.Fprint(writer, solution(N))

	writer.Flush()
}

func solution(N int) int {
	var count int
	for i := 1; i <= N; i++ {
		if i < 100 {
			count += 1
		} else {
			hun, ten, one := i/100, (i/10)%10, i%10
			if hun-ten == ten-one {
				count += 1
			}
		}
	}
	return count
}
