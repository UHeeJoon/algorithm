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
	for j := 0; j <= n/3; j++ {
		for i := 0; i <= n/5; i++ {
			if i*5+j*3 == n {
				fmt.Fprintln(writer, i+j)
				return
			}
		}
	}
	fmt.Fprintln(writer, -1)
}
