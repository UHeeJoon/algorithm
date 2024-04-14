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
	var n int
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Fprint(writer, " ")
		}
		for j := 0; j < n-i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer)
	}
}
