package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var H, M int
	fmt.Fscanln(reader, &H, &M)

	M -= 45
	if M < 0 {
		H, M = H-1, M+60
		if H < 0 {
			H += 24
		}
	}
	fmt.Fprint(writer, H, M)

	writer.Flush()
}
