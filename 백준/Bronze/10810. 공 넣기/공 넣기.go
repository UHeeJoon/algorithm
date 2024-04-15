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
	var n, m, i, j, k int
	fmt.Fscan(reader, &n, &m)
	var arr = make([]int, n)
	for idx := 0; idx < m; idx++ {
		fmt.Fscan(reader, &i, &j, &k)
		for arr_idx := i - 1; arr_idx <= j-1; arr_idx++ {
			arr[arr_idx] = k
		}
	}
	for _, data := range arr {
		fmt.Fprint(writer, data, " ")
	}
}
