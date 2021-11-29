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
		if array[i][1] == array[j][1] {
			return array[i][0] < array[j][0]
		} else {
			return array[i][1] < array[j][1]
		}
	})

	var count, temp int
	for i := 0; i < N; i++ {
		if temp <= array[i][0] {
			temp, count = array[i][1], count+1
		}
	}

	fmt.Fprint(writer, count)

	writer.Flush()
}
