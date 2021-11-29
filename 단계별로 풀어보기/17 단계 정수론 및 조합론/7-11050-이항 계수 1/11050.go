package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, K int
	fmt.Fscanln(reader, &N, &K)
	fmt.Fprintf(writer, "%d", factorial(N)/(factorial(K)*factorial(N-K)))
	writer.Flush()
}

func factorial(number int) int {
	if number != 1 {
		answer := 1
		for i := number; i >= 1; i-- {
			answer *= i
		}
		return answer
	} else {
		return 1
	}
}

// 메모리 초과
func factorial2(number int) int {
	if number == 1 {
		return number
	} else {
		return number * factorial(number-1)
	}
}
