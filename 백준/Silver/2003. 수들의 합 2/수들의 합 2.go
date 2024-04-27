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
	var (
		n     int
		m     int
		dp    []int
		count int
	)
	fmt.Fscan(reader, &n, &m)
	dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &dp[i])
		dp[i] += dp[i-1]
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[i]-dp[j] == m {
				count++
			}
		}
	}
	fmt.Fprintln(writer, count)

}
