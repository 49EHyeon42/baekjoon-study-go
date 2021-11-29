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

	var west, east int
	dp := make([][]int, 30)
	for i := range dp {
		dp[i] = make([]int, 30)
	}
	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &west, &east)
		fmt.Fprintf(writer, "%d\n", combination(dp, east, west))
	}
	writer.Flush()
}

func combination(dp [][]int, n, r int) int {
	if dp[n][r] > 0 {
	} else if n == r || r == 0 {
		dp[n][r] = 1
	} else {
		dp[n][r] = combination(dp, n-1, r-1) + combination(dp, n-1, r)
	}
	return dp[n][r]
}

/*
	조합으로 중복 없이 선택

	조합 공식의 성질
	n+1Cr+1 = nCr + nCr+1
	nC0 = nCn = 1

	방법 2가지 = 재귀, 반복문(bottom-up ?)
*/
