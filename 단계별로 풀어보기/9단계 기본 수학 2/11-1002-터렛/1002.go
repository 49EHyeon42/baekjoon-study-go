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
	var x1, y1, r1, x2, y2, r2 float64
	fmt.Fscanln(reader, &T)
	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &x1, &y1, &r1, &x2, &y2, &r2)
		fmt.Fprintln(writer, point(x1, y1, r1, x2, y2, r2))
	}
	writer.Flush()
}

func point(x1, y1, r1, x2, y2, r2 float64) int {
	distancePow := int((math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)))
	switch {
	case x1 == x2 && y1 == y2 && r1 == r2:
		return -1
	case distancePow > int(math.Pow(r1+r2, 2)):
		return 0
	case distancePow < int(math.Pow(r2-r1, 2)):
		return 0
	case distancePow == int(math.Pow(r2-r1, 2)):
		return 1
	case distancePow == int(math.Pow(r1+r2, 2)):
		return 1
	default:
		return 2
	}
}
