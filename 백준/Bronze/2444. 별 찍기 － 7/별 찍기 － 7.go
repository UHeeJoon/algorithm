package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	idx := 1
	fmt.Fscan(reader, &n)
	for i := 0; i < n && i >= 0; i += idx {
		for j := 0; j < n-i-1; j++ {
			fmt.Fprint(writer, " ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer)
		if i == n-1 {
			idx = -idx
		}
	}
}
