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
	var n, m, i, j int
	fmt.Fscan(reader, &n, &m)
	var arr = make([]int, n)
	for c := 0; c < n; c++ {
		arr[c] = c + 1
	}
	for c := 0; c < m; c++ {
		fmt.Fscan(reader, &i, &j)
		arr[i-1], arr[j-1] = arr[j-1], arr[i-1]
	}
	for _, data := range arr {
		fmt.Fprint(writer, data, " ")
	}
}
