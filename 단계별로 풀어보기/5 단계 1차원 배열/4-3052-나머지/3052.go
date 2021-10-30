package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	array, input := make([]int, 10), 0
	for i := 0; i < 10; i++ {
		fmt.Fscan(reader, &input)
		array[i] = input % 42
	}

	output := 0
	for i := 0; i < 10; i++ {
		count := 0
		for j := 0; j < i; j++ {
			if array[i] == array[j] {
				count += 1
			}
		}
		if count == 0 {
			output += 1
		}
	}
	fmt.Fprint(writer, output)

	writer.Flush()
}
