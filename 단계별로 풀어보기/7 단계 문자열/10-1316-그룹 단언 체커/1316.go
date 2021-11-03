package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscan(reader, &N)
	var count int = N
	for i := 0; i < N; i++ {
		var str string
		fmt.Fscan(reader, &str)
		var array = make([]bool, 26)
		for j := 0; j < len(str); j++ {
			if array[str[j]-97] == true && str[j-1] != str[j] {
				count -= 1
				break
			} else {
				array[str[j]-97] = true
			}
		}
	}
	fmt.Fprint(writer, count)
	writer.Flush()
}

// * 더 짧고 간단한 알고리즘이 있을 것 같음
