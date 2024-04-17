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
	var n int
	fmt.Fscan(reader, &n)
	result := 1
	tmp := 1
	for i := 1; ; i++ {
		if n <= tmp {
			break
		}
		tmp += 6 * i
		result++
	}
	fmt.Fprintln(writer, result)
}
