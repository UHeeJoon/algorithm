package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrom(str []rune) byte {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return 0
		}
	}
	return 1
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var input string
	fmt.Fscan(reader, &input)
	fmt.Fprintln(writer, isPalindrom([]rune(input)))
}
