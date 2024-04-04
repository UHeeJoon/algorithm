package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var a, b int
	str := []rune("")
	fmt.Fscan(reader, &a, &b)
	if a == b {
		str = []rune("==")
	} else if a > b {
		str = []rune(">")
	} else {
		str = []rune("<")
	}
	fmt.Fprintln(writer, string(str))
}
