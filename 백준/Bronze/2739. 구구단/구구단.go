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
	express := "%d * %d = %d\n"
	fmt.Fscan(reader, &n)
	for i := 1; i <= 9; i++ {
		fmt.Fprintf(writer, express, n, i, i*n)
	}

}
