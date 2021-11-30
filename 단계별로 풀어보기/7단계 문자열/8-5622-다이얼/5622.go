package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var str string
	fmt.Fscan(reader, &str)

	var times []int = []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6,
		7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 10, 10, 10, 10}
	var answer int
	for _, r := range str {
		answer += times[r-'A']
	}
	fmt.Fprint(writer, answer)

	writer.Flush()
}
