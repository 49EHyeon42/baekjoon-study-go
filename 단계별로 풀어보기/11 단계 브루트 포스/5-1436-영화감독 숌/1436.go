package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscanln(reader, &N)

	answer, count := 666, 1
	for count != N {
		answer += 1
		if strings.Contains(strconv.Itoa(answer), "666") {
			count += 1
		}
	}
	fmt.Fprintln(writer, answer)
	writer.Flush()
}

// 7684KB, 156ms
// int형 말고 문자형으로 풀어보기
