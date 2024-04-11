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
	var alpha [26]int
	for idx := range alpha {
		alpha[idx] = -1
	}
	var input string
	fmt.Fscan(reader, &input)
	for idx, data := range []rune(input) {
		if alpha[data-'a'] == -1 {
			alpha[data-'a'] = idx
		}
	}
	for _, data := range alpha {
		fmt.Fprint(writer, data, " ")
	}
}
