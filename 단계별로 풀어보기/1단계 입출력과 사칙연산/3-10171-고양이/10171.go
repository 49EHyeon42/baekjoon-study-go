package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(writer, "\\    /\\")
	fmt.Fprintln(writer, " )  ( ')")
	fmt.Fprintln(writer, "(  /  )")
	fmt.Fprintln(writer, " \\(__)|")
	writer.Flush()
}
