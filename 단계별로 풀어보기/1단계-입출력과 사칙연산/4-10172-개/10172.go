package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(writer, "|\\_/|")
	fmt.Fprintln(writer, "|q p|   /}")
	fmt.Fprintln(writer, "( 0 )\"\"\"\\")
	fmt.Fprintln(writer, "|\"^\"`    |")
	fmt.Fprint(writer, "||_/=\\\\__|")
	writer.Flush()
}
