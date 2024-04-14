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
	var n, a, b int
	fmt.Fscan(reader, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d + %d = %d\n", i, a, b, (a + b))
	}
}
