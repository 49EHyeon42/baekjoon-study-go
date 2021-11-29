package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(writer, "Hello World!")
	writer.Flush()
}
