package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fprintln(writer, 3, "\ndbgmlwns2")
}
