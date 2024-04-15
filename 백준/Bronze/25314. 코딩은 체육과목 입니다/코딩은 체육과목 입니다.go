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
	for i := 0; i < n/4; i++ {
		fmt.Fprint(writer, "long ")
	}
	fmt.Fprintln(writer, "int")
}
