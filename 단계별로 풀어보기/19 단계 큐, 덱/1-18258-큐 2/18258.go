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

	var queue []int
	command, value := "", 0
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &command)
		switch command {
		case "push":
			fmt.Fscanln(reader, &value)
			queue = append(queue, value)
		case "pop":
			if len(queue) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				value = queue[0]
				queue = queue[1:]
				fmt.Fprintln(writer, value)
			}
		case "size":
			fmt.Fprintln(writer, len(queue))
		case "empty":
			if len(queue) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case "front":
			if len(queue) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, queue[0])
			}
		case "back":
			if len(queue) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, queue[len(queue)-1])
			}
		}
	}
	writer.Flush()
}

// 메모리, 시간 너무 큼
