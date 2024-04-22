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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	fmt.Fprintln(writer, (n+m)*(n-m))
}
