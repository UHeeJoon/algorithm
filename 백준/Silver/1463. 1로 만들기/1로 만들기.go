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
	var arr = make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = 999_999_999
	}
	arr[1] = 0
	for i := 1; i <= n; i++ {
		next := arr[i] + 1
		if i+1 <= n && next < arr[i+1] {
			arr[i+1] = next
		}
		if i*2 <= n && next < arr[i*2] {
			arr[i*2] = next
		}
		if i*3 <= n && next < arr[i*3] {
			arr[i*3] = next
		}

	}
	fmt.Fprintln(writer, arr[n])
}
