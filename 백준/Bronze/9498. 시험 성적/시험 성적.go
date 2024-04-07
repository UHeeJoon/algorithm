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
	var n int
	var grade rune
	fmt.Fscan(reader, &n)
	switch {
	case n >= 90:
		grade = rune('A')
	case n >= 80:
		grade = rune('B')
	case n >= 70:
		grade = rune('C')
	case n >= 60:
		grade = rune('D')
	default:
		grade = rune('F')
	}
	fmt.Fprintln(writer, string(grade))
}
