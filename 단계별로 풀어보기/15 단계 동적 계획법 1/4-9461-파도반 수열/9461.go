package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var T int
	fmt.Fscanln(reader, &T)

	var N int
	dp := make([]int, 101)
	dp[1], dp[2], dp[3], dp[4], dp[5] = 1, 1, 1, 2, 2
	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &N)
		for j := 6; j <= N; j++ {
			dp[j] = dp[j-1] + dp[j-5]
		}
		fmt.Fprintln(writer, dp[N])
	}
	writer.Flush()
}
