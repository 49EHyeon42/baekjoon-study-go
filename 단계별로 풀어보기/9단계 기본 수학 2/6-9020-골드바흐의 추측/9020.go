package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	max := 10000
	isNotPrimeArr := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		if isNotPrimeArr[i] {
			continue
		}
		for j := i * i; j <= max; j += i {
			isNotPrimeArr[j] = true
		}
	}

	var T, n int
	fmt.Fscan(reader, &T)
	for i := 0; i < T; i++ {
		fmt.Fscan(reader, &n)
		first, second := n/2, n/2
		for {
			if isNotPrimeArr[first] == false && isNotPrimeArr[second] == false {
				fmt.Fprintf(writer, "%d %d\n", first, second)
				break
			}
			first, second = first-1, second+1
		}
	}
	writer.Flush()
}
