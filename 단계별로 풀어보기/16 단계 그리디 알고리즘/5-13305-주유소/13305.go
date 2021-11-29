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

	distance, cost := make([]int, N-1), make([]int, N)
	for i := range distance {
		fmt.Fscan(reader, &distance[i])
	}
	for i := range cost {
		fmt.Fscan(reader, &cost[i])
	}

	sum, minCost := 0, cost[0]
	for i := 0; i < N-1; i++ {
		if minCost > cost[i] {
			minCost = cost[i]
		}
		sum += minCost * distance[i]
	}
	fmt.Fprint(writer, sum)
	writer.Flush()
}
