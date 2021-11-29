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

	dp := make([]int, 41)

	var N int
	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &N)
		if N == 0 {
			fmt.Fprintf(writer, "%d %d\n", 1, 0)
		} else if N == 1 {
			fmt.Fprintf(writer, "%d %d\n", 0, 1)
		} else {
			fibonacci(dp, N)
			fmt.Fprintf(writer, "%d %d\n", dp[N-1], dp[N])
		}
	}
	writer.Flush()
}

func fibonacci(dp []int, number int) int {
	if number <= 0 {
		dp[0] = 0
		return 0
	} else if number == 1 {
		dp[1] = 1
		return 1
	}
	if dp[number] != 0 {
		return dp[number]
	} else {
		dp[number] = fibonacci(dp, number-1) + fibonacci(dp, number-2)
		return dp[number]
	}
}
