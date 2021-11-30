package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var x, y [4]int
	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &x[i], &y[i])
	}
	check(&x)
	check(&y)

	fmt.Fprintln(writer, x[3], y[3])

	writer.Flush()
}

func check(arr *[4]int) {
	if arr[0] == arr[1] {
		arr[3] = arr[2]
	} else if arr[0] == arr[2] {
		arr[3] = arr[1]
	} else {
		arr[3] = arr[0]
	}
}
