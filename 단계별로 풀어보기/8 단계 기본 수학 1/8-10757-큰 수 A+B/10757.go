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
	reverseBytes(A)
	reverseBytes(B)

	if len(A) > len(B) {
		difference := len(A) - len(B)
		for i := 0; i < difference; i++ {
			B = append(B, '0')
		}
	} else {
		difference := len(B) - len(A)
		for i := 0; i < difference; i++ {
			A = append(A, '0')
		}
	}

	answer, carry := make([]byte, len(A)+1), 0
	for i := 0; i < len(A); i++ {
		sum := int(A[i]-'0'+B[i]-'0') + carry
		if sum < 0 {
			sum += '0'
		} else if sum > 9 {
			carry = 1
		} else {
			carry = 0
		}
		answer[i] = byte(sum%10) + '0'
	}
	if carry == 1 {
		answer[len(A)] = '1'
	}
	reverseBytes(answer)
	fmt.Fprintln(writer, string(answer))
	writer.Flush()
}

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
