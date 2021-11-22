package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	answer, col, N = 0, make([]int, 15), 0
)

func main() {
	fmt.Fscanln(reader, &N)
	dfs(0)
	fmt.Fprint(writer, answer)
	writer.Flush()
}

func dfs(x int) {
	if x == N {
		answer += 1
		return
	}
	for i := 0; i < N; i++ {
		col[x] = i
		if isValid(x) {
			dfs(x + 1)
		}
	}

}

func isValid(level int) bool {
	for i := 0; i < level; i++ {
		if col[i] == col[level] || level-i == int(math.Abs(float64(col[level]-col[i]))) {
			return false
		}
	}
	return true
}

// * 다시 풀어보기, 시간 단축 가능
