package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, M int
	fmt.Fscan(reader, &N, &M)

	var array = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &array[i])
	}
	var max int
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				sum := array[i] + array[j] + array[k]
				if sum > max && sum <= M {
					max = sum
				}
			}
		}
	}
	fmt.Fprint(writer, max)
	writer.Flush()
}

/*
	방법
	1. 반복문
	2. 재귀
	3. Next Permutation
*/
