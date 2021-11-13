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

	/*
		파스칼 삼각형
		이항계수[N][K]=이항계수[N-1][K-1]+이항계수[N-1][K]
	*/
	array := make([][]int, 1001)
	for i := range array {
		array[i] = make([]int, 1001)
	}
	for i := 1; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if i == j || j == 0 {
				array[i][j] = 1
			} else {
				array[i][j] = (array[i-1][j-1] + array[i-1][j]) % 10007
			}
		}
	}
	fmt.Fprintf(writer, "%d", array[N][K])
	writer.Flush()
}
