package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	N, _ := reader.ReadBytes('\n')

	for i := 0; i < len(N)-1; i++ {
		for j := i; j < len(N); j++ {
			if N[i] < N[j] {
				N[i], N[j] = N[j], N[i]
			}
		}
	}

	writer.WriteString(string(N))

	writer.Flush()
}
