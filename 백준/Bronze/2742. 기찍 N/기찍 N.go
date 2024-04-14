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
	for i := n; i > 0; i-- {
		fmt.Fprintln(writer, i)
		// if i%1000 == 0 {
		// 	writer.Flush()
		// }
	}
}
