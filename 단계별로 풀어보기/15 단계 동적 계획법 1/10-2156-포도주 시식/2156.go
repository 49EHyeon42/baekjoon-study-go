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
	fmt.Fscanln(reader, &N)

	dp := make([][]int, 101)
	for i := range dp {
		dp[i] = make([]int, 11)
	}

	for i := 1; i <= 9; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= N; i++ {
		dp[i][0] = dp[i-1][1]
		for j := 1; j <= 9; j++ {
			dp[i][j] = (dp[i-1][j-1] + dp[i-1][j+1]) % 1000000000
		}
	}

	answer := 0
	for i := 0; i < 10; i++ {
		answer += dp[N][i]
	}
	fmt.Fprint(writer, answer%1000000000)
	writer.Flush()
}

// ! 어려움 다시보기
