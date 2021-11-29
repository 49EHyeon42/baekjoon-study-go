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

	command, _ := reader.ReadBytes('\n')

	for i := 1; i < N; i++ {
		temp, _ := reader.ReadBytes('\n')
		for j := 0; j < len(command); j++ {
			if command[j] != temp[j] {
				command[j] = '?'
			}
		}
	}

	fmt.Fprintln(writer, string(command))

	writer.Flush()
}
