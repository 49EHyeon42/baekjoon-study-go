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
	defer writer.Flush()

	var N, M, target int
	fmt.Fscanln(reader, &N)
	array := make([]int, N)
	for i := range array {
		fmt.Fscan(reader, &array[i])
	}
	reader.ReadString('\n') // 버퍼 삭제
	sort.Ints(array)
	fmt.Fscanln(reader, &M)
	for i := 0; i < M; i++ {
		fmt.Fscanf(reader, "%d", &target)
		fmt.Fprintln(writer, binarySearch(array, target))
	}
}

func binarySearch(array []int, target int) int {
	start, end := 0, len(array)-1
	mid := (start + end) / 2

	for end-start >= 0 {
		if array[mid] == target {
			return 1
		} else if array[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = (end + start) / 2
	}
	return 0
}
