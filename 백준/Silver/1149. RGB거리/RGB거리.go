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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscanln(reader, &n)
	var dp = make([][]int, n+1)
	dp[0] = make([]int, 3)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, 3)
		fmt.Fscanln(reader, &dp[i][0], &dp[i][1], &dp[i][2])
	}
	for i := 1; i <= n; i++ {
		dp[i][0] += min(dp[i-1][1], dp[i-1][2])
		dp[i][1] += min(dp[i-1][0], dp[i-1][2])
		dp[i][2] += min(dp[i-1][0], dp[i-1][1])
	}
	fmt.Fprintln(writer, min(dp[n][0], min(dp[n][1], dp[n][2])))
}
