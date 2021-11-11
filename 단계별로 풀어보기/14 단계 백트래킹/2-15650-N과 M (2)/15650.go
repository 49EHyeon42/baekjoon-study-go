package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	var N, M int
	fmt.Fscanln(reader, &N, &M)

	array, comb := make([]int, N), make([]int, M)
	for i := 1; i <= N; i++ {
		array[i-1] = i
	}

	combination(array, comb, M, 0, 0)

	writer.Flush()
}

func combination(array, comb []int, M, index, depth int) {
	if M == 0 {
		for i := 0; i < len(comb); i++ {
			fmt.Fprintf(writer, "%d ", comb[i])
		}
		fmt.Fprintf(writer, "\n")
	} else if depth == len(array) {
	} else {
		comb[index] = array[depth]
		combination(array, comb, M-1, index+1, depth+1)
		combination(array, comb, M, index, depth+1)
	}
}

/*
* 백트래킹 문제
* 조합 알고리즘으로 판단
 */
