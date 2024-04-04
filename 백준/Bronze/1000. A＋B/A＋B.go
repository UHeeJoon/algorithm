package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer rw.Flush()
	var a, b int
	fmt.Fscan(rw, &a, &b)
	fmt.Fprintln(rw, a+b)
}