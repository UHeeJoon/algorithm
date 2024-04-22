package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func abs(a, b int) int {
	if a-b < 0 {
		return -(a - b)
	}
	return a - b
}
func main() {
	defer writer.Flush()
	var n, m int
	fmt.Fscan(reader, &n, &m)
	fmt.Fprintln(writer, abs(n, m))
}
