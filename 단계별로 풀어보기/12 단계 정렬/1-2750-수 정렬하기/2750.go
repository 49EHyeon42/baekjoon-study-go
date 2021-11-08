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
	fmt.Fscan(reader, &N)

	var array = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &array[i])
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N-(i+1); j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	for i := 0; i < N; i++ {
		fmt.Fprintln(writer, array[i])
	}

	writer.Flush()
}

// 버블정렬
// 다른 정렬도 작성
