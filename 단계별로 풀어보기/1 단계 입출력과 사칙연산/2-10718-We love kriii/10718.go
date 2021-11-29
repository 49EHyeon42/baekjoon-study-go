package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fprintln(writer, "강한친구 대한육군")
	fmt.Fprint(writer, "강한친구 대한육군")

	writer.Flush()
}
