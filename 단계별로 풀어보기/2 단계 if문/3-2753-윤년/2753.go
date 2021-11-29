package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var year int
	fmt.Fscanln(reader, &year)

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Fprint(writer, 1)
	} else {
		fmt.Fprint(writer, 0)
	}

	writer.Flush()
}
