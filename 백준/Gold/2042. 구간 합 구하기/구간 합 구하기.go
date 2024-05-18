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

func init_fenw(fenw []int64, data int64, idx int) {
	for idx <= len(fenw) {
		fenw[idx] += data
		idx += (idx & -idx)
	}
}

func query_fenw(fenw []int64, idx int) int64 {
	var answer int64
	for idx > 0 {
		answer += fenw[idx]
		idx -= (idx & -idx)
	}
	return answer
}

func main() {
	defer writer.Flush()
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	var arr = make([]int64, n)
	var fenw = make([]int64, n+4)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &arr[i-1])
		init_fenw(fenw, arr[i-1], i)
	}
	var a, b int
	var c, ans int64
	for i := 0; i < m+k; i++ {
		fmt.Fscan(reader, &a, &b, &c)
		if a == 1 {
			init_fenw(fenw, c-arr[b-1], b)
			arr[b-1] = c
		} else if a == 2 {
			ans = query_fenw(fenw, int(c)) - query_fenw(fenw, b-1)
			fmt.Fprintln(writer, ans)
		}
	}
}
