package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, firstRing, otherRing int
	fmt.Fscanln(reader, &N)
	fmt.Fscan(reader, &firstRing)
	for i := 1; i < N; i++ {
		fmt.Fscan(reader, &otherRing)
		gcd := Gcd(firstRing, otherRing)
		fmt.Fprintf(writer, "%d/%d\n", firstRing/gcd, otherRing/gcd)
	}
	writer.Flush()
}

func Gcd(first, second int) int {
	if second == 0 {
		return first
	} else {
		return Gcd(second, first%second)
	}
}
