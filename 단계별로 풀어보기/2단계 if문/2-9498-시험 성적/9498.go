package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var testScore int
	fmt.Fscanln(reader, &testScore)

	if testScore >= 90 {
		fmt.Fprint(writer, "A")
	} else if testScore >= 80 {
		fmt.Fprint(writer, "B")
	} else if testScore >= 70 {
		fmt.Fprint(writer, "C")
	} else if testScore >= 60 {
		fmt.Fprint(writer, "D")
	} else {
		fmt.Fprint(writer, "F")
	}

	writer.Flush()
}
