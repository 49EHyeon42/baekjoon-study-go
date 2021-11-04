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
	people := make([][]int, N)
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscanln(reader, &x, &y)
		people[i] = []int{x, y}
	}

	for i := range people {
		answer := 0
		for j := range people {
			if people[i][0] < people[j][0] && people[i][1] < people[j][1] {
				answer += 1
			}
		}
		fmt.Fprintf(writer, "%d ", answer+1)
	}
	writer.Flush()
}
