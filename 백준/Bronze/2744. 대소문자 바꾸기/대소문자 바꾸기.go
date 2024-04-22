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
	var str string
	fmt.Fscan(reader, &str)
	for _, s := range []rune(str) {
		if s >= 97 {
			fmt.Fprint(writer, string(s-32))
		} else {
			fmt.Fprint(writer, string(s+32))
		}
	}
}
