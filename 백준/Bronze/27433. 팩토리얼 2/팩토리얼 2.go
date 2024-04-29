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
	var n, result int = 0, 1
	fmt.Fscan(reader, &n)
	for i := 2; i <= n; i++ {
		result *= i
	}
	fmt.Fprintln(writer, result)

}
