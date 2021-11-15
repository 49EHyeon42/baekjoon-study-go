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

	if N < 4 {
		fmt.Fprint(writer, N)
	} else {
		temp1, temp2 := 1, 2
		for i := 0; i < N-2; i++ {
			temp1, temp2 = temp2%15746, (temp1+temp2)%15746
		}
		fmt.Fprint(writer, temp2)
	}
	writer.Flush()
}

/*
	! 런타임 에러
	dp := make([]int, N+1)
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= N; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 15746
	}
	fmt.Fprint(writer, dp[N])

	* 방법 1
	if N == 1 {
		fmt.Fprint(writer, 1)
	} else if N == 2 {
		fmt.Fprint(writer, 2)
	} else {
		dp := make([]int, 3)
		dp[0], dp[1], dp[2] = 0, 1, 2
		for i := 0; i <= N-2; i++ {
			dp[0] = dp[1]
			dp[1] = dp[2]
			dp[2] = (dp[0] + dp[1]) % 15746
		}
		fmt.Fprint(writer, dp[1])
	}
*/
