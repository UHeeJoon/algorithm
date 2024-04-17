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
	var tc, n int
	fmt.Fscan(reader, &tc)
	for ; tc > 0; tc-- {
		fmt.Fscan(reader, &n)
		fmt.Fprint(writer, n/25, " ")
		n %= 25
		fmt.Fprint(writer, n/10, " ")
		n %= 10
		fmt.Fprint(writer, n/5, " ")
		n %= 5
		fmt.Fprintln(writer, n/1)
	}
}
