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

	queue := make([]int, N)
	for i := range queue {
		queue[i] = i + 1
	}
	for len(queue) > 1 {
		queue = queue[1:]
		queue = append(queue, queue[0])
		queue = queue[1:]
	}
	fmt.Fprint(writer, queue[0])
	writer.Flush()
}

// 메모리 소모 매우큼, 시간 소모 큼
