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
	var str1, str2 string
	fmt.Fscan(reader, &str1, &str2)
	fmt.Fprintln(writer, max(reverseString(str1), reverseString(str2)))
}
