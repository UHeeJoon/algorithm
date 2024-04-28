package main

import (
	"bufio"
	"fmt"
	"os"
)

const DP_DEFAULT = 999_999_999

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var dp = make([]int, n+6)
	for i := 0; i < n+6; i++ {
		dp[i] = DP_DEFAULT
	}
	dp[3] = 1
	dp[5] = 1
	for i := 5; i <= n; i++ {
		dp[i] = min(dp[i], dp[i-3]+1)
		dp[i] = min(dp[i], dp[i-5]+1)
	}
	if dp[n] == DP_DEFAULT {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, dp[n])
	}
}
