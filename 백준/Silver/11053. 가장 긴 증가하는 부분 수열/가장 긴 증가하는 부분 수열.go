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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()
	var n, cnt int
	fmt.Fscan(reader, &n)
	arr := make([]int, n+1)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			if cnt < dp[i] {
				cnt = dp[i]
			}
		}
	}

	fmt.Fprintln(writer, cnt)
}
