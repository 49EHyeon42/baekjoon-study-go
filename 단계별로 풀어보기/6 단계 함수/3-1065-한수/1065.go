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
	if N > 100 {
		var count int = 99
		for i := 100; i <= N; i++ {
			hun, ten, one := i/100, (i/10)%10, i%10
			if hun-ten == ten-one {
				count += 1
			}
		}
		return count
	} else {
		return N
	}
}
