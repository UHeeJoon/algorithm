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
	var n int
	fmt.Fscanln(reader, &n)
	var arr, dp = make([]int, n+3), make([]int, n+3)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &arr[i])
	}
	dp[1] = arr[1]
	dp[2] = arr[1] + arr[2]
	dp[3] = arr[3] + max(arr[1], arr[2])
	for i := 4; i <= n; i++ {
		dp[i] = arr[i] + max(dp[i-2], arr[i-1]+dp[i-3])
	}
	fmt.Fprintln(writer, dp[n])
}
