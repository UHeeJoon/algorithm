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
	var n, m, k int
	fmt.Fscan(reader, &n, &m)
	var arr = make([][]int, n+1)
	arr[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		arr[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			fmt.Fscan(reader, &arr[i][j])
			arr[i][j] += arr[i][j-1] + arr[i-1][j] - arr[i-1][j-1]
		}
	}
	fmt.Fscan(reader, &k)
	var x1, y1, x2, y2 int
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		fmt.Fprintln(writer, arr[x2][y2]-arr[x2][y1-1]-arr[x1-1][y2]+arr[x1-1][y1-1])
	}

}
