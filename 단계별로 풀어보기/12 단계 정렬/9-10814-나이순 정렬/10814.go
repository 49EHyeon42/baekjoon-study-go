package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type member struct {
	age  int
	name string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscanln(reader, &N)

	array := make([]member, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &array[i].age, &array[i].name)
	}
	sort.SliceStable(array, func(i, j int) bool {
		return array[i].age < array[j].age
	})
	for i := range array {
		fmt.Fprintf(writer, "%d %s\n", array[i].age, array[i].name)
	}
	writer.Flush()
}

// 메모리 5780KB, 276ms, sort.Slice 사용으로 안풀림
// sort.SliceStable() 해결 무슨 차이?
