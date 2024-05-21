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
	var n int
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			fmt.Fprint(writer, " ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer)
	}
}
