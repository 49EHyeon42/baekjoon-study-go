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
	fmt.Fscanln(reader, &N)

	array := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &array[i])
	}
	sort.Ints(array)
	/*
		arr[1]-arr[0]=M*(arr[1]/M - arr[0]/M)
		arr[2]-arr[1]=M*(arr[2]/M - arr[1]/M)
		arr[i]-arr[i-1]=M*(arr[i]/M - arr[i-1]/M)
	*/
	gcd := array[1] - array[0]
	for i := 2; i < N; i++ {
		gcd = Gcd(gcd, array[i]-array[i-1])
	}
	var result []int
	for i := 1; i*i <= gcd; i++ {
		if gcd%i == 0 {
			result = append(result, i)
			if i != gcd/i {
				result = append(result, gcd/i)
			}
		}
	}
	sort.Ints(result)
	for i := 1; i < len(result); i++ {
		fmt.Fprintf(writer, "%d ", result[i])
	}
	writer.Flush()
}

func Gcd(first, second int) int {
	if second == 0 {
		return first
	} else {
		return Gcd(second, first%second)
	}
}
