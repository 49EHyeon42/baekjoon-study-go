package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscan(reader, &N)

	array := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &array[i])
	}
	var answer1 float64
	answer2, answer3, answer4 := 0, 0, 0
	for _, number := range array {
		answer1 += float64(number)
	}
	answer1 /= float64(N)

	sort.Ints(array)
	answer2 = array[(N+1)/2-1]

	answer3 = mode(array)

	answer4 = array[N-1] - array[0]

	fmt.Fprintf(writer, "%.0f\n%d\n%d\n%d", answer1, answer2, answer3, answer4)

	writer.Flush()
}

// ! 이해 못함, 어려움, 지저분함
func mode(array []int) int {
	temp, count, index, max := make([]int, 8001), 0, 0, 0
	for i := 0; i < len(array); i++ {
		index = array[i] + 4000
		temp[index] += 1

		if temp[index] > max {
			max = temp[index]
		}
	}
	for i := 0; i < 8001; i++ {
		if temp[i] == 0 {
			continue
		}
		if temp[i] == max {
			if count == 0 {
				index, count = i, count+1
			} else if count == 1 {
				index = i
				break
			}
		}
	}
	return index - 4000
}
