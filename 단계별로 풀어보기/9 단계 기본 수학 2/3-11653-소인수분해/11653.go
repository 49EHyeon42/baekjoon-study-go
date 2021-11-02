package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var N int
	fmt.Fscan(reader, &N)
	for i := 2; i <= int(math.Sqrt(float64(N))); i++ {
		for N%i == 0 {
			fmt.Fprintln(writer, i)
			N /= i
		}
	}
	if N != 1 {
		fmt.Fprint(writer, N)
	}

	writer.Flush()
}

// * 메모리 924KB, 시간 100ms
// for i := 2; i <= N; i++ {
// 	for N%i == 0 {
// 		fmt.Fprintln(writer, i)
// 		N /= i
// 	}
// }

// * 메모리 948KB, 시간 4ms
// for i := 2; i <= int(math.Sqrt(float64(N))); i++ {
// 	for N%i == 0 {
// 		fmt.Fprintln(writer, i)
// 		N /= i
// 	}
// }
// if N != 1 {
// 	fmt.Fprint(writer, N)
// }
