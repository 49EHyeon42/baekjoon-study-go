package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var T int
	fmt.Fscanln(reader, &T)

	var docsCount, myDocsIdx int
	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &docsCount, &myDocsIdx)

	}
	writer.Flush()
}

// 지저분함, 메모리 조금 큼, 시간 12ms
