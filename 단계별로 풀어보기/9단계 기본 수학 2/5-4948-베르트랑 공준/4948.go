package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	max := 246912
	isNotPrimeArr, primeNumArr := make([]bool, max+1), make([]int, max+1)
	for i := 2; i <= max; i++ {
		if isNotPrimeArr[i] {
			continue
		}
		for j := i * i; j <= max; j += i {
			isNotPrimeArr[j] = true
		}
	}
	for i := 2; i <= max; i++ {
		primeNumArr[i] = primeNumArr[i-1]
		if !isNotPrimeArr[i] {
			primeNumArr[i] += 1
		}
	}

	var n int
	for {
		fmt.Fscan(reader, &n)
		if n == 0 {
			break
		}
		fmt.Fprintln(writer, primeNumArr[n*2]-primeNumArr[n])
	}
	writer.Flush()
}
