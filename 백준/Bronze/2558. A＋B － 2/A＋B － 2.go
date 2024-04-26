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
	var a, b int
	fmt.Fscan(reader, &a, &b)
	fmt.Fprintln(writer, a+b)

}
