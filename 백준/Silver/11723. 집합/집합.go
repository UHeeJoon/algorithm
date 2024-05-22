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
	var (
		n     int
		set   int
		str   string
		value int
	)
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &str, &value)
		switch str {
		case "add":
			set |= (1 << value)
		case "remove":
			set &= ^(1 << value)
		case "check":
			if (set & (1 << value)) == 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, 1)
			}
		case "toggle":
			set ^= (1 << value)
		case "all":
			set = 1<<22 - 1
		case "empty":
			set = 0
		}
	}
}
