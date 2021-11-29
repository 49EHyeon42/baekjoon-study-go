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
	fmt.Fscan(reader, &str)

	str = strings.ToUpper(str)

	var array = make([]int, 26)
	for i := range str {
		array[str[i]-65] += 1
	}

	var max int
	var answer string
	for i := 0; i < 26; i++ {
		if array[i] > max {
			max = array[i]
			answer = string(byte(i) + 65)
		} else if array[i] == max {
			answer = "?"
		}
	}
	fmt.Fprint(writer, answer)
	writer.Flush()
}

// * 소스 지저분한거 같음
