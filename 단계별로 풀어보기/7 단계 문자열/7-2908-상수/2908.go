package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var str1, str2 string
	fmt.Fscan(reader, &str1, &str2)

	str1, str2 = reverse(str1), reverse(str2)
	if str1 > str2 {
		fmt.Fprint(writer, str1)
	} else {
		fmt.Fprint(writer, str2)
	}

	writer.Flush()
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
