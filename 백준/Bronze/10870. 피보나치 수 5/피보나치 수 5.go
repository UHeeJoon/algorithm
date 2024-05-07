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

func main() {
	defer writer.Flush()
	var n, f, s = 0, 0, 1
	fmt.Fscan(reader, &n)
	for i := 1; i <= n; i++ {
		f, s = s, f+s
	}
	fmt.Fprintln(writer, f)
}
