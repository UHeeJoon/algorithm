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

func chore_max(a, b, c int) int {
	if (b <= a && a <= c) || (c <= a && a <= b) {
		return a
	}
	if (a <= b && b <= c) || (c <= b && b <= a) {
		return b
	}
	return c
}
func main() {
	defer writer.Flush()
	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)
	fmt.Fprintln(writer, chore_max(a, b, c))
}
