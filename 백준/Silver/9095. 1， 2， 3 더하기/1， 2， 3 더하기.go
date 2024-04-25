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

func main() {
	defer writer.Flush()
	var tc, n int
	var arr = make([]int, 13)
	arr[1] = 1
	arr[2] = 2
	arr[3] = 4
	for i := 4; i < 13; i++ {
		arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
	}
	fmt.Fscan(reader, &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscan(reader, &n)
		fmt.Fprintln(writer, arr[n])
	}

}
