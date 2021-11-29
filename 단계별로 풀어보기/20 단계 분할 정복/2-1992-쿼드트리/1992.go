package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscanln(reader, &N)

	image := make([][]byte, N)
	for i := range image {
		image[i] = make([]byte, N)
	}

	var temp []byte
	for i := 0; i < N; i++ {
		line, _, _ := reader.ReadLine()
		for _, value := range line {
			if value != 13 {
				temp = append(temp, value)
			}
		}
		image[i] = temp
	}

	fmt.Println(image)

	// partition(image, 0, 0, N)

	// fmt.Fprint(writer, string(answer))

	writer.Flush()
}

var (
	answer []byte
)

func partition(image [][]byte, x, y, size int) {
	if isSquare(image, x, y, size) {
		if image[x][y] == 0 {
			answer = append(answer, image[x][y])
		}
		return
	}
	answer = append(answer, '(')
	newSize := size / 2
	partition(image, x, y, newSize)
	partition(image, x, y+newSize, newSize)
	partition(image, x+newSize, y, newSize)
	partition(image, x+newSize, y+newSize, newSize)
	answer = append(answer, ')')
}

func isSquare(image [][]byte, x, y, size int) bool {
	value := image[x][y]
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			if image[i][j] != value {
				return false
			}
		}
	}
	return true
}
