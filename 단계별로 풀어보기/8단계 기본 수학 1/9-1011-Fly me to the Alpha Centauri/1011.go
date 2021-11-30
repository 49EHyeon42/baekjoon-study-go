package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var T int
	fmt.Fscan(reader, &T)

	var x, y, answer, length int
	for i := 0; i < T; i++ {
		fmt.Fscan(reader, &x, &y)

		answer, length = int(math.Sqrt(float64(y-x))), y-x
		if length == answer*answer {
			answer = answer*2 - 1
		} else if answer*answer < length && length <= answer*answer+answer {
			answer = answer * 2
		} else if answer*answer+answer < length && length < (answer+1)*(answer+1) {
			answer = answer*2 + 1
		}
		fmt.Fprintln(writer, answer)
	}

	writer.Flush()
}
