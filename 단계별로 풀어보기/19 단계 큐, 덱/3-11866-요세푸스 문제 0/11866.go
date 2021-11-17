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

	queue := make([]int, N)
	for i := range queue {
		queue[i] = i + 1
	}
	fmt.Fprintf(writer, "<")
	value := 0
	for len(queue) != 0 {
		for i := 0; i < K-1; i++ {
			queue = append(queue, queue[0])
			queue = queue[1:]
		}
		value = queue[0]
		queue = queue[1:]
		if len(queue) > 0 {
			fmt.Fprintf(writer, "%d, ", value)
		}
	}
	fmt.Fprintf(writer, "%d>", value)
	writer.Flush()
}

// 지저분함, 메모리 조금 큼, 시간 12ms
