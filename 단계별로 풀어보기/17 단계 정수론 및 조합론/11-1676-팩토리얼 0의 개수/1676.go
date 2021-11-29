package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, count int
	fmt.Fscanln(reader, &N)
	for N >= 5 {
		count, N = count+N/5, N/5
	}
	fmt.Fprint(writer, count)
	writer.Flush()
}

/*
	팩토리얼 후 0 검사는 크기가 커서 안됨
	10 = 2 * 5
*/
