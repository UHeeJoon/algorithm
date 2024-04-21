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
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	var dp = make([][]int, n+1)
	dp[0] = make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			fmt.Fscan(reader, &dp[i][j])
			dp[i][j] += dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}
	var x1, y1, x2, y2 int
	for m > 0 {
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		fmt.Fprintln(writer, dp[x2][y2]-dp[x2][y1-1]-dp[x1-1][y2]+dp[x1-1][y1-1])
		m--
	}
}
