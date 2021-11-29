package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var T int
	fmt.Fscanln(reader, &T)

	var n int
	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &n)
		answer, countKind := 1, make(map[string]int, n)
		for j := 0; j < n; j++ {
			var name, kind string
			fmt.Fscanln(reader, &name, &kind)
			_, exists := countKind[kind]
			if !exists {
				countKind[kind] = 1
			} else {
				countKind[kind] += 1
			}
		}
		for _, count := range countKind {
			answer *= count + 1
		}
		fmt.Fprintf(writer, "%d\n", answer-1)
	}
	writer.Flush()
}

/*
	종류 만큼 곱 * (옷 개수C옷 선택) - 1
*/
