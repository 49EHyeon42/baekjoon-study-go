package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var str string
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)

	if str == "" {
		writer.WriteString("0")
	} else {
		strArray := strings.Split(str, " ")
		fmt.Fprint(writer, len(strArray))
	}
	writer.Flush()
}
