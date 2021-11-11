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

	array, visited := make([]int, N+1), make([]bool, N+1)

	permutation(array, visited, N, M, 0)

	writer.Flush()
}

func permutation(array []int, visited []bool, N, M, depth int) {
	if depth == M {
		for i := 0; i < M; i++ {
			fmt.Fprintf(writer, "%d ", array[i])
		}
		fmt.Fprintf(writer, "\n")
		return
	}
	for i := 1; i <= N; i++ {
		if !visited[i] {
			visited[i], array[depth] = true, i
			permutation(array, visited, N, M, depth+1)
			visited[i] = false
		}
	}
}

/*
* 백트래킹 문제
* 순열 알고리즘으로 판단
* 방법 2가지 = 재귀 사용, 힙 알고리즘 사용

! 수열 출력 가능, 사전 순으로 출력 불가
func main() {
	var N, M int
	fmt.Fscanln(reader, &N, &M)

	array := make([]int, N)
	for i := 1; i <= N; i++ {
		array[i-1] = i
	}

	permutation(array, N, M, 0)

	writer.Flush()
}

func permutation(array []int, N, M, depth int) {
	if depth == M {
		for i := 0; i < M; i++ {
			fmt.Fprintf(writer, "%d ", array[i])
		}
		fmt.Fprintf(writer, "\n")
	}
	for i := depth; i < N; i++ {
		array[depth], array[i] = array[i], array[depth]
		permutation(array, N, M, depth+1)
		array[depth], array[i] = array[i], array[depth]
	}
}
*/
