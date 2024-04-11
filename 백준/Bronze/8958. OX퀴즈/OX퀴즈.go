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
	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var str string
		score := 0
		count := 0
		fmt.Fscan(reader, &str)
		for _, data := range []rune(str) {
			if data == 'O' {
				count++
				score += count
			} else {
				count = 0
			}
		}
		fmt.Fprintln(writer, score)
	}
}
