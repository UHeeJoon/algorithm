package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	fmt.Fscan(reader, &n, &m)
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &A[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var num int
			fmt.Fscan(reader, &num)
			fmt.Fprint(writer, A[i][j]+num, " ")
		}
		fmt.Fprintln(writer)
	}

}
