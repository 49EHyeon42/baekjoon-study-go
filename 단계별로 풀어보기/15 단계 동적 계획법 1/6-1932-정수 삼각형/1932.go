package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)
	if n == 1 {
		fmt.Fscan(reader, &n)
		fmt.Fprint(writer, n)
	} else {
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j <= i; j++ {
				fmt.Fscan(reader, &dp[i][j])
			}
		}

		answer := 0
		for i := 1; i < n; i++ {
			for j := 0; j <= i; j++ {
				if j == 0 {
					dp[i][j] += dp[i-1][j]
				} else if j == i {
					dp[i][j] += dp[i-1][j-1]
				} else {
					dp[i][j] += max(dp[i-1][j], dp[i-1][j-1])
				}
				if dp[i][j] > answer {
					answer = dp[i][j]
				}
			}
		}
		fmt.Fprint(writer, answer)
	}
	writer.Flush()
}

func max(first, second int) int {
	if first >= second {
		return first
	}
	return second
}

// 메모리 시간 조금 큼, 수정 필요
