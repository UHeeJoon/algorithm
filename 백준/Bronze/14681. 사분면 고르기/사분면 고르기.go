package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var x, y, result int
	fmt.Fscan(reader, &x, &y)
	if x >= 0 && y >= 0 {
		result = 1
	} else if x < 0 && y >= 0 {
		result = 2
	} else if x < 0 && y < 0 {
		result = 3
	} else {
		result = 4
	}
	fmt.Fprintln(writer, result)
}
