package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	var N, M int
	fmt.Fscanln(reader, &N, &M)

	array := make([]int, N)
	for i := 1; i <= N; i++ {
		array[i-1] = i
	}

	Rpermutation(array, N, M, 0)

	writer.Flush()
}

func Rpermutation(array []int, N, M, depth int) {
	if depth == M {
		for i := 0; i < M; i++ {
			fmt.Fprintf(writer, "%d ", array[i])
		}
		fmt.Fprintf(writer, "\n")
		return
	}
	for i := 1; i <= N; i++ {
		array[depth] = i
		Rpermutation(array, N, M, depth+1)
	}
}

/*
* 백트래킹 문제
* 중복순열 알고리즘으로 판단
! 통과, 시간 오래걸림, 수정 필요
*/
