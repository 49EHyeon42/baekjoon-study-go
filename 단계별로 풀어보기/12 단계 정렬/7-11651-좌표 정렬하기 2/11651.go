package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscan(reader, &N)

	array := make([][]int, N)
	for i := 0; i < N; i++ {
		temp := make([]int, 2)
		fmt.Fscan(reader, &temp[0], &temp[1])
		array[i] = temp
	}

	sort.Slice(array, func(i, j int) bool {
		if array[i][1] < array[j][1] {
			return true
		} else if array[i][1] == array[j][1] && array[i][0] < array[j][0] {
			return true
		}
		return false
	})

	for i := range array {
		fmt.Fprintln(writer, array[i][0], array[i][1])
	}

	writer.Flush()
}
