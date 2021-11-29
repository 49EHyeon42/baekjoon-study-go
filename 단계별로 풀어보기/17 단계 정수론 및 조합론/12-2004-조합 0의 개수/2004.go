package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	// ! 어려움
	base2 := basePowerNum(2, n) - basePowerNum(2, n-m) - basePowerNum(2, m)
	base5 := basePowerNum(5, n) - basePowerNum(5, n-m) - basePowerNum(5, m)

	if base2 < base5 {
		fmt.Fprint(writer, base2)
	} else {
		fmt.Fprint(writer, base5)
	}
	writer.Flush()
}

// base 승수 카운트
func basePowerNum(base, number int) (count int) {
	for number >= base {
		count, number = count+number/base, number/base
	}
	return
}
