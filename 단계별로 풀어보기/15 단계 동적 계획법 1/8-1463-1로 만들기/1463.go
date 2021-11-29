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

	dp := make([]int, N+1)
	for i := 2; i <= N; i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = min(dp[i], dp[i/2]+1)
		}
		if i%3 == 0 {
			dp[i] = min(dp[i], dp[i/3]+1)
		}
	}
	fmt.Fprint(writer, dp[N])
	writer.Flush()
}

func min(arr ...int) int {
	min := 9223372036854775807
	for _, value := range arr {
		if min > value {
			min = value
		}
	}
	return min
}

/*
	더 가벼운 방법
	1. 큐

	2.
	package main
import ("fmt")

func main() {
    var n, count int
    fmt.Scan(&n)
    values := make([]map[int]bool, 0)
    values = append(values, make(map[int]bool))
    values[0][n] = true

    for !values[count][1] {
        values = append(values, make(map[int]bool))
        count++

        for value, _ := range values[count-1] {
            values[count][value-1] = true

            if value % 2 == 0 {
                values[count][value/2] = true
            }

            if value % 3 == 0 {
                values[count][value/3] = true
            }
        }
    }

    fmt.Printf("%d\n", count)
}
*/
