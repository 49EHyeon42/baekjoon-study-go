package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var str string
	fmt.Fscan(reader, &str)
	var count int
	for i := 0; i < len(str); i++ {
		if i < len(str)-1 && str[i] == 99 {
			if str[i+1] == 61 || str[i+1] == 45 {
				i += 1
			}
		} else if str[i] == 100 {
			if i < len(str)-2 && str[i+1] == 122 && str[i+2] == 61 {
				i += 2
			} else if i < len(str)-1 && str[i+1] == 45 {
				i += 1
			}
		} else if i < len(str)-1 && str[i] == 108 && str[i+1] == 106 {
			i += 1
		} else if i < len(str)-1 && str[i] == 110 && str[i+1] == 106 {
			i += 1
		} else if i < len(str)-1 && str[i] == 115 && str[i+1] == 61 {
			i += 1
		} else if i < len(str)-1 && str[i] == 122 && str[i+1] == 61 {
			i += 1
		}
		count += 1
	}
	fmt.Fprint(writer, count)
	writer.Flush()
}
