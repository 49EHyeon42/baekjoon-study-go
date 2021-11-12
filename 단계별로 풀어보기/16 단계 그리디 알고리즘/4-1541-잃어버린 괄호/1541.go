package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var array string
	fmt.Fscanln(reader, &array)

	isMin, sum, temp := false, 0, 0
	for i := 0; i < len(array); i++ {
		if array[i] == '+' || array[i] == '-' {
			if !isMin {
				sum, temp = sum+temp, 0
			} else {
				sum, temp = sum-temp, 0
			}
			if array[i] == '-' {
				isMin = true
			}
		} else {
			temp *= 10
			temp += int(array[i] - '0')
		}
	}
	if !isMin {
		sum, temp = sum+temp, 0
	} else {
		sum, temp = sum-temp, 0
	}
	fmt.Fprint(writer, sum)
	writer.Flush()
}

/*
* 방법 2가지
* 1. array 배열로 합 출력
* 2. '-', '+' 분해 후 합 출력
 */
