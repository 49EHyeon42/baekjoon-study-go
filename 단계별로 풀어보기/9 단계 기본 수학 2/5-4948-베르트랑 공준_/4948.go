package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n int
	for {
		fmt.Fscan(reader, &n)
		if n != 0 {
			var count int
			for i := n; i <= n*2; i++ {
				if isPrimeNumber(i) {
					count += 1
				}
			}
			fmt.Fprintln(writer, count)
		} else {
			break
		}
	}
	writer.Flush()
}

func isPrimeNumber(number int) bool {
	if number != 1 {
		var i int = 2
		for i*i <= number {
			if number%i == 0 {
				return false
			}
			i += 1
		}
		return true
	} else {
		return false
	}
}

// ! 틀림
// for {
// 	fmt.Fscan(reader, &n)
// 	if n != 0 {
// 		var count int
// 		for i := n; i <= n*2; i++ {
// 			if isPrimeNumber(i) {
// 				count += 1
// 			}
// 		}
// 		fmt.Fprintln(writer, count)
// 	} else {
// 		break
// 	}
// }

// * 나중에  에라스토테네스의 체 알고리즘으로 구현
