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
	for i := M; i <= N; i++ {
		if isPrimeNumber(i) {
			fmt.Fprintln(writer, i)
		}
	}
	writer.Flush()
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

// * 나중에  에라스토테네스의 체 알고리즘으로 구현
