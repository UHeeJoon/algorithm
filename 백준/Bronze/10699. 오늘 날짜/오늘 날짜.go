package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	now := time.Now()
	fmt.Fprintln(writer, now.Format("2006-01-02"))
}