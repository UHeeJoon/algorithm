package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, left, mid, right int64
	fmt.Fscan(reader, &n)
	left = 0
	right = n
	for left <= right {
		mid = (left + right) / 2
		if float64(mid)*float64(mid) < float64(n) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Fprintln(writer, left)
}
