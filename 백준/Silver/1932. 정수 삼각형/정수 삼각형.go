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

func max_arr(arr []int) int {
	_max := -1
	for _, d := range arr {
		if _max < d {
			_max = d
		}
	}
	return _max
}

func max_int(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var arr = make([][]int, n)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, i+1)
		dp[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			fmt.Fscan(reader, &arr[i][j])
		}
	}
	dp[0][0] = arr[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + arr[i][j]
			} else if j == i {
				dp[i][j] = dp[i-1][j-1] + arr[i][j]
			} else {
				dp[i][j] = arr[i][j] + max_int(dp[i-1][j], dp[i-1][j-1])
			}
		}
	}
	fmt.Fprintln(writer, max_arr(dp[n-1]))
}
