package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var A, B []byte
	fmt.Fscanln(reader, &A, &B)

	if len(A) < len(B) {
		A, B = B, A
	}

	initBytes(A)
	initBytes(B)

	difference := len(A) - len(B)
	for i := 0; i < difference; i++ {
		B = append(B, 0)
	}

	answer := make([]byte, len(A)+1)
	for i := 0; i < len(A); i++ {
		value := A[i] + B[i] + answer[i]
		answer[i], answer[i+1] = value%10, answer[i+1]+(value/10)
	}

	if answer[len(answer)-1] != 0 {
		fmt.Fprintf(writer, "%c", answer[len(answer)-1]+'0')
	}
	for i := len(answer) - 2; i >= 0; i-- {
		fmt.Fprintf(writer, "%c", answer[i]+'0')
	}

	writer.Flush()
}

func initBytes(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = b[i] - '0'
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
