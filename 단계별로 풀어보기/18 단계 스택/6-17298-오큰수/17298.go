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

	answer, array, stack := make([]int, N), make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &array[i])
	}
	for i := N - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= array[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			answer[i] = -1
		} else {
			answer[i] = stack[len(stack)-1]
		}
		stack = append(stack, array[i])
	}
	for _, value := range answer {
		fmt.Fprintf(writer, "%d ", value)
	}
	writer.Flush()
}

// * 통과, 메모리, 시간 너무 많이 먹음, 조정 필요
