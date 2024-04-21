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
	var arr = make([][]int, n)
	var dp = make([][]int, n+1)
	dp[0] = make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		dp[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &arr[i][j])
			dp[i+1][j+1] = arr[i][j] + dp[i+1][j] + dp[i][j+1] - dp[i][j]
		}
	}
	var x1, y1, x2, y2 int
	for m > 0 {
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		fmt.Fprintln(writer, dp[x2][y2]-dp[x2][y1-1]-dp[x1-1][y2]+dp[x1-1][y1-1])
		m--
	}
}
