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
	fmt.Fscanln(reader, &N)

	answer, array := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &array[i])
	}
	copy(answer, array)

	sort.Ints(array)

	arrMap, count := make(map[int]int), 0
	for _, key := range array {
		_, exists := arrMap[key]
		if !exists {
			arrMap[key] = count
			count += 1
		}
	}

	for _, key := range answer {
		fmt.Fprintf(writer, "%d ", arrMap[key])
	}
	writer.Flush()
}
