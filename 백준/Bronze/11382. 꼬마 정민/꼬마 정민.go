package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var a, b, c big.Int
	fmt.Fscan(reader, &a, &b, &c)

	fmt.Fprintln(writer, a.Add(&a, b.Add(&b, &c)))
}
