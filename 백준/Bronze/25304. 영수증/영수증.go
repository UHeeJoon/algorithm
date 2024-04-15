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
	var x, n, a, b int
	fmt.Fscan(reader, &x, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a, &b)
		x -= a * b
	}
	if x != 0 {
		fmt.Fprintln(writer, "No")
	} else {
		fmt.Fprintln(writer, "Yes")
	}
}
