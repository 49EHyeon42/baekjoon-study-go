package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	var array []bool = make([]bool, 10001)

	isSelfNumber(array)

	for i := 1; i < len(array); i++ {
		if !array[i] {
			fmt.Fprintln(writer, i)
		}
	}

	writer.Flush()
}

func isSelfNumber(array []bool) {
	for i := 1; i < len(array); i++ {
		var sum, number int = i, i
		for number != 0 {
			sum = sum + number%10
			number = number / 10
		}
		if sum < len(array) {
			array[sum] = true
		}
	}
}
