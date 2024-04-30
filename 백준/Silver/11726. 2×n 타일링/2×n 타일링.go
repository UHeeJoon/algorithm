package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

const D_NUM = 10_007

func main() {
	defer writer.Flush()
	var n, num, result int = 0, 1, 1
	fmt.Fscan(reader, &n)
	for i := 2; i <= n; i++ {
		num, result = result, (num+result)%D_NUM
	}
	fmt.Fprintln(writer, result)
}
