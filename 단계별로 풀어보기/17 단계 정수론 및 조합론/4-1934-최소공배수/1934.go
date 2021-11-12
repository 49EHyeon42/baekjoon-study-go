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

	var first, second int
	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &first, &second)
		if first < second {
			first, second = second, first
		}
		fmt.Fprintf(writer, "%d\n", lcm(first, second))
	}
	writer.Flush()
}

func gcd(first, second int) int {
	if second == 0 {
		return first
	} else {
		return gcd(second, first%second)
	}
}

func lcm(first, second int) int {
	return first * second / gcd(first, second)
}
