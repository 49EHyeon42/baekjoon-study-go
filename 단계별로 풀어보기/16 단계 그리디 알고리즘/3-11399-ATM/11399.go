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

	mins := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &mins[i])
	}

	sort.Ints(mins)

	sum := 0
	for i, min := range mins {
		sum += (N - i) * min
	}

	fmt.Fprint(writer, sum)

	writer.Flush()
}
