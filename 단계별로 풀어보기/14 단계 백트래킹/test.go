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
	// 순열로 사용
	// for i := 1; i <= N; i++ {
	// 	array[i-1] = i
	// }`

	// perm1(array, N, M, 0)

	// perm := make([]int, M)
	// visited := make([]bool, N)
	// perm2(visited, array, perm, 0)

	fmt.Fprintln(writer, comb1(N, M))

	//comb2(array, N, M, 0, 0)

	comb := make([]int, M)
	comb3(array, comb, 0, 0)

	writer.Flush()
}

// 순열1, 사전순으로 정렬 안됨
func perm1(array []int, N, M, depth int) {
	if depth == M {
		for i := 0; i < M; i++ {
			fmt.Fprintf(writer, "%d ", array[i])
		}
		fmt.Fprintf(writer, "\n")
		return
	}
	for i := depth; i < N; i++ {
		array[depth], array[i] = array[i], array[depth]
		perm1(array, N, M, depth+1)
		array[depth], array[i] = array[i], array[depth]
	}
}

// 순열2, 사전적으로 정렬, 배열 출력, 방문 체크
func perm2(visited []bool, array, perm []int, depth int) {
	if len(perm) == depth {
		for i := 0; i < len(perm); i++ {
			fmt.Fprintf(writer, "%d ", perm[i])
		}
		fmt.Fprintf(writer, "\n")
		return
	}
	for i := 0; i < len(array); i++ {
		if !visited[i] {
			visited[i], perm[depth] = true, array[i]
			perm2(visited, array, perm, depth+1)
			visited[i] = false
		}
	}
}

// 조합1, 결과만 출력
func comb1(n, r int) int {
	if n == r || r == 0 {
		return 1
	} else {
		return comb1(n-1, r-1) + comb1(n-1, r)
	}
}

// 조합2, 출력
func comb2(array []int, N, M, index, depth int) {
	if M == 0 {
		for i := 0; i < index; i++ {
			fmt.Fprintf(writer, "%d ", array[i])
		}
		fmt.Fprintf(writer, "\n")
	} else if depth == N {
	} else {
		array[index] = depth
		comb2(array, N, M-1, index+1, depth+1)
		comb2(array, N, M, index, depth+1)
	}
}

// 조합3, 배열 출력
func comb3(array, comb []int, index, depth int) {
	if len(comb) == depth {
		for i := 0; i < len(comb); i++ {
			fmt.Fprintf(writer, "%d ", comb[i])
		}
		fmt.Fprintf(writer, "\n")
	} else {
		for i := index; i < len(array); i++ {
			comb[depth] = array[index]
			comb3(array, comb, index+1, depth+1)
		}
	}
}

// 중복순열
func permR() {}

// 중복조합
func combR() {

}

// https://ansohxxn.github.io/categories/algorithm
