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
	var a, b, v int
	fmt.Fscan(reader, &a, &b, &v)
	result := 1
	if (v-a)%(a-b) == 0 {
		result += (v - a) / (a - b)
	} else {
		result += (v-a)/(a-b) + 1
	}
	fmt.Fprintln(writer, result)
}
