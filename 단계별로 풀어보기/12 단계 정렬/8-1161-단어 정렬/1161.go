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

	array := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &array[i])
	}

	sort.Strings(array)

	sort.Slice(array, func(i, j int) bool {
		if len(array[i]) == len(array[j]) {
			return array[i] < array[j]
		} else {
			return len(array[i]) < len(array[j])
		}
	})

	var overlap string
	for i := range array {
		if overlap != array[i] {
			fmt.Fprintln(writer, array[i])
			overlap = array[i]
		}
	}

	writer.Flush()
}
