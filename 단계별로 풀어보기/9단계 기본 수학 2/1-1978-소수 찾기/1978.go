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
	var count int
	for i := 0; i < N; i++ {
		var number int
		fmt.Fscan(reader, &number)
		if isPrimeNumber(number) {
			count += 1
		}
	}
	fmt.Fprint(writer, count)
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
