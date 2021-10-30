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

	data, max, min := 0, -1000000, 1000000
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &data)
		if max < data {
			max = data
		}
		if min > data {
			min = data
		}
	}
	fmt.Fprint(writer, min, max)

	writer.Flush()
}

/*
* 방법 1
var data, max, min int
for i := 0; i < N; i++ {
	fmt.Fscan(reader, &data)
	if max < data {
		max = data
	}
	if min > data {
		min = data
	}
}
fmt.Fprint(writer, min, max)


* 방법 2
sort.Ints(array)
fmt.Fprint(writer, array[0], array[N-1])
*/
