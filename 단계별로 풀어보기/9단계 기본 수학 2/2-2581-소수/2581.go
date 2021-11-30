package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var M, N int
	fmt.Fscan(reader, &M, &N)
	sum, min := solution(M, N)
	if sum != 0 && min != 10001 {
		fmt.Fprintf(writer, "%d\n%d", sum, min)
	} else {
		fmt.Fprint(writer, -1)
	}
	writer.Flush()
}

func solution(M, N int) (int, int) {
	var sum, min int = 0, 10001
	for i := M; i <= N; i++ {
		if isPrimeNumber(i) {
			if i < min {
				min = i
			}
			sum += i
		}
	}
	return sum, min
}

func isPrimeNumber(number int) bool {
	if number != 1 {
		var i int = 2
		for i*i <= number {
			if number%i == 0 {
				return false
			}
			i += 1
		}
		return true
	} else {
		return false
	}
}
