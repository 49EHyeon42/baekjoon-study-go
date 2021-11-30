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

	var number float64
	fmt.Fscanln(reader, &number)
	number = math.Pow(number, 2)
	fmt.Fprintf(writer, "%f\n%f\n", math.Pi*number, 2*number)

	writer.Flush()
}
