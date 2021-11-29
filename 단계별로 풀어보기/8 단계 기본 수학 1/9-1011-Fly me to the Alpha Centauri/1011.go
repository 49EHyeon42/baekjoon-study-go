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
	fmt.Fscan(reader, &T)

	for i := 0; i < T; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)

		fmt.Fprintln(writer, solution(x, y))
	}

	writer.Flush()
}

func solution(x, y int) int {

}

func getMinDistance() {

}
